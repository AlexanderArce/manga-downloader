package grabber

import "strings"

// Chapter represents a manga chapter
type Chapter struct {
	// Title is the chapter title
	Title string
	// Number is the chapter number
	// TODO: Why is this a float64? a chapter number is always an integer
	Number float64
	// PagesCount is the number of pages of the chapter
	PagesCount int64
	// Pages is the list of pages of the chapter
	Pages Pages
	// Language is the chapter language
	Language string
}

// Chapters is a slice of Chapter
type Chapters []Chapter

// Page represents a manga page
type Page struct {
	// Number is the page number
	Number int64
	// URL is the page URL
	URL string
}

// Pages is a slice of Page
type Pages []Page

// GetNumber returns the chapter number
func (c Chapter) GetNumber() float64 {
	return c.Number
}

// GetTitle returns the chapter title
func (c Chapter) GetTitle() string {
	return strings.TrimSpace(c.Title)
}
