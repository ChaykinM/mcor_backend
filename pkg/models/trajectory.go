package models

type Trajectory struct {
	Id           int    `json:"id"`
	Mech_id      int    `json:"mech_id"`
	MechConfigId int    `json:"mech_config_id"`
	Time         string `json:"time"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DKTpoint     string `json:"dkt_point"`
	DKTpolinoms  string `json:"dkt_polinoms"`
	IKTpoint     string `json:"ikt_point"`
	IKTpolinoms  string `json:"ikt_polinoms"`
}

type TrajectoryAddRequest struct {
	Mech_id     int    `json:"mech_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DKTpoint    string `json:"dkt_point"`
	DKTpolinoms string `json:"dkt_polinoms"`
	IKTpoint    string `json:"ikt_point"`
	IKTpolinoms string `json:"ikt_polinoms"`
}

type DKT struct {
	Id           int    `json:"traj_id"`
	Mech_id      int    `json:"mech_id"`
	MechConfigId int    `json:"mech_config_id"`
	DKTpoint     string `json:"dkt_point"`
	DKTpolinoms  string `json:"dkt_polinoms"`
}

type IKT struct {
	Id           int    `json:"id"`
	Mech_id      int    `json:"mech_id"`
	MechConfigId int    `json:"mech_config_id"`
	IKTpoint     string `json:"ikt_point"`
	IKTpolinoms  string `json:"ikt_polinoms"`
}

type MsomDirectTaskRequest struct {
	Id   int     `json:"id"`
	Time float64 `json:"time"`
	Q_1  float64 `json:"q_1"`
	Q_2  float64 `json:"q_2"`
	Q_3  float64 `json:"q_3"`
	Q_4  float64 `json:"q_4"`
	Q_5  float64 `json:"q_5"`
	Q_6  float64 `json:"q_6"`
}

type FiveBarDirectTaskRequest struct {
}

type ThreeRotDirectTaskRequest struct {
}

type MsomDirectTaskSolution struct {
	Id    int     `json:"id"`
	Time  float64 `json:"time"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Alpha float64 `json:"alpha"`
	Betta float64 `json:"betta"`
}

type FiveBarDirectTaskSolution struct {
}

type ThreeRotDirectTaskSolution struct {
}

type MsomInverseTaskRequest struct {
	Id    int     `json:"id"`
	Time  float64 `json:"time"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Alpha float64 `json:"alpha"`
	Betta float64 `json:"betta"`
}

type FiveBarInverseTaskRequest struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ThreeRotInverseTaskRequest struct {
}

type MsomInverseTaskSolution struct {
	Id   int     `json:"id"`
	Time float64 `json:"time"`
	Q_1  float64 `json:"q_1"`
	Q_2  float64 `json:"q_2"`
	Q_3  float64 `json:"q_3"`
	Q_4  float64 `json:"q_4"`
	Q_5  float64 `json:"q_5"`
	Q_6  float64 `json:"q_6"`
}

type CubicSplineData struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	D float64 `json:"d"`
	X float64 `json:"x"`
}

type InterpolationRequest struct {
	Id   int       `json:"id"`
	Time []float64 `json:"time"`
	Val  []float64 `json:"val"`
}

type InterpolationSolution struct {
	Id      int                `json:"id"`
	T       float64            `json:"T"`
	Splines []*CubicSplineData `json:"splines"`
}
