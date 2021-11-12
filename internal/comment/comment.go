package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines our comment structure
type Comment struct {
	gorm.Model

	// Slug - path where this comment was posted
	Slug   string
	Body   string
	Author string
}

// ICommentService - the if for our comment service
type ICommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service.
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
