package controllers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateArticleInput struct {
	CONTENT  string `json:"content" binding:"required,min=200"`
	TITLE    string `json:"title" binding:"required,min=20"`
	CATEGORY string `json:"category" binding:"required,min=3"`
	STATUS   string `json:"status" binding:"required"`
}

type UpdateArticleInput struct {
	CONTENT  string `json:"content" binding:"required,min=200"`
	TITLE    string `json:"title" binding:"required,min=20"`
	CATEGORY string `json:"category" binding:"required,min=3"`
	STATUS   string `json:"status" binding:"required"`
}

func Findarticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var article []models.Article
	db.Find(&article)
	if len(article) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": article})
	}
}

func Addarticle(c *gin.Context) {
	var input CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{CONTENT: input.CONTENT, TITLE: input.TITLE, CATEGORY: input.CATEGORY, STATUS: input.STATUS}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&article)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func FindarticlebyId(c *gin.Context) {
	var article models.Article
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func Updatearticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var article models.Article
	if err := db.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	var input UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var UpdateInput models.Article
	UpdateInput.CATEGORY = input.CATEGORY
	UpdateInput.CONTENT = input.CONTENT
	UpdateInput.STATUS = input.STATUS
	UpdateInput.TITLE = input.TITLE
	db.Model(&article).Updates(UpdateInput)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func Deletearticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var article models.Article
	if err := db.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	db.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Findbylimit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))
	var article []models.Article
	if err := db.Limit(limit).Offset(offset).Find(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": article})
}
