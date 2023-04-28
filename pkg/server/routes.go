package server

func (s *Server) routerSettings() {
	s.authRoutes()
	s.mechanismsRoutes()
	s.motorsRoutes()
	s.encodersRoutes()
	s.paramsRoutes()
	s.customParamsRoutes()
	s.historySeansRoutes()
	s.historyCustomParams()
	s.userRoutes()

	s.assetsStaticRoutes()
}

func (s *Server) authRoutes() {
	authRouters := s.router.Group("/auth")
	{
		authRouters.POST("/login", s.loginHandler)
		authRouters.POST("/register", s.registerHandler)
	}
}

func (s *Server) mechanismsRoutes() {
	mechanismsRoutes := s.router.Group("/mechanisms/", s.userIdentification)
	{
		mechanismsRoutes.GET("/", s.getAllMechanisms)
		mechanismsRoutes.GET(":mech_id", s.getMechanism)

		mechanismsRoutes.GET("/info", s.getAllMechanismsInfo)
		mechanismsRoutes.GET(":mech_id/info", s.getMechanismInfo)
		mechanismsRoutes.POST("/:mech_id/info/edit", s.editMechanismInfo)

		mechanismsRoutes.GET(":mech_id/configs", s.getMechAllConfigs)                         // получить все доступные конфигурации механизма
		mechanismsRoutes.GET(":mech_id/configs/:config_id", s.getMechConfig)                  // получить конфигурацию механизма по её id
		mechanismsRoutes.GET(":mech_id/configs/current", s.getCurrentMechConfig)              // получить текущую конфигурацию механизма
		mechanismsRoutes.GET(":mech_id/configs/standard", s.getStandardMechConfig)            // получить стандартную конфигурацию механизма
		mechanismsRoutes.POST(":mech_id/configs/:config_id/setActive", s.setActiveMechConfig) // сделать конфигурацию текущей
		mechanismsRoutes.POST(":mech_id/configs/add", s.addMechConfig)                        // добавить новую конфигурацию механизма
		mechanismsRoutes.POST(":mech_id/configs/:config_id/edit", s.editMechConfig)           // отредактировать конфигурацию механизма

		mechanismsRoutes.GET("/:mech_id/motors", s.getMechMotors)
		mechanismsRoutes.POST("/:mech_id/motors/add", s.addMechMotor)

		mechanismsRoutes.GET("/:mech_id/motors/:mech_motor_id", s.getMechMotor)
		mechanismsRoutes.POST("/:mech_id/motors/:mech_motor_id/edit", s.editMechMotor)
		mechanismsRoutes.DELETE("/:mech_id/motors/:mech_motor_id/del", s.deleteMechMotor)

		mechanismsRoutes.GET("/:mech_id/motors/params", s.getMechMotorsParamsAll)
		mechanismsRoutes.GET("/:mech_id/motors/:mech_motor_id/params", s.getMechMotorParams)
		mechanismsRoutes.POST("/:mech_id/motors/:mech_motor_id/params/add", s.addMechMotorParam)
		mechanismsRoutes.DELETE("/:mech_id/motors/:mech_motor_id/params/:mech_motor_param_id/del", s.deleteMechMotorParam)

		mechanismsRoutes.GET("/:mech_id/motors/encoders", s.getMechMotorEncodersAll)
		mechanismsRoutes.GET("/:mech_id/motors/:mech_motor_id/encoders", s.getMechMotorEncoders)
		mechanismsRoutes.POST("/:mech_id/motors/:mech_motor_id/encoders/add", s.addMechMotorEncoders)
		mechanismsRoutes.DELETE("/:mech_id/motors/:mech_motor_id/encoders/:mech_motor_enc_id/del", s.deleteMechMotorEncoder)

		mechanismsRoutes.GET("/:mech_id/histories", s.getMechHistorySeans)
		// параметры для истории сеансов

		// Решения траекторной задачи
		mechanismsRoutes.GET("/:mech_id/trajectories", s.getMechTrajectories)
		mechanismsRoutes.GET("/:mech_id/trajectories/:traj_id", s.getMechTrajectoryById)
		mechanismsRoutes.GET("/:mech_id/trajectories/:traj_id/dkt", s.getMechTrajectoryDKT)
		mechanismsRoutes.GET("/:mech_id/trajectories/:traj_id/ikt", s.getMechTrajectoryIKT)
		mechanismsRoutes.POST("/:mech_id/trajectories/save", s.saveMechTrajectory)

		// Задание траекторий
		// mechanismsRoutes.POST("/:mech_id/kinematics/set_trajectory")
		mechanismsRoutes.POST("/:mech_id/kinematics/dkt_solver", s.mechDKTsolver)
		mechanismsRoutes.POST("/:mech_id/kinematics/ikt_solver", s.mechIKTsolver)
		mechanismsRoutes.POST("/:mech_id/kinematics/interpolation", s.mechIKTinterpolation)

		// ?????????????
		// get /:mech_id/kinematics/dkt -> получить все пзк
		// post /:mech_id/kinematics/dkt/soilver -> решить ПЗК

	}
}

