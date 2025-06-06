package eventAdapter

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	EventUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/event"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type EventHandler struct {
	EventUsecase EventUsecase.EventUsecase
}

func NewEventAdapter(eventUsecase EventUsecase.EventUsecase) *EventHandler {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}
func (h *EventHandler) GetAll(c *fiber.Ctx) error {
	events, err := h.EventUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", events)
}
func (h *EventHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	event, err := h.EventUsecase.GetByID(&eventID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", event)
}
func (h *EventHandler) Create(c *fiber.Ctx) error {
	var event Entities.Event

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", err.Error(), nil)
	}
	if form == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", "failed to parse form", nil)
	}
	if form.Value["author_id"] == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Author ID is required", "Author ID is required", nil)
	}
	authorID, err := utils.StringToUint(form.Value["author_id"][0])
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid author ID format", err.Error(), nil)
	}
	event.AuthorID = authorID
	if form.Value["title"] != nil {
		event.Title = form.Value["title"][0]
	}
	if form.Value["description"] != nil {
		event.Description = form.Value["description"][0]
	}
	if form.Value["date_time"] != nil {
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", form.Value["date_time"][0])
		if err != nil {
			return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid date_time format", err.Error(), nil)
		}
		event.DateTime = parsedTime
	}
	fmt.Println("event.DateTime", event.DateTime)
	if form.Value["location"] != nil {
		event.Location = form.Value["location"][0]
	}
	if form.Value["author_id"] != nil {
		authorID, err := utils.StringToUint(form.Value["author_id"][0])
		if err != nil {
			return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid author ID format", err.Error(), nil)
		}
		event.AuthorID = authorID
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.BucketKey))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create storage client", err.Error(), nil)
	}
	defer client.Close()

	files := form.File["event_image"]
	if len(files) >= 1 {
		bucketName := config.BucketName
		shareableURLs := make([]string, 0, len(files))

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Error opening file", err.Error(), nil)
			}

			objectName := fmt.Sprintf("images/%d_%s", time.Now().UnixNano(), fileHeader.Filename)
			token := uuid.New().String()

			wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
			wc.Metadata = map[string]string{
				"firebaseStorageDownloadTokens": token,
			}

			if _, err = io.Copy(wc, file); err != nil {
				file.Close()
				wc.Close()
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to write file to bucket", err.Error(), nil)
			}
			file.Close()
			if err := wc.Close(); err != nil {
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to close writer", err.Error(), nil)
			}

			shareableURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s",
				bucketName, url.QueryEscape(objectName), token)
			shareableURLs = append(shareableURLs, shareableURL)
		}
		event.Event_image = shareableURLs[0]
	}
	fmt.Println("user.UserImage", event.Event_image)

	if err := h.EventUsecase.Create(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", event)
}
func (h *EventHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	var event Entities.Event
	if err := c.BodyParser(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	event.ID = eventID
	if err := h.EventUsecase.Update(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", event)
}
func (h *EventHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	if err := h.EventUsecase.DeleteByID(&eventID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}
func (h *EventHandler) GetByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID is required", "User ID is required", nil)
	}
	userIDUint, err := utils.StringToUint(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}
	events, err := h.EventUsecase.GetByUserID(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", events)
}

func (h *EventHandler) JoinEvent(c *fiber.Ctx) error {
	eventIDStr := c.Params("id")
	if eventIDStr == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Event ID is required", "Event ID is required", nil)
	}
	eventID, err := utils.StringToUint(eventIDStr)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid Event ID format", err.Error(), nil)
	}
	userID := c.Locals("userID")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid user ID type", "Invalid user ID type", nil)
	}
	fmt.Println("UserID from locals:", userID)
	userIDUint := uint(userIDFloat)
	if err := h.EventUsecase.JoinEvent(&eventID, &userIDUint); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}