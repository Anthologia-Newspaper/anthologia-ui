package sources

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

func getUserId(c *gin.Context) (int32, error) {

	bearerToken := c.Request.Header.Get("Authorization")
	tokenString := ""

	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString = strings.Split(bearerToken, " ")[1]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
	}

	tokenPure, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := tokenPure.Claims.(jwt.MapClaims); ok && tokenPure.Valid {
		userID, _ := claims["user_id"].(float64)
		return int32(userID), nil
	} else {
		return 0, errors.New("Invalid token")
	}
}

type UserInformation struct {
	Email    string
	Username string
}

// * fetch database with userId get from the function above, return a different object with the Author name and Email
func getUserInformation() UserInformation{
	var userInfo UserInformation;

	userInfo.Email = "";
	userInfo.Username = "";

	return userInfo;
}

type Article struct {
	gorm.Model
	Id      int32
	UserId  int32
	Title   string
	Subtitle string
	Content string
	Topic string
	Draft bool
	Likes   pq.Int32Array `gorm:"type:integer[]"`
}
