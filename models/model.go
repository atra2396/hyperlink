package model

type Text struct {
	ID      uint
	Title   string
	Body    string
	Source  EntryType
	Locator string
}

type Fragment struct {
	ID         uint
	Body       string
	ParentText Text
	ParentId   uint
	Link       []Link
}

type Link struct {
	ID        uint
	Name      string
	Relations []Link
}

type EntryType int

const (
	InvalidSource EntryType = iota
	ManualEntry
	Website
	LocalFile
)
