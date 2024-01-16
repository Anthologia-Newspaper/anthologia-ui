package sources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type LikesResponse struct {
	Amount   int
	Accounts pq.Int32Array `gorm:"type:integer[]"`
}

func addIfNotPresent(arr pq.Int32Array, key int32) pq.Int32Array {
	for _, val := range arr {
		if val == key {
			return arr
		}
	}
	return append(arr, key)
}

func GetAllArticles(c *gin.Context, db *gorm.DB) {
	articles := new([]Article)

	result := db.Where(Article{}).Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context, db *gorm.DB) {
	article := new(Article)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "article id could not be retrieved"})
		return
	}

	result := db.Where(Article{Id: int32(id)}).Find(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if article.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetMyArticles(c *gin.Context, db *gorm.DB) {
	articles := new([]Article)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where(Article{UserId: userId}).Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetMyLikedArticles(c *gin.Context, db *gorm.DB) {
	var articles []Article

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	var filteredArticles []Article
	filteredArticles = []Article{}
	for _, article := range articles {
		for _, like := range article.Likes {
			if like == userId {
				filteredArticles = append(filteredArticles, article)
				break
			}
		}
	}

	c.JSON(http.StatusOK, filteredArticles)
}

func GetLikesInfo(c *gin.Context, db *gorm.DB) {
	article := new(Article)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	result := db.Where(Article{Id: int32(id)}).Find(&article)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if article.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	c.JSON(http.StatusOK, LikesResponse{len(article.Likes), article.Likes})
}


func GetTopArticles(c *gin.Context, db *gorm.DB) {
    var articles []Article

   
    result := db.Order("array_length(BU, 1) DESC").Limit(10).Find(&articles)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
        return
    }

    c.JSON(http.StatusOK, articles)
}