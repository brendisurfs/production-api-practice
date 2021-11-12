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

// GetComment - retrieves comments by their id fro mteh db
func (s *Service) GetComment(ID uint) (Comment, error) {

	var comment Comment

	// retrieve the first comment with the ID.
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	// if this all passes, return the comment.
	return comment, nil
}

// GetCommentBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment

	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

// PostComment - adds a new comment to the db.
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates commment by id with new Comment info.
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)

	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// DeleteComment - deletes a comment from the db by ID.
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - gets all comments from the db.
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment

	// if the comments cant be found, handle error.
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}

	// if all good, return the comments.
	return comments, nil
}
