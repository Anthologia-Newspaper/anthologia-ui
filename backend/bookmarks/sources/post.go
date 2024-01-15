package sources

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type TitleArticle struct {
	Id    int32
	Title string
}

func addIfNotPresent(arr pq.Int32Array, key int32) pq.Int32Array {
	for _, val := range arr {
		if val == key {
			return arr
		}
	}

	return append(arr, key)
}

func AddBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.Bind(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid arguments"})
		return
	}

	if bookmark.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookmark must have a title"})
		return
	}
	bookmark.UserId = userId
	bookmark.Articles = pq.Int32Array{}

	result := db.Create(&bookmark)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating bookmark"})
		return
	}
	c.JSON(http.StatusCreated, bookmark)
}

func AddArticleInBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookmarkId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "Bookmark id could not be retrieved"})
		return
	}
	articleId, err := strconv.ParseInt(c.Param("id-article"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "Article id could not be retrieved"})
		return
	}

	result := db.Where(Bookmark{Id: int32(bookmarkId), UserId: userId}).Find(&bookmark)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interacting with database"})
		return
	} else if bookmark.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}
	bookmark.Articles = addIfNotPresent(bookmark.Articles, int32(articleId))
	db.Save(&bookmark)
	c.JSON(http.StatusOK, bookmark)
}
