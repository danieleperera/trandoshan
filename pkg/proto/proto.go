package proto

import "time"

const (
	// URLTodoSubject is the subject used when an URL is schedule for crawling
	URLTodoSubject = "url.todo"
	// URLFoundSubject is the subject used when an URL is extracted from resource
	URLFoundSubject = "url.found"
	// NewResourceSubject is the subject used when a new resource has been crawled
	NewResourceSubject = "resource.new"
)

// URLTodoMsg represent an URL to crawl
type URLTodoMsg struct {
	URL string `json:"url"`
}

func (msg *URLTodoMsg) Subject() string {
	return URLTodoSubject
}

// URLFoundMsg represent a found URL
type URLFoundMsg struct {
	URL string `json:"url"`
}

func (msg *URLFoundMsg) Subject() string {
	return URLFoundSubject
}

type NewResourceMsg struct {
	URL  string `json:"url"`
	Body string `json:"body"`
}

func (msg *NewResourceMsg) Subject() string {
	return NewResourceSubject
}

// ResourceDto represent a resource as given by the API
type ResourceDto struct {
	URL   string    `json:"url"`
	Body  string    `json:"body"`
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
}
