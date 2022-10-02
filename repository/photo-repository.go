package repository

import (
	"github.com/ikhsanfrcn/vix-btpn/entity"
	"gorm.io/gorm"
)

// PhotoRepository is a ....
type PhotoRepository interface {
	InsertPhoto(b entity.Photo) entity.Photo
	UpdatePhoto(b entity.Photo) entity.Photo
	DeletePhoto(b entity.Photo)
	AllPhoto() []entity.Photo
	FindPhotoByID(photoID uint64) entity.Photo
}

type photoConnection struct {
	connection *gorm.DB
}

// NewPhotoRepository creates an instance PhotoRepository
func NewPhotoRepository(dbConn *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: dbConn,
	}
}

func (db *photoConnection) InsertPhoto(b entity.Photo) entity.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) UpdatePhoto(b entity.Photo) entity.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) DeletePhoto(b entity.Photo) {
	db.connection.Delete(&b)
}

func (db *photoConnection) FindPhotoByID(photoID uint64) entity.Photo {
	var photo entity.Photo
	db.connection.Preload("User").Find(&photo, photoID)
	return photo
}

func (db *photoConnection) AllPhoto() []entity.Photo {
	var photos []entity.Photo
	db.connection.Preload("User").Find(&photos)
	return photos
}
