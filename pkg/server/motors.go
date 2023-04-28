package server

import (
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"main.go/pkg/models"
	// "../models"
)

func (s *Server) getMotor(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if motors, err := s.database.GetMotors(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"motors": motors,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMotorById(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if motor_id, err := strconv.Atoi(c.Param("motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if motor, err := s.database.GetMotorById(motor_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"motor": motor,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) editMotor(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	userStatus, _ := c.Get("Status")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else if userStatus == "observer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Недостаточно прав на выполнение действий",
		})
	} else {
		var motor models.Motor
		if err := c.ShouldBindJSON(&motor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование привода").Error(),
			})
		} else {
			if err := s.database.EditMotor(&motor); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Привод был успешно отредактирован",
				})
			}
		}
	}
}

func (s *Server) uploadImgMotor(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	userStatus, _ := c.Get("Status")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else if userStatus == "observer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Недостаточно прав на выполнение действий",
		})
	} else {
		if motor_id, err := strconv.Atoi(c.Param("motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			file, err := c.FormFile("file")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "No file is received",
					"error":   err.Error(),
				})
				return
			}

			extension := filepath.Ext(file.Filename)
			newFileName := uuid.New().String() + extension
			imgUrl := "/assets/motors/" + newFileName
			filePath := "." + imgUrl

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
					"error":   err.Error(),
				})
				return
			}

			if err := s.database.UploadImgMotor(motor_id, imgUrl); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Изображение успешно загружено",
				"img_url": imgUrl,
			})
		}
	}
}

func (s *Server) addMotor(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	userStatus, _ := c.Get("Status")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else if userStatus == "observer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Недостаточно прав на выполнение действий",
		})
	} else {
		var addRequest models.MotorAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление привода в компонентную базу").Error(),
			})
		} else {
			if id, err := s.database.AddMotor(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "Привод был успешно добавлен в компонентную базу",
				})
			}
		}
	}
}
