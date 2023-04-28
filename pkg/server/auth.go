package server

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
	// "../models"
)

func (s *Server) loginHandler(c *gin.Context) {
	fmt.Println("отрабатываю пост запрос на авторизацию")

	var loginRequest models.LoginRequest
	fmt.Println("запрос на авторизацию", loginRequest)
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("Incorrect authorization form data").Error()})
	} else {
		passwordHash := md5.New()
		passwordHash.Write([]byte(loginRequest.Password))
		passwordMd5 := hex.EncodeToString(passwordHash.Sum(nil))
		loginRequest.Password = passwordMd5

		log.Println(loginRequest)
		if id, status, err := s.database.LoginAuthorization(&loginRequest); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("Не правильные имя пользователя или пароль. Пожалуйста, проверьте ваши данные для авторизации.").Error()})
		} else {
			authToken := s.generateAuthToken(id, status)
			c.JSON(http.StatusOK, gin.H{
				"Message": "Успешная авторизация",
				"Token":   authToken,
			})
		}
	}
}

func (s *Server) registerHandler(c *gin.Context) {
	var registerRequest models.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("Incorrect registration form data").Error()})
	} else {
		passwordHash := md5.New()
		passwordHash.Write([]byte(registerRequest.Password))
		passwordMd5 := hex.EncodeToString(passwordHash.Sum(nil))
		registerRequest.Password = passwordMd5

		if id, err := s.database.RegisterUser(&registerRequest); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("Такой пользователь уже существует. Измените регистрационные данные.").Error()})
		} else {
			status := "observer"
			authToken := s.generateAuthToken(id, status)

			c.JSON(http.StatusOK, gin.H{
				"Message": "Успешная авторизация",
				"Token":   authToken,
			})
		}

	}
}
