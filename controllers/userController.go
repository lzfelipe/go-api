package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lzfelipe/go-crud/initializers"
	model "github.com/lzfelipe/go-crud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateUser(c *gin.Context) {
	//read from body
	var body model.UserInput
	c.Bind(&body)

	//create user
	hash, err := HashPassword(body.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error while encrypting password",
		})
		return
	}

	user := model.User{Username: body.Username, Email: body.Email, Password: hash}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return
	c.JSON(200, gin.H{
		"user": user,
	})
}

func LoginUser(c *gin.Context)  {
	var body struct{
		Email string
		Password string
	}

	c.Bind(&body)

	var user model.User
	result := initializers.DB.Select("id", "email", "deleted_at", "password").Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"user": nil,
			"error": "Unable to login. Check your credentials.",
		})
		return
	}

	match := CheckPasswordHash(body.Password, user.Password)

	if !match {
		c.JSON(400, gin.H{
			"user": nil,
			"error": "Unable to login. Check your credentials.",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": user.ID,
		"deleted_at": user.DeletedAt,
		"username": user.Username,
		"email": user.Email,
	})
}

func ReadUser(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get the user data
	var result struct{
		gorm.Model
		Username string
		Email string
	}

	// initializers.DB.Select("ID", "username", "email", "CreatedAt", "UpdatedAt", "DeletedAt").Where("ID = ?", id).Find(&user)
	initializers.DB.Table("users").Select("ID", "username", "email", "created_at", "updated_at", "deleted_at").Where("id = ?", id).Scan(&result)

	if result.ID == 0 {
		c.JSON(400, gin.H{
			"user": nil,
			"error": "User not found",
		})
		return
	}

	//return
	c.JSON(200, gin.H{
		"user": result,
	})
}
