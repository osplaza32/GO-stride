package models

import (
	"time"
)

type Colletion struct {
	ID        string `gorm:"primary_key;uuid"`
	CreatedAt time.Time
	Score float32
	UserID  string     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
	points []Points `gorm:"foreignkey:UserRefer"`
	Type string
	Description string


}