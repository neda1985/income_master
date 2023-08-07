package models

type UserRegister struct {
	Name string `json:"name" biding:"required"`
	LastName string `json:"lastName" biding:"required"`

}
