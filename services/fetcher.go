package services

import (
	"go_crawler/model"
	"log"
	"math/rand"
	"time"
)

type fetcher struct {
	in  chan *model.Link
	out chan *model.Article
}

func NewFetcher(in chan *model.Link, out chan *model.Article) *fetcher {
	return &fetcher{in: in, out: out}
}

func (f *fetcher) Listen() {
	for {
		select {
		case link := <-f.in:
			go f.fetch(link)
		}
	}
}

func (f *fetcher) fetch(link *model.Link) {
	log.Printf("fetching: %s \n", link.URL)
	// simulate fetching time random from 5 seconds
	time.Sleep(time.Duration(time.Duration(rand.Intn(5)) * time.Second))

	article := &model.Article{
		Title:   "Mỹ, Canada sử dụng chó nghiệp vụ để ngăn chặn tả heo châu Phi",
		Content: "(TBKTSG Online) - Mỹ và Canada đang tăng cường sử dụng các đội chó nghiệp vụ ở các cảng biển, sân bay để kiểm tra hàng hóa, giúp phát hiện sớm các sản phẩm thịt heo nhập lậu có thể nhiễm virus dịch tả heo châu Phi (ASF).",
		Author:  "Lê Linh",
		URL:     link.URL,
	}
	log.Printf("fetched article: %s \n", article.Title)

	f.out <- article
}
