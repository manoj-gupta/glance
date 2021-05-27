package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// needed for "postgres" driver
	_ "github.com/lib/pq"

	"github.com/manoj-gupta/glance/internal/db"
	"github.com/manoj-gupta/glance/internal/models"
	"github.com/manoj-gupta/glance/internal/utils"
)

func getSecretKey() string {
	return utils.GetEnv("JWT_SECRET_KEY", "secret")
}

func getCookieName() string {
	return utils.GetEnv("JWT_COOKIE_NAME", "cookie")
}

// Register .. Register route handler
func Register(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "bad input data"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	fmt.Println(user)

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotAcceptable,
			gin.H{"error": "user exists"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Login .. Login a user
func Login(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "bad input data"})
		return
	}

	var user models.User

	db.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "incorrect password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(getSecretKey()))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not login"})
		return
	}

	c.SetCookie(getCookieName(), token, 60*60*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message:": "success"})
}

// User .. Retrieve logged-in user
func User(c *gin.Context) {
	cookie, err := c.Cookie(getCookieName())
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"message": "Not logged in"})
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"message": "unauthenticated"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	db.DB.Where("id = ?", claims.Issuer).First(&user)

	c.JSON(http.StatusOK, user)
}

// Logout .. logout user
func Logout(c *gin.Context) {
	// set cookie time in past
	c.SetCookie(getCookieName(), "", -60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message:": "success"})
}
