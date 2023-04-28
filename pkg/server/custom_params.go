package server

import (
	"errors"
	"net/http"
	"strconv"

	// "../models"
	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
)

func (s *Server) getCustomParams(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if params, err := s.database.GetCustomParams(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"custom_params": params,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) addCustomParam(c *gin.Context) {
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
		var addRequest models.CustomParamAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление дополнительного параметра").Error(),
			})
		} else {
			if id, err := s.database.AddCustomParam(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "Дополнительный параметр был успешно добавлен",
				})
			}
		}
	}
}

func (s *Server) editCustomParam(c *gin.Context) {
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
		var param models.CustomParam
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование дополнительного параметра").Error(),
			})
		} else {
			if err := s.database.EditCustomParam(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Дополнительный параметр был успешно отредактирован",
				})
			}
		}
	}
}

func (s *Server) deleteCustomParam(c *gin.Context) {
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
		if param_id, err := strconv.Atoi(c.Param("custom_param_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteCustomParam(param_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Дополнительный параметр был успешно удален",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getCustomParamById(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if param_id, err := strconv.Atoi(c.Param("custom_param_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if param, err := s.database.GetParamById(param_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"custom_param": param,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}
