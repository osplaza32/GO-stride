package models

import (
	"time"
)

type Colletion struct {
	ID        string `gorm:"primary_key;uuid"`
	CreatedAt time.Time
	UserID  string     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
	Points []Points `gorm:"foreignkey:ColletionRefer"`
	Proyect string `gorm:"default:'Stride'"`
	Type string
	Description string


}
func (s Colletion) IsEmpty() bool {
	return s.Equal(Colletion{})
}
func (s Colletion) Equal(o Colletion) bool {

	if(s.ID != o.ID) { return false }
	if(s.Proyect != o.Proyect) { return false }
	if(s.Type != o.Type) { return false }
	if(s.UserID != o.UserID) { return false }
	//if(s.Password != o.Password) { return false }
	return true
}