package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Sensor represents a smart home sensor
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}

// SensorCreate represents the data needed to create a new sensor
type UserCreate struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

// SensorUpdate represents the data that can be updated for a sensor
type UserUpdate struct {
	Address string `json:"address"`
}

var users = []User{
	{ID: 1, Name: "Петя", Phone: "Петин телефон", Address: "Петин адрес", LastUpdated: time.Now(), CreatedAt: time.Now()},
	{ID: 2, Name: "Вася", Phone: "Васин телефон", Address: "Васин адрес", LastUpdated: time.Now(), CreatedAt: time.Now()},
}

// CreateUser handles POST /api/v2/user
func CreateUser(c *gin.Context) {
	// Здесь был бы код для записи в базу
	var userCreate UserCreate
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var new_user = User{len(users) + 1, userCreate.Name, userCreate.Phone, "DefAddress", time.Now(), time.Now()}
	users = append(users, new_user)

	c.JSON(http.StatusCreated, new_user)
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if id <= 0 || id > len(users) {

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, users[id-1])
}
