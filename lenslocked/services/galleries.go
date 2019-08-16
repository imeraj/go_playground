package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/imeraj/go_playground/lenslocked/models"
	galleriesrepo "github.com/imeraj/go_playground/lenslocked/repositories/gallery"
)

type GalleryService struct {
	repo *galleriesrepo.GalleryRepo
}

func NewGalleryService() *GalleryService {
	repo := galleriesrepo.NewGalleryRepo()

	return &GalleryService{
		repo: repo,
	}
}

func (gs *GalleryService) Create(gallery *models.Gallery) error {
	return gs.repo.Create(gallery)
}

func (gs *GalleryService) ByUserID(userID uint) ([]models.Gallery, error) {
	return gs.repo.ByUserID(userID)
}

func (gs *GalleryService) ByID(galleryID uint) (*models.Gallery, error) {
	return gs.repo.ByID(galleryID)
}

func (gs *GalleryService) Delete(galleryID uint) error {
	gallery, err := gs.ByID(galleryID)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(gs.imagePath(gallery.ID)); err != nil {
		return err
	}

	return gs.repo.Delete(galleryID)
}

func (gs *GalleryService) Update(gallery *models.Gallery) error {
	return gs.repo.Update(gallery)
}

func (gs *GalleryService) ProcessImages(galleryPath string, files []*multipart.FileHeader) error {
	for _, f := range files {
		file, err := f.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join(galleryPath, f.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gs *GalleryService) DeleteImage(gallery *models.Gallery, image *models.Image) error {
	if err := os.Remove(image.RelativePath()); err != nil {
		return err
	}

	return nil
}

func (gs *GalleryService) ByGalleryID(galleryID uint) ([]models.Image, error) {
	path := gs.imagePath(galleryID)
	strings, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}

	ret := make([]models.Image, len(strings))
	for i, imageStr := range strings {
		ret[i] = models.Image{
			Filename:  filepath.Base(imageStr),
			GalleryID: galleryID,
		}
	}

	return ret, nil
}

func (gs *GalleryService) imagePath(galleryID uint) string {
	return filepath.Join("public", "images", "galleries", fmt.Sprintf("%v", galleryID))
}
