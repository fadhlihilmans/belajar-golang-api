package controllers

import (
	"belajar-golang-api/helpers"
	"belajar-golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

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
		errors := helpers.FormatValidationError(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name: input.Name,
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
		errors := helpers.FormatValidationError(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// if err := models.DB.Delete(&category).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"success": false, 
	// 		"message": "Cannot delete category: it is still used by some posts",
	// 	})
	// 	return
	// }
	if err := models.DB.Delete(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
					"success": false, 
					"message": "Gagal menghapus data",
					"debug_error": err.Error(), 
			})
			return
	}

	models.DB.Delete(&category)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Category Deleted Successfully",
	})
}