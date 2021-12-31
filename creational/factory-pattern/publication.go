package main

const MIN_PAGE_LENGTH = 15

type iPublication interface {
	SetName(string)
	SetPages(int)
	SetPublisher(string)
	GetName() string
	GetPages() int
	GetPublisher() string
}

type Publication struct {
	name      string
	publisher string
	pages     int
}

func (p *Publication) SetName(name string) {
	p.name = name
}

func (p *Publication) SetPages(pages int) {
	p.pages = pages
}

func (p *Publication) SetPublisher(publisher string) {
	p.publisher = publisher
}

func (p *Publication) GetName() string {
	return p.name
}

func (p *Publication) GetPages() int {
	return p.pages
}

func (p *Publication) GetPublisher() string {
	return p.publisher
}
