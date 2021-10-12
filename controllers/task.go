package controllers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"restapiGin/models"
	"restapiGin/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Deadline   string    `json:"deadline"`
	CategoryId uuid.UUID `json:"category_id"`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	tokenAuth, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 1", "err": err})
	}
	userId, err := service.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 2", "err": err})
	}
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Where("user_id = ?", userId).Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	tokenAuth, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 1", "err": err})
	}
	userId, err := service.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 2", "err": userId})
	}
	userUuid, _ := uuid.FromString(userId)

	var categories models.Category

	if categoryErr := db.Where("id = ? AND user_id = ? ", input.CategoryId, userUuid).First(&categories).Error; categoryErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!!"})
		return
	}

	// Create task
	task := models.Task{AssingedTo: input.AssingedTo, Task: input.Task, Deadline: deadline, UserId: userUuid, CategoryId: input.CategoryId}

	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
