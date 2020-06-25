package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string
	PasswordHash string
	Email        string
	Texts        []Text
	Fragments    []Fragment
	Links        []Link
}

type Text struct {
	gorm.Model
	Title     string
	Body      string
	Source    EntryType
	Locator   string
	UserID    uint
	Fragments []Fragment
}

type Fragment struct {
	gorm.Model
	Body   string
	TextID uint
	Link   []*Link `gorm:"many2many:fragment_links"`
	UserID uint
}

type Link struct {
	gorm.Model
	Name      string
	Relations []Link      `gorm:"many2many:link_relations;association_jointable_foreignkey:relation_id"`
	Fragments []*Fragment `gorm:"many2many:fragment_links"`
	UserID    uint
}

type EntryType int

const (
	InvalidSource EntryType = iota
	ManualEntry
	Website
	LocalFile
)
