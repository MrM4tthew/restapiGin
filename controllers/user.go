package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"restapiGin/models"
)

type RegisterUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GET /users
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
func GetUser(c *gin.Context) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Register(c *gin.Context) {
	// Validate input
	var input RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Bcrypt password
	hash, _ := HashPassword(input.Password)

	// Create user
	user := models.User{Name: input.Name, Username: input.Username, Email: input.Email, Password: hash}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}

func Login(c *gin.Context) {
	// Validate input
	var input LoginUserInput
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	// Check input first
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// Do this if input is successfully executed
		// Check if user is in database
		if err2 := db.Where("email = ?", input.Email).First(&user).Error; err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		} else {
			match := CheckPasswordHash(input.Password, user.Password)
			if match == false {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong password"})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": "user is in"})
			}
		}

	}

}

// Password Bcrypt function
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
