package services

import (
	"github.com/lyquocnam/go-crawler-learning/model"
	"github.com/lyquocnam/go-crawler-learning/repo"
	"log"
)

type updater struct {
	in       chan *model.Article
	out      chan *model.Article
	linkRepo repo.LinkRepo
}

func NewUpdater(in chan *model.Article, out chan *model.Article, linkRepo repo.LinkRepo) *updater {
	return &updater{
		in:       in,
		out:      out,
		linkRepo: linkRepo,
	}
}

func (u *updater) Listen() {
	for {
		select {
		case article := <-u.in:
			go u.updateFetchedLink(article)
		}
	}
}

func (u *updater) updateFetchedLink(article *model.Article) {
	_, err := u.linkRepo.Update(article.URL, true)
	if err != nil {
		log.Println(err)
	} else {
		u.out <- article
		log.Printf("updated %s -> is_fetched = true", article.URL)
	}
}
