package CommentUsecase

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
)

type CommentUsecase interface {
	GetAll() ([]Entities.Comment, error)
	GetByID(id *uint) (*Entities.Comment, error)
	Create(comment *Entities.Comment) error
	Update(comment *Entities.Comment) error
	DeleteByID(id *uint) error
}
type commentService struct {
	commentRepository CommentRepository
}

func NewCommentService(repo CommentRepository) CommentUsecase {
	return &commentService{
		commentRepository: repo,
	}
}
func (s *commentService) GetAll() ([]Entities.Comment, error) {
	comments, err := s.commentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (s *commentService) GetByID(id *uint) (*Entities.Comment, error) {
	comment, err := s.commentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (s *commentService) Create(comment *Entities.Comment) error {
	err := s.commentRepository.Create(comment)
	if err != nil {
		return err
	}
	return nil
}
func (s *commentService) Update(comment *Entities.Comment) error {
	err := s.commentRepository.Update(comment)
	if err != nil {
		return err
	}
	return nil
}
func (s *commentService) DeleteByID(id *uint) error {
	err := s.commentRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
