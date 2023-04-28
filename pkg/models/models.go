package models

type DbConnectInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Mechanism struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	StendImgUrl  string       `json:"stend_img_url"`
	StructImgUrl string       `json:"struct_img_url"`
	MechConfig   MechConfig   `json:"curr_mech_config"`
	Motors       []*MechMotor `json:"motors"`
}

type MechanismInfo struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	StendImgUrl  string `json:"stend_img_url"`
	StructImgUrl string `json:"struct_img_url"`
}

type MechInfoEditRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MechConfig struct {
	Id           int                `json:"id"`
	Mech_id      int                `json:"mech_id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	ConfigParams []*MechConfigParam `json:"config_params"`
	Type         string             `json:"type"`
	Current      bool               `json:"current"`
}

type MechConfigParam struct {
	Id          int     `json:"id"`
	Tag         string  `json:"tag"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type MechConfigAddRequest struct {
	Mech_id      int                `json:"mech_id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	ConfigParams []*MechConfigParam `json:"config_params"`
	Type         string             `json:"type"`
}

type MechConfigEditRequest struct {
	Id           int                `json:"id"`
	Mech_id      int                `json:"mech_id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	ConfigParams []*MechConfigParam `json:"config_params"`
	Type         string             `json:"type"`
}

type MechMotor struct {
	Id        int                 `json:"mech_motor_id"`
	Name      string              `json:"name"`
	MotorData Motor               `json:"motor_data"`
	Encoders  []*MechMotorEncoder `json:"encoders"`
	Params    []*MechMotorParam   `json:"params"`
}

type MechMotorAddRequest struct {
	MechId  int    `json:"mech_id"`
	MotorId int    `json:"motor_id"`
	Name    string `json:"name"`
}

type MechMotorEditRequest struct {
	MechMotorId int    `json:"mech_motor_id"`
	MotorId     int    `json:"motor_id"`
	Name        string `json:"name"`
}

type MechMotorParams struct {
	MechMotorId   int               `json:"mech_motor_id"`
	MechMotorName string            `json:"name"`
	Params        []*MechMotorParam `json:"params"`
}

type MechMotorParam struct {
	MechMotorParamId int   `json:"mech_motor_param_id"`
	ParamData        Param `json:"param_data"`
}

type MechMotorParamAddRequest struct {
	ParamId int `json:"param_id"`
}

type MechMotorEncoders struct {
	MechMotorId int                 `json:"mech_motor_id"`
	Encoders    []*MechMotorEncoder `json:"encoders"`
}

type MechMotorEncoder struct {
	MechMotorEncId int     `json:"mech_motor_enc_id"`
	EncoderData    Encoder `json:"encoder_data"`
}

type MechMotorEncoderAddRequest struct {
	EncoderId int `json:"enc_id"`
}
