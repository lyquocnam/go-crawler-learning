package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lyquocnam/go-crawler-learning/model"
	"github.com/lyquocnam/go-crawler-learning/repo"
	"log"
	"time"
)

type watcher struct {
	linkRepo repo.LinkRepo
	out      chan *model.Link
}

func NewWatcher(linkRepo repo.LinkRepo, out chan *model.Link) *watcher {
	return &watcher{
		linkRepo: linkRepo,
		out:      out,
	}
}

func (w *watcher) Watch(deplay int) {
	for {
		<-time.After(time.Duration(deplay) * time.Second)
		go w.watchFromDb()
	}
}

func (w *watcher) watchFromDb() {
	link, err := w.linkRepo.GetUnFetchedLink()
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			log.Println(err)
		}
		return
	}

	log.Println(link)

	w.out <- link
}
