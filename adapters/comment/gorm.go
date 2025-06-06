package CommentAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"gorm.io/gorm"
)

type CommentGorm struct {
	db *gorm.DB
}

func NewCommentGorm(db *gorm.DB) *CommentGorm {
	return &CommentGorm{
		db: db,
	}
}
func (g *CommentGorm) GetAll() ([]Entities.Comment, error) {
	var comments []Entities.Comment
	if err := g.db.Preload("Author").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (g *CommentGorm) GetByID(id *uint) (*Entities.Comment, error) {
	var comment Entities.Comment
	if err := g.db.Preload("Author").Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
func (g *CommentGorm) Create(comment *Entities.Comment) error {
	if err := g.db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
func (g *CommentGorm) Update(comment *Entities.Comment) error {
	if err := g.db.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}
func (g *CommentGorm) DeleteByID(id *uint) error {
	if err := g.db.Delete(&Entities.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}
