package database

import (
	"github.com/brendisurfs/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comments table.
// 1. takes in a model.
// 2. defines all the columns and fields + predifined gorm fields.
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
