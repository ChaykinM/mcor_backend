package models

type Param struct {
	Id          int    `json:"param_id"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type ParamAddRequest struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

// type MotorParam struct {
// 	MechMotorId int      `json:"mech_motor_id"`
// 	Name        string   `json:"name"`
// 	Params      []*Param `json:"params"`
// }

type CustomParam struct {
	Id          int    `json:"param_id"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type CustomParamAddRequest struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
