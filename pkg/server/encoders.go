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

func (s *Server) getEncoders(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if encoders, err := s.database.GetEncoders(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"encoders": encoders,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getEncoder(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if enc_id, err := strconv.Atoi(c.Param("enc_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if encoder, err := s.database.GetEncoder(enc_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"encoder": encoder,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) editEncoder(c *gin.Context) {
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
		var encoder models.EncoderEditRequest
		if err := c.ShouldBindJSON(&encoder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование энкодера").Error(),
			})
		} else {
			if err := s.database.EditEncoder(&encoder); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Энкодер был успешно отредактирован",
				})
			}
		}
	}
}

func (s *Server) uploadImgEncoder(c *gin.Context) {
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
		if enc_id, err := strconv.Atoi(c.Param("enc_id")); err != nil {
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
			imgUrl := "/assets/encoders/" + newFileName
			filePath := "." + imgUrl

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
					"error":   err.Error(),
				})
				return
			}

			if err := s.database.UploadImgEncoder(enc_id, imgUrl); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Your file has been successfully uploaded.",
				"img_url": imgUrl,
			})
		}
	}
}

func (s *Server) addEncoder(c *gin.Context) {
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
		var addRequest models.EncoderAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление энкодера в компонентную базу").Error(),
			})
		} else {
			if id, err := s.database.AddEncoder(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "Энкодер был успешно добавлен в компонентную базу",
				})
			}
		}
	}
}
