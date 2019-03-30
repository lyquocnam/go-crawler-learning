package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lyquocnam/go-crawler-learning/model"
	"github.com/lyquocnam/go-crawler-learning/repo"
	"github.com/lyquocnam/go-crawler-learning/services"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(model.Article{}, model.Link{})
	//fmt.Printf("\x0cOn %d/10", i)

	linkRepo := repo.NewLinkRepo(db)
	articleRepo := repo.NewArticleRepo(db)
	unFetchedLinkChan := make(chan *model.Link, 1)

	// Theo doi DB -> get link
	timeDelay, _ := strconv.Atoi(os.Getenv("TIME_DELAY"))
	watcher := services.NewWatcher(linkRepo, unFetchedLinkChan)
	go watcher.Watch(timeDelay) // delay 5s

	// fetch article khi co link
	articleChan := make(chan *model.Article)
	fetcher := services.NewFetcher(unFetchedLinkChan, articleChan)
	go fetcher.Listen()

	// update lai thong tin sau khi fetch
	updatedArticleChan := make(chan *model.Article)
	updater := services.NewUpdater(articleChan, updatedArticleChan, linkRepo)
	go updater.Listen()

	// luu article
	savedArticleChan := make(chan *model.Article)
	saver := services.NewSaver(updatedArticleChan, savedArticleChan, articleRepo)
	go saver.Listen()

	// thong ke
	for {
		select {
		case link := <-unFetchedLinkChan:
			fmt.Printf("link: %v \n", link)
		}
	}
}
