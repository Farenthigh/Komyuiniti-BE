package eventAdapter

import (
	"fmt"

	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type EventGorm struct {
	db *gorm.DB
}

func NewEventGorm(db *gorm.DB) *EventGorm {
	return &EventGorm{
		db: db,
	}
}
func (g *EventGorm) GetAll() ([]Entities.Event, error) {
	var events []Entities.Event
	if err := g.db.Preload("Author").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
func (g *EventGorm) GetByID(id *uint) (*Entities.Event, error) {
	var event Entities.Event
	if err := g.db.Preload("Members").Preload("Author").Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
func (g *EventGorm) Create(event *Entities.Event) error {
	if err := g.db.Create(&event).Error; err != nil {
		return err
	}
	return nil
}
func (g *EventGorm) Update(event *Entities.Event) error {
	if err := g.db.Save(&event).Error; err != nil {
		return err
	}
	return nil
}
func (g *EventGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Event{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (g *EventGorm) GetByUserID(userID *uint) ([]Entities.Event, error) {
	var events []Entities.Event
	if err := g.db.Preload("Author").Where("user_id = ?", userID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
func (g *EventGorm) JoinEvent(eventID *uint, userID *uint) error {
	var event Entities.Event
	if err := g.db.Preload("Members").Where("id = ?", eventID).First(&event).Error; err != nil {
		return err
	}

	var user Entities.User
	if err := g.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}
	fmt.Println(&user)
	event.Members = append(event.Members, user)
	if err := g.db.Save(&event).Error; err != nil {
		return err
	}
	return nil
}
