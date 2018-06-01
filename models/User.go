package models

import "time"

type User struct {
	ID        string `gorm:"primary_key;uuid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time

	points []Points `gorm:"foreignkey:UserRefer"`

}

func (s User) IsEmpty() bool {
	return s.Equal(User{})
}
func (s User) Equal(o User) bool {

	if(s.ID != o.ID) { return false }
	if(s.FirstName != o.FirstName) { return false }
	if(s.LastName != o.LastName) { return false }
	if(s.Email != o.Email) { return false }
	if(s.Password != o.Password) { return false }
	return true
}