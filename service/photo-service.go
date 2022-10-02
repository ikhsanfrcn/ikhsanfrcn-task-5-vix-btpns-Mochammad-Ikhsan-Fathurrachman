package service

import (
	"fmt"
	"log"

	"github.com/ikhsanfrcn/vix-btpn/dto"
	"github.com/ikhsanfrcn/vix-btpn/entity"
	"github.com/ikhsanfrcn/vix-btpn/repository"
	"github.com/mashingan/smapping"
)

// PhotoService is a ....
type PhotoService interface {
	Insert(b dto.PhotoCreateDTO) entity.Photo
	Update(b dto.PhotoUpdateDTO) entity.Photo
	Delete(b entity.Photo)
	All() []entity.Photo
	FindByID(photoID uint64) entity.Photo
	IsAllowedToEdit(userID string, photoID uint64) bool
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

// NewPhotoService .....
func NewPhotoService(photoRepo repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepo,
	}
}

func (service *photoService) Insert(b dto.PhotoCreateDTO) entity.Photo {
	photo := entity.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.InsertPhoto(photo)
	return res
}

func (service *photoService) Update(b dto.PhotoUpdateDTO) entity.Photo {
	photo := entity.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.UpdatePhoto(photo)
	return res
}

func (service *photoService) Delete(b entity.Photo) {
	service.photoRepository.DeletePhoto(b)
}

func (service *photoService) All() []entity.Photo {
	return service.photoRepository.AllPhoto()
}

func (service *photoService) FindByID(photoID uint64) entity.Photo {
	return service.photoRepository.FindPhotoByID(photoID)
}

func (service *photoService) IsAllowedToEdit(userID string, photoID uint64) bool {
	b := service.photoRepository.FindPhotoByID(photoID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
