package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"restapiGin/models"
	"restapiGin/service"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type UpdateCategoryInput struct {
	Name string `json:"name"`
}

// GET /categories
func FindCategories(c *gin.Context) {
	tokenAuth, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 1", "err": err})
	}
	userId, err := service.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 2", "err": err})
	}
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Where("user_id = ?", userId).Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GET /categories/:id
func FindCategory(c *gin.Context) {
	tokenAuth, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 1", "err": err})
	}
	userId, err := service.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 2", "err": err})
	}

	db := c.MustGet("db").(*gorm.DB)
	var category models.Category

	if categoryErr := db.Where("user_id = ? AND id = ? ", userId, c.Param("id")).Find(&category).Error; categoryErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// POST /categories
func CreateCategory(c *gin.Context) {
	// Validate input
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenAuth, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 1", "err": err})
	}
	userId, err := service.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized 2", "err": userId})
	}

	// Convert id from string to uuid type
	userUuid, _ := uuid.FromString(userId)
	//userUuid2, _ := uuid.FromString("4ea0b2d3-df62-4542-9f36-5d62ae6dc229")

	// Create category
	category := models.Category{ Name: input.Name, UserId: userUuid}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}
