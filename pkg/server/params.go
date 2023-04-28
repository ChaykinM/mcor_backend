package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"main.go/pkg/models"
	// "../models"
)

func (s *Server) getParams(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if params, err := s.database.GetParams(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"params": params,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) addParam(c *gin.Context) {
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
		var addRequest models.ParamAddRequest
		if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление параметра").Error(),
			})
		} else {
			if id, err := s.database.AddParam(&addRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"id":    id,
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      id,
					"message": "Параметр был успешно добавлен",
				})
			}
		}
	}
}

func (s *Server) editParam(c *gin.Context) {
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
		var param models.Param
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование параметра").Error(),
			})
		} else {
			if err := s.database.EditParam(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Параметр был успешно отредактирован",
				})
			}
		}
	}
}

func (s *Server) deleteParam(c *gin.Context) {
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
		if param_id, err := strconv.Atoi(c.Param("param_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteParam(param_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Параметр был успешно удален",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getParamById(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if param_id, err := strconv.Atoi(c.Param("param_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if param, err := s.database.GetParamById(param_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"param": param,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}
