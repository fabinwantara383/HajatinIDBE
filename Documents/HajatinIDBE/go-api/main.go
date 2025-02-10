package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struktur data untuk response JSON
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Data dummy
var users = []User{
	{ID: 1, Name: "Alice", Age: 25},
	{ID: 2, Name: "Bob", Age: 30},
}

func main() {
	r := gin.Default()

	// Endpoint GET untuk mengambil semua user
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Endpoint POST untuk menambah user baru
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	// Jalankan server di port 8080
	r.Run(":8080")
}
