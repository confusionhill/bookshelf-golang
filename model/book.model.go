package model

import "time"

type BookModel struct {
	Id        string
	Name      string
	Year      int32
	Author    string
	Summary   string
	Publisher string
	PageCount int32
	ReadPage  int32
	Reading   bool
	Finished  bool
	InsertAt  time.Time
	UpdateAt  time.Time
}

type BookInput struct {
	Id        string
	Name      string
	Year      int32
	Author    string
	Summary   string
	Publisher string
	PageCount int32
	ReadPage  int32
	Reading   bool
}

// func convertToBook() *BookModel {

// }
