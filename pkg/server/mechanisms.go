package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"main.go/pkg/kinematics"

	"main.go/pkg/models"
	// "../kinematics"
	// "../models"
	"github.com/gin-gonic/gin"
)

func (s *Server) getAllMechanisms(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mechanisms, err := s.database.GetAllMechanisms(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mechanisms": mechanisms,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechanism(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mechanism, err := s.database.GetMechanism(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mechanism": mechanism,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

}

func (s *Server) getAllMechanismsInfo(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mechanisms, err := s.database.GetAllMechanismsInfo(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mechanisms": mechanisms,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechanismInfo(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mechanism, err := s.database.GetMechanismInfo(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mechanism": mechanism,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) editMechanismInfo(c *gin.Context) {
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
		var mechInfo models.MechInfoEditRequest
		if err := c.ShouldBindJSON(&mechInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование механизма").Error(),
			})
		} else {
			if err := s.database.EditMechanismInfo(&mechInfo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Механизм был успешно отредактирован",
				})
			}
		}
	}

}

func (s *Server) getMechAllConfigs(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if configs, err := s.database.GetMechAllConfigs(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"configs": configs,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechConfig(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config_id, err := strconv.Atoi(c.Param("config_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config, err := s.database.GetMechConfig(mech_id, config_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"config": config,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getCurrentMechConfig(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config, err := s.database.GetCurrentMechConfig(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"configs": config,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getStandardMechConfig(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config, err := s.database.GetStandardMechConfig(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"configs": config,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) setActiveMechConfig(c *gin.Context) {
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
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config_id, err := strconv.Atoi(c.Param("config_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			if err = s.database.SetActiveNechConfig(mech_id, config_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Конфигурация механизма была успешно установлена активной",
				})
			}
		}
	}
}

func (s *Server) addMechConfig(c *gin.Context) {
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
		var mechConfig *models.MechConfigAddRequest
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&mechConfig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование конфигурации механизма").Error(),
			})
		} else if mech_id != mechConfig.Mech_id {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование конфигурации механизма. Параметр идентификатора механизма не совпадает со значением в запросе на редактирование.").Error(),
			})
		} else {
			if id, err := s.database.AddMechConfig(mechConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message":   "Конфигурация механизма была успешно добавлена",
					"config_id": id,
				})
			}
		}
	}
}

func (s *Server) editMechConfig(c *gin.Context) {
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
		var mechConfig *models.MechConfigEditRequest
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if config_id, err := strconv.Atoi(c.Param("config_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&mechConfig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование конфигурации механизма").Error(),
			})
		} else if mech_id != mechConfig.Mech_id || config_id != mechConfig.Id {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование конфигурации механизма. Параметры идентификаторов механизма и его конфигурации не совпадают со значениями в запросе на редактирование.").Error(),
			})
		} else {
			if err := s.database.EditMechConfig(mechConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Конфигурация механизма была успешно отредактирован",
				})
			}
		}
	}
}

func (s *Server) getMechMotors(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if motors, err := s.database.GetMechMotors(mech_id); err == nil {
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

func (s *Server) getMechMotor(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if motor, err := s.database.GetMechMotor(mech_id, mech_motor_id); err == nil {
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

func (s *Server) editMechMotor(c *gin.Context) {
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
		var mech models.MechMotorEditRequest
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if _, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&mech); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на редактирование привода механизма").Error(),
			})
		} else if err := s.database.EditMechMotor(&mech); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Привод механизма был успешно отредактирован",
			})
		}
	}

}

func (s *Server) addMechMotor(c *gin.Context) {
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
		var addRequest models.MechMotorAddRequest
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_id != addRequest.MechId {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("Не корректный запрос на добавление привода механизма").Error(),
			})
		} else if id, err := s.database.AddMechMotor(&addRequest); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"id":      id, // просто для теста пока так
				"message": "Новый привод для манипуляционного механизма был успешно добавлен",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) deleteMechMotor(c *gin.Context) {
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
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteMechMotor(mech_motor_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Привод манипуляционного механизма был успешно удалён",
			})
		}
	}
}

