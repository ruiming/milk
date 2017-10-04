package models

import "time"

// Comment Model
type Comment struct {
	Id        int
	Name      string
	Email     string
	Website   string
	Post      *Post     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
