package EventUsecase

import Entities "github.com/Farenthigh/Fitbuddy-BE/entities"

type EventUsecase interface {
	GetAll() ([]Entities.Event, error)
	GetByID(id *uint) (*Entities.Event, error)
	Create(event *Entities.Event) error
	Update(event *Entities.Event) error
	DeleteByID(id *uint) error
	GetByUserID(userID *uint) ([]Entities.Event, error)
	JoinEvent(eventID *uint, UserID *uint) error
}

type EventService struct {
	eventRepo Eventrepository
}

func NewEventService(eventRepo Eventrepository) EventUsecase {
	return &EventService{
		eventRepo: eventRepo,
	}
}

func (service *EventService) GetAll() ([]Entities.Event, error) {
	events, err := service.eventRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (service *EventService) GetByID(id *uint) (*Entities.Event, error) {
	event, err := service.eventRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}
func (service *EventService) Create(event *Entities.Event) error {
	if err := service.eventRepo.Create(event); err != nil {
		return err
	}
	return nil
}
func (service *EventService) Update(event *Entities.Event) error {
	selectedEvent, err := service.eventRepo.GetByID(&event.ID)
	if err != nil {
		return err
	}
	if selectedEvent == nil {
		return nil
	}
	selectedEvent.Title = event.Title
	selectedEvent.Description = event.Description
	selectedEvent.Location = event.Location
	selectedEvent.DateTime = event.DateTime

	if err := service.eventRepo.Update(selectedEvent); err != nil {
		return err
	}
	return nil
}
func (service *EventService) DeleteByID(id *uint) error {
	if err := service.eventRepo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
func (service *EventService) GetByUserID(userID *uint) ([]Entities.Event, error) {
	events, err := service.eventRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (service *EventService) JoinEvent(eventID *uint, UserID *uint) error {
	if err := service.eventRepo.JoinEvent(eventID, UserID); err != nil {
		return err
	}
	return nil
}