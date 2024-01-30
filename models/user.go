package models

type User struct {
	ID       int    `json:"id" gorm:"primary key;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gomr:"unique"`
	Password []byte `json:"-"`
}
