package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Friends  []User
	LoggedIn bool
}