func (s *Server) getMechMotorsParamsAll(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if motorParams, err := s.database.GetMechMotorsParamsAll(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mech_id":     mech_id,
				"motorParams": motorParams,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechMotorParams(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if params, err := s.database.GetMechMotorParams(mech_motor_id); err == nil {
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

func (s *Server) addMechMotorParam(c *gin.Context) {
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
		var addRequest models.MechMotorParamAddRequest
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if id, err := s.database.AddMechMotorParam(mech_motor_id, addRequest.ParamId); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"id":      id,
				"message": "Новый параметр для наблюдения привода был успешно добавлен",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) deleteMechMotorParam(c *gin.Context) {
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
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if _, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_param_id, err := strconv.Atoi(c.Param("mech_motor_param_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteMechMotorParam(mech_motor_param_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Параметр для наблюдений работы привода механизма был успешно удалён",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechMotorEncodersAll(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if allMechMotorEncoders, err := s.database.GetMechMotorEncodersAll(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mech_id":              mech_id,
				"allMechMotorEncoders": allMechMotorEncoders,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechMotorEncoders(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mechMotorEncoders, err := s.database.GetMechMotorEncoders(mech_motor_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"mech_motor_encoders": mechMotorEncoders,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) addMechMotorEncoders(c *gin.Context) {
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
		var addRequest models.MechMotorEncoderAddRequest
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_id, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := c.ShouldBindJSON(&addRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if id, err := s.database.AddMechMotorEncoder(mech_motor_id, addRequest.EncoderId); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"id":      id, // просто для теста пока так
				"message": "Новый датчик обратной связи для привода механизма успешно добавлен",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) deleteMechMotorEncoder(c *gin.Context) {
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
		if _, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if _, err := strconv.Atoi(c.Param("mech_motor_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if mech_motor_enc_id, err := strconv.Atoi(c.Param("mech_motor_enc_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if err := s.database.DeleteMechMotorEncoder(mech_motor_enc_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Датчик обратной связи для привода механизма был успешно удалён",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechHistorySeans(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if histories, err := s.database.GetMechHistorySeans(mech_id); err == nil {
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

func (s *Server) getMechTrajectories(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if trajectories, err := s.database.GetMechTrajectories(mech_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"trajectories": trajectories,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechTrajectoryById(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if traj_id, err := strconv.Atoi(c.Param("traj_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if trajectory, err := s.database.GetMechTrajectoryById(mech_id, traj_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"trajectory": trajectory,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechTrajectoryDKT(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if traj_id, err := strconv.Atoi(c.Param("traj_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if dkt, err := s.database.GetMechTrajectoryDKT(mech_id, traj_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"dkt": dkt,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) getMechTrajectoryIKT(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		if mech_id, err := strconv.Atoi(c.Param("mech_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if traj_id, err := strconv.Atoi(c.Param("traj_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if ikt, err := s.database.GetMechTrajectoryIKT(mech_id, traj_id); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"ikt": ikt,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func (s *Server) saveMechTrajectory(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован в системе",
		})
	} else {
		var addRequest models.TrajectoryAddRequest
		mech_id, err := strconv.Atoi(c.Param("mech_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		mechConfig, _ := s.database.GetCurrentMechConfig(mech_id)
		err = c.ShouldBindJSON(&addRequest)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var id int
		id, err = s.database.AddMechTrajectory(mechConfig.Id, &addRequest)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": "Траектория была успешно добавлена",
		})
	}
}

func (s *Server) mechDKTsolver(c *gin.Context) {
	// Добавить авторизацию
	mech_id, err := strconv.Atoi(c.Param("mech_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	mechConfig, _ := s.database.GetCurrentMechConfig(mech_id)
	mechInfo, _ := s.database.GetMechanismInfo(mech_id)
	k := kinematics.New(mechInfo.Type, mechConfig)

	switch mechInfo.Type {
	case "msom":
		var points []models.MsomDirectTaskRequest
		if err := c.ShouldBindJSON(&points); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var solution []*models.MsomDirectTaskSolution
		for i, task := range points {
			point_ikt := k.MsomDirectSolver(&task)
			point_ikt.Id = i
			solution = append(solution, point_ikt)
		}
		c.JSON(http.StatusOK, gin.H{
			"solution": solution,
		})
		return
	case "5R":
		log.Println("jhkjh")
		// case "3R"
	}

}

func (s *Server) mechIKTsolver(c *gin.Context) {
	mech_id, err := strconv.Atoi(c.Param("mech_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	mechConfig, _ := s.database.GetCurrentMechConfig(mech_id)
	mechInfo, _ := s.database.GetMechanismInfo(mech_id)
	k := kinematics.New(mechInfo.Type, mechConfig)

	switch mechInfo.Type {
	case "msom":
		var points []models.MsomInverseTaskRequest
		if err := c.ShouldBindJSON(&points); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var solution []*models.MsomInverseTaskSolution
		for i, task := range points {
			point_ikt := k.MsomInverseSolver(&task)
			log.Println(task)
			point_ikt.Id = i
			solution = append(solution, point_ikt)
		}
		c.JSON(http.StatusOK, gin.H{
			"solution": solution,
		})
		return
	case "5R":
		log.Println("jhkjh")
		// case "3R"
	}
}

func (s *Server) mechIKTinterpolation(c *gin.Context) {
	mech_id, err := strconv.Atoi(c.Param("mech_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var interpolationTask []*models.InterpolationRequest
	if err := c.ShouldBindJSON(&interpolationTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(interpolationTask)
	mechConfig, _ := s.database.GetCurrentMechConfig(mech_id)
	mechInfo, _ := s.database.GetMechanismInfo(mech_id)
	k := kinematics.New(mechInfo.Type, mechConfig)

	var solutions []*models.InterpolationSolution
	for _, task := range interpolationTask {
		splines := k.GetCubicSplines(task.Time, task.Val)
		var solution models.InterpolationSolution
		solution.Id = task.Id
		solution.T = task.Time[len(task.Time)-1]
		solution.Splines = splines
		solutions = append(solutions, &solution)
	}

	c.JSON(http.StatusOK, gin.H{
		"solution": solutions,
	})
}
