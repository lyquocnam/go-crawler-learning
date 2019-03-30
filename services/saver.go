package services

import (
	"github.com/lyquocnam/go-crawler-learning/model"
	"github.com/lyquocnam/go-crawler-learning/repo"
	"log"
)

type saver struct {
	in          chan *model.Article
	out         chan *model.Article
	articleRepo repo.ArticleRepo
}

func NewSaver(in chan *model.Article, out chan *model.Article, articleRepo repo.ArticleRepo) *saver {
	return &saver{
		in:          in,
		out:         out,
		articleRepo: articleRepo,
	}
}

func (s *saver) Listen() {
	for {
		select {
		case article := <-s.in:
			go s.save(article)
		}
	}
}

func (s *saver) save(article *model.Article) {
	exist, err := s.articleRepo.ExistByUrl(article.URL)
	if err != nil {
		log.Println(err)
		return
	}

	if !exist {
		err := s.articleRepo.Create(article)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("created article: %s", article.Title)
		}
	}
}
