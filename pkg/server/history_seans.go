package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
	// "../models"
)

func (s *Server) getAllHistorySeans(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if history_seans, err := s.database.GetAllHistorySeans(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"history_seans": history_seans,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getHistorySeansById(c *gin.Context) {
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
		} else if histories, err := s.database.GetHistorySeansById(id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"history_seans": histories,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) editHistorySeans(c *gin.Context) {
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
		var history models.HistorySeans
		if err := c.ShouldBindJSON(&history); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование истории сеанса механизма").Error(),
			})
		} else {
			if err := s.database.EditHistorySeans(&history); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "История сеанса механизма была успешна отредактирована",
				})
			}
		}
	}
}

func (s *Server) addNewHistorySeans(c *gin.Context) {
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
		var addRequest models.HistorySeansAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление истории сенаса наблюдений за механизмом").Error(),
			})
		} else {
			if id, err := s.database.AddNewHistorySeans(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "История сеанса наблюдений за механизмом была успешно добавлена",
				})
			}
		}
	}
}

func (s *Server) deleteHistorySeans(c *gin.Context) {
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
		} else if err := s.database.DeleteHistorySeans(id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "История сеанса наблюдений за механизмом была успешно удалена",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}
