package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"main.go/pkg/models"
	// "../models"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAllHistoryCustomParams(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if history_custom_params, err := s.database.GetAllHistoryCustomParams(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"history_custom_params": history_custom_params,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getHistoryCustomParamsById(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if id, err := strconv.Atoi(c.Param("id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if histories, err := s.database.GetHistoryCustomParamsById(id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"history_custom_params": histories,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) editHistoryCustomParams(c *gin.Context) {
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
		var history models.HistoryCustomParams
		if err := c.ShouldBindJSON(&history); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование истории дополнительных наблюдений").Error(),
			})
		} else {
			if err := s.database.EditHistoryCustomParams(&history); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "История дополнительных наблюдений была успешна отредактирована",
				})
			}
		}
	}
}

func (s *Server) addNewHistoryCustomParams(c *gin.Context) {
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
		var addRequest models.HistoryCustomParamsAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление истории дополнительных наблюдений").Error(),
			})
		} else {
			if id, err := s.database.AddNewHistoryCustomParams(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "История дополнительных наблюдений была успешно добавлена",
				})
			}
		}
	}
}

func (s *Server) deleteHistoryCustomParams(c *gin.Context) {
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
		if id, err := strconv.Atoi(c.Param("id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteHistoryCustomParams(id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "История дополнительных наблюдений была успешно удален",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}
