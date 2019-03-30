package repo

import (
	"github.com/jinzhu/gorm"
	"go_crawler/model"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *articleRepo {
	return &articleRepo{db: db}
}

type ArticleRepo interface {
	ExistByUrl(url string) (bool, error)
	Create(article *model.Article) error
}

func (r *articleRepo) ExistByUrl(url string) (bool, error) {
	count := 0
	err := r.db.New().Model(model.Article{}).Where("url = ?", url).Count(&count).Error
	return count > 0, err
}

func (r *articleRepo) Create(article *model.Article) error {
	return r.db.New().Create(&article).Error
}
