package controllers

import (
	"belajar-golang-api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateCategoryInput struct {
	Name   string `json:"name" binding:"required"`
}

// type ErrorMsg struct {
// 	Field   string `json:"field"`
// 	Message string `json:"message"`
// }


// func GetErrorMsg(fe validator.FieldError) string {
// 	switch fe.Tag() {
// 	case "required":
// 		return "This field is required"
// 	}
// 	return "Unknown error"
// }

func FindCategories(c *gin.Context) {

	var categories []models.Category
	models.DB.Select("id", "name").Find(&categories)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Categories",
		"data":    categories,
	})
}

func StoreCategory(c *gin.Context) {
	var input ValidateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	category := models.Category{
		Name:   input.Name,
	}
	models.DB.Create(&category)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Category Created Successfully",
		"data":    category,
	})
}

func FindCategoryById(c *gin.Context) {
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Category By ID : " + c.Param("id"),
		"data":    category,
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input ValidateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	models.DB.Model(&category).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Category Updated Successfully",
		"data":    category,
	})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&category)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Category Deleted Successfully",
	})
}