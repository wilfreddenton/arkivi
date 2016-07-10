package main

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

// User
type User struct {
	gorm.Model
	Username string
	Password string `json:"-"`
	Admin    bool
	Settings Settings
}

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSendJson struct {
	CreatedAt time.Time
	Username  string
	Admin     bool
	NumImages int
	Settings  Settings
}

// Settings
type Settings struct {
	gorm.Model
	UserID       uint
	Camera       string
	Film         string
	Public       bool
	Registration bool // admin only setting
}

// Image
type Image struct {
	gorm.Model
	MonthID     uint
	UserID      uint
	Title       string
	Description string `gorm:"size:1500"`
	TakenAt     *time.Time
	Camera      string
	Film        string
	Tags        []Tag `gorm:"many2many:image_tags;"`
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

func (i *Image) GetPaths() []string {
	var paths []string
	old := "/static/"
	new := "assets/"
	if i.ThumbUrl != "" {
		paths = append(paths, strings.Replace(i.ThumbUrl, old, new, 1))
	}
	if i.SmallUrl != "" {
		paths = append(paths, strings.Replace(i.SmallUrl, old, new, 1))
	}
	if i.MediumUrl != "" {
		paths = append(paths, strings.Replace(i.MediumUrl, old, new, 1))
	}
	if i.LargeUrl != "" {
		paths = append(paths, strings.Replace(i.LargeUrl, old, new, 1))
	}
	if i.Url != "" {
		paths = append(paths, strings.Replace(i.Url, old, new, 1))
	}
	return paths
}

type ImageMini struct {
	ID int
}

type ImageJson struct {
	ID          int
	Title       string
	TakenAt     string
	Description string
	Camera      string
	Film        string
	Tags        []TagJson
	Published   bool
}

// Tag
type Tag struct {
	gorm.Model
	Name   string
	Images []*Image `gorm:"many2many:image_tags"`
}

type TagJson struct {
	Name string
}

type TagCountJson struct {
	Name  string
	Count int
}
type TagMini struct {
	ImageID int
}

// Action
type Action struct {
	IDs   []int
	Value interface{}
}

type ActionTags struct {
	IDs   []int
	Value []TagJson
}

// Month
type Month struct {
	gorm.Model
	Month     string
	Year      int
	NumImages int
}

// Year
type Year struct {
	Year   int
	Months []Month
}

// misc
type count struct {
	Count int
}

type UrlParam struct {
	Name    string
	Value   string
	IsFirst bool
	IsLast  bool
}
