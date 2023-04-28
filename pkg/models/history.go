package models

type HistorySeans struct {
	Id          int    `json:"id"`
	MechId      int    `json:"mech_id"`
	MechName    string `json:"mech_name"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      string `json:"params"`
}

type HistorySeansAddRequest struct {
	MechId      int    `json:"mech_id"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      string `json:"params"`
}

type HistoryCustomParams struct {
	Id          int    `json:"id"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      string `json:"params"`
}
type HistoryCustomParamsAddRequest struct {
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      string `json:"params"`
}
