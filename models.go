package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Image
type Image struct {
	gorm.Model
	Title       string
	Description string `gorm:"size:1500"`
	TakenAt     time.Time
	Camera      string
	Film        string
	Tags        []Tag
	Views       int
	Name        string
	Ext         string
	Width       int
	Height      int
	ThumbUrl    string
	SmallUrl    string
	MediumUrl   string
	LargeUrl    string
	Url         string
	Published   bool
}

func (i *Image) Save() {
	DB.Create(i)
}

// Tag
type Tag struct {
	ID      int
	ImageID int `gorm:"index"`
	Tag     string
}
