package repo

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go_crawler/model"
)

type linkRepo struct {
	db *gorm.DB
}

func NewLinkRepo(db *gorm.DB) *linkRepo {
	return &linkRepo{db: db}
}

type LinkRepo interface {
	Get(id uint) (*model.Link, error)
	GetByURL(url string) (*model.Link, error)
	GetUnFetchedLink() (*model.Link, error)
	Create(url string) (*model.Link, error)
	Update(url string, isFetched bool) (*model.Link, error)
}

func (s *linkRepo) Get(id uint) (*model.Link, error) {
	var result model.Link
	err := s.db.New().First(&result, "id = ?", id).Error
	return &result, err
}

func (s *linkRepo) GetByURL(url string) (*model.Link, error) {
	var result model.Link
	err := s.db.New().First(&result, "url = ?", url).Error
	return &result, err
}

func (s *linkRepo) Create(url string) (*model.Link, error) {
	result := model.Link{
		URL: url,
	}
	err := s.db.New().Create(&result).Error
	return &result, err
}

func (s *linkRepo) Update(url string, isFetched bool) (*model.Link, error) {
	item, err := s.GetByURL(url)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("link has not exist")
	}

	item.IsFetched = isFetched

	err = s.db.New().Save(&item).Error
	return item, err
}

func (s *linkRepo) GetUnFetchedLink() (*model.Link, error) {
	var result model.Link
	err := s.db.New().First(&result, "is_fetched = ?", false).Error
	return &result, err
}
