package templates

// Record corresponds with a line from repos.csv (or source)
type Record struct {
	URL      string `csv:"url"`
	Category string `csv:"category"`
}

// Store are all records on source
type Store []Record

// Repository maps a repo information
type Repository struct {
	Name        string
	Description string
	// UpdatedAt   interface{} // date?
	Stars int
}
