package services

type crawler struct {
	out chan string
	in  chan string
}
