package articles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

/*
DeleteAllArticles deletes all the articles of the connected user
*/
func DeleteAllArticles(c *gin.Context, db *gorm.DB) {

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where(Article{UserId: userId}).Delete(&Article{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "all articles have been successfully deleted"})
}

/*
DeleteArticle deletes one of the articles of the connected user
*/
func DeleteArticle(c *gin.Context, db *gorm.DB) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"delete": "article id could not be retrieved"})
		return
	}
	result := db.Where(Article{UserId: userId, Id: uint(id)}).Delete(&Article{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found or was not created by the current user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "article has been successfully deleted"})
}

/*
RemoveLike removes the like of a given post from the connected user
*/
func RemoveLike(c *gin.Context, db *gorm.DB) {
	article := new(Article)
	i := 0

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	result := db.Where(Article{Id: uint(id)}).Find(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if article.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	for _, value := range article.Likes {
		if value != userId {
			article.Likes[i] = value
			i++
		}
	}
	article.Likes = article.Likes[:i]
	db.Save(&article)
	c.JSON(http.StatusOK, article)
}