func (s *Server) motorsRoutes() {
	motorsRoutes := s.router.Group("/motors", s.userIdentification)
	{
		motorsRoutes.GET("/", s.getMotor)
		motorsRoutes.GET("/:motor_id", s.getMotorById)
		motorsRoutes.POST("/:motor_id/edit", s.editMotor)
		motorsRoutes.POST("/:motor_id/uploadImg", s.uploadImgMotor)
		motorsRoutes.POST("/add", s.addMotor)
	}
}

func (s *Server) encodersRoutes() {
	encodersRoutes := s.router.Group("/encoders", s.userIdentification)
	{
		encodersRoutes.GET("/", s.getEncoders)
		encodersRoutes.GET("/:enc_id", s.getEncoder)
		encodersRoutes.POST("/:enc_id/edit", s.editEncoder)
		encodersRoutes.POST("/:enc_id/uploadImg", s.uploadImgEncoder)
		encodersRoutes.POST("/add", s.addEncoder)
	}
}

func (s *Server) paramsRoutes() {
	paramsRoutes := s.router.Group("/params", s.userIdentification)
	{
		paramsRoutes.GET("/", s.getParams)
		paramsRoutes.GET("/:param_id", s.getParamById)
		paramsRoutes.POST("/add", s.addParam)
		paramsRoutes.POST("/:param_id/edit", s.editParam)
		paramsRoutes.DELETE("/:param_id/del", s.deleteParam)

	}
}

func (s *Server) customParamsRoutes() {
	customParamsRoutes := s.router.Group("/custom_params", s.userIdentification)
	{
		customParamsRoutes.GET("/", s.getCustomParams)
		customParamsRoutes.GET("/:custom_param_id", s.getCustomParamById)
		customParamsRoutes.POST("/add", s.addCustomParam)
		customParamsRoutes.POST("/:custom_param_id/edit", s.editCustomParam)
		customParamsRoutes.DELETE("/:custom_param_id/del", s.deleteCustomParam)
	}
}

func (s *Server) historySeansRoutes() {
	historySeansRoutes := s.router.Group("histories", s.userIdentification)
	{
		historySeansRoutes.GET("/", s.getAllHistorySeans)
		historySeansRoutes.GET("/:id", s.getHistorySeansById)
		historySeansRoutes.POST("/:id/edit", s.editHistorySeans)
		historySeansRoutes.POST("/add", s.addNewHistorySeans)
		historySeansRoutes.DELETE("/:id/del", s.deleteHistorySeans)
	}
}

func (s *Server) historyCustomParams() {
	historyCustomParamsRoutes := s.router.Group("history_custom_params", s.userIdentification)
	{
		historyCustomParamsRoutes.GET("/", s.getAllHistoryCustomParams)
		historyCustomParamsRoutes.GET("/:id", s.getHistoryCustomParamsById)
		historyCustomParamsRoutes.POST("/:id/edit", s.editHistoryCustomParams)
		historyCustomParamsRoutes.POST("/add", s.addNewHistoryCustomParams)
		historyCustomParamsRoutes.DELETE("/:id/del", s.deleteHistoryCustomParams)
	}
}

func (s *Server) userRoutes() {
	userRoutes := s.router.Group("users", s.userIdentification)
	{
		userRoutes.GET("/", s.getUsers)
		userRoutes.GET("/:user_id", s.getUserByID)
		userRoutes.POST("/:user_id/edit", s.editUser)
		userRoutes.DELETE("/:user_id/del", s.deleteUser)
		userRoutes.POST("/:user_id/updPass", s.updateUserPassword)
	}
}

func (s *Server) assetsStaticRoutes() {
	s.router.Static("/assets/mechanisms", "./assets/mechanisms")
	s.router.Static("/assets/encoders", "./assets/encoders")
	s.router.Static("/assets/motors", "./assets/motors")
}
