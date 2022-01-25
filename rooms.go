package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name      string    `json:"name"`
	CreatedBy User      `json:"created-by"`
	Users     []User    `json:"users"`
	Messages  []Message `json:"messages"`
}

func GetAllRooms(c *gin.Context) {
	room := Room{
		Name: "Gaming",
		CreatedBy: User{
			Name:     "Bob",
			LoggedIn: true,
		},
		Users: []User{
			{
				Name:     "Bob",
				Bio:      "",
				Image:    "",
				Friends:  []User{},
				LoggedIn: true,
			},
		},
		Messages: []Message{},
	}

	c.JSON(http.StatusOK, room)
}

func GetOneRoom(c *gin.Context) {
}

func CreateRoom(c *gin.Context) {

	room := Room{}
	c.BindJSON(room)

	c.JSON(http.StatusOK, room)
}

func DeleteRoom(c *gin.Context) {
}
