package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	users := make([]User, 0)

	router := gin.Default()

	router.GET("/users", getUsersHandler(users))
	router.GET("/users/:id", getUserHandler(users))
	router.POST("/users", createUserHandler(users))
	router.PATCH("users/:id", updateUserHandler(users))
	router.DELETE("/users/:id", deleteUserHandler(users))

	router.Run()
}

func getUsersHandler(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	}
}

func getUserHandler(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, user := range users {
			if user.Id == c.Param("id") {
				c.JSON(http.StatusOK, user)
			}
		}
	}
}

func createUserHandler(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uuid := uuid.NewString()
		newUser := User{Id: uuid, Name: user.Name, Password: user.Password, Email: user.Email}
		users = append(users, newUser)

		c.JSON(http.StatusCreated, gin.H{
			"id":       newUser.Id,
			"name":     newUser.Name,
			"password": newUser.Password,
			"email":    newUser.Email,
		})
	}
}

func updateUserHandler(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, user := range users {
			if user.Id == c.Param("id") {
				users[i].Name = updatedUser.Name
				users[i].Password = updatedUser.Password
				users[i].Email = updatedUser.Email

				c.JSON(http.StatusOK, gin.H{
					"id":       users[i].Id,
					"name":     users[i].Name,
					"password": users[i].Password,
					"email":    users[i].Email,
				})
			}
		}
	}
}

func deleteUserHandler(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := range users {
			if users[i].Id == c.Param("id") {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusNoContent, nil)
			}
		}
	}
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
