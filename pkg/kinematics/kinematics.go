package kinematics

import (
	"errors"
	"log"

	"main.go/pkg/models"
	// "../models"
)

type Kinematics struct {
	mech_type string
	config    *models.MechConfig
	// functions
	// directSolver(*task);
	// inverseSolver();
	// workspace();
	// singularity(); and etc
}

func New(mech_type string, config *models.MechConfig) *Kinematics {
	log.Println(config)
	return &Kinematics{
		mech_type: mech_type,
		config:    config,
	}
}

func (k *Kinematics) findConfigParam(tag string) (float64, error) {
	var configParam *models.MechConfigParam
	for _, configParam = range k.config.ConfigParams {
		if configParam.Tag == tag {
			return configParam.Value, nil
		}
	}
	return configParam.Value, errors.New("not found param by tag")
}
