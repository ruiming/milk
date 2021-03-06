package models

import "time"

// Tag Model
type Tag struct {
	Id        int
	Name      string
	Posts     []*Post   `orm:"reverse(many)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
