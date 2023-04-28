package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetAllMechanisms() ([]*models.Mechanism, error) {
	var mechanisms []*models.Mechanism
	request := fmt.Sprintf("SELECT public.mechanisms.id, public.mechanisms.name, public.mechanisms.description, public.mechanisms.stend_img_url, public.mechanisms.struct_img_url, public.mech_configs.id, public.mech_configs.name, public.mech_configs.description, public.mech_configs.config, public.mech_configs.type FROM public.mechanisms LEFT JOIN public.mech_configs ON public.mechanisms.id = public.mech_configs.mech_id WHERE public.mech_configs.current = true;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return mechanisms, err
	} else {
		for rows.Next() {
			var id, configId int
			var name, description, stendImgUrl, structImgUrl, configName, configDescription, configStr, configType string
			var config []*models.MechConfigParam

			rows.Scan(&id, &name, &description, &stendImgUrl, &structImgUrl, &configId, &configName, &configDescription, &configStr, &configType)
			json.Unmarshal([]byte(configStr), &config)
			var mechanism models.Mechanism
			mechanism.Id = id
			mechanism.Name = name
			mechanism.Description = description
			mechanism.StendImgUrl = stendImgUrl
			mechanism.StructImgUrl = structImgUrl
			mechanism.MechConfig.Id = configId
			mechanism.MechConfig.Mech_id = id
			mechanism.MechConfig.Name = configName
			mechanism.MechConfig.Description = configDescription
			mechanism.MechConfig.ConfigParams = config
			mechanism.MechConfig.Type = configType
			mechanism.MechConfig.Current = true
			if motors, err := d.GetMechMotors(id); err == nil {
				mechanism.Motors = motors
			}
			mechanisms = append(mechanisms, &mechanism)
		}
	}
	return mechanisms, nil
}

func (d *Database) GetMechanism(id int) (*models.Mechanism, error) {
	var mechanism models.Mechanism
	request := fmt.Sprintf("SELECT public.mechanisms.name, public.mechanisms.description, public.mechanisms.stend_img_url, public.mechanisms.struct_img_url, public.mech_configs.id, public.mech_configs.name, public.mech_configs.description, public.mech_configs.config, public.mech_configs.type FROM public.mechanisms LEFT JOIN public.mech_configs ON public.mechanisms.id = public.mech_configs.mech_id WHERE public.mech_configs.current = true AND public.mechanisms.id = %d;", id)
	rows := d.dbDriver.QueryRow(request)

	var configId int
	var name, description, stendImgUrl, structImgUrl, configName, configDescription, configStr, configType string
	var config []*models.MechConfigParam

	err := rows.Scan(&name, &description, &stendImgUrl, &structImgUrl, &configId, &configName, &configDescription, &configStr, &configType)
	json.Unmarshal([]byte(configStr), &config)

	mechanism.Id = id
	mechanism.Name = name
	mechanism.Description = description
	mechanism.StendImgUrl = stendImgUrl
	mechanism.StructImgUrl = structImgUrl
	mechanism.MechConfig.Id = configId
	mechanism.MechConfig.Mech_id = id
	mechanism.MechConfig.Name = configName
	mechanism.MechConfig.Description = configDescription
	mechanism.MechConfig.ConfigParams = config
	mechanism.MechConfig.Type = configType
	mechanism.MechConfig.Current = true
	if motors, err := d.GetMechMotors(id); err == nil {
		mechanism.Motors = motors
	}

	return &mechanism, err
}

func (d *Database) GetAllMechanismsInfo() ([]*models.MechanismInfo, error) {
	var mechanisms []*models.MechanismInfo
	request := fmt.Sprintf("SELECT public.mechanisms.id, public.mechanisms.name, public.mechanisms.description, public.mechanisms.type, public.mechanisms.stend_img_url, public.mechanisms.struct_img_url FROM public.mechanisms ORDER BY public.mechanisms.id;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return mechanisms, err
	} else {
		for rows.Next() {
			var id int
			var name, description, mechType, stendImgUrl, structImgUrl string

			rows.Scan(&id, &name, &description, &mechType, &stendImgUrl, &structImgUrl)

			var mechanism models.MechanismInfo
			mechanism.Id = id
			mechanism.Name = name
			mechanism.Description = description
			mechanism.Type = mechType
			mechanism.StendImgUrl = stendImgUrl
			mechanism.StructImgUrl = structImgUrl
			mechanisms = append(mechanisms, &mechanism)
		}
	}
	return mechanisms, nil
}

func (d *Database) GetMechanismInfo(id int) (*models.MechanismInfo, error) {
	var mechanism models.MechanismInfo
	request := fmt.Sprintf("SELECT public.mechanisms.name, public.mechanisms.description, public.mechanisms.type, public.mechanisms.stend_img_url, public.mechanisms.struct_img_url FROM public.mechanisms WHERE public.mechanisms.id = %d;", id)
	rows := d.dbDriver.QueryRow(request)

	var name, description, mechType, stendImgUrl, structImgUrl string

	err := rows.Scan(&name, &description, &mechType, &stendImgUrl, &structImgUrl)

	mechanism.Id = id
	mechanism.Name = name
	mechanism.Description = description
	mechanism.Type = mechType
	mechanism.StendImgUrl = stendImgUrl
	mechanism.StructImgUrl = structImgUrl

	return &mechanism, err
}

func (d *Database) EditMechanismInfo(mechInfo *models.MechInfoEditRequest) error {
	request := fmt.Sprintf("UPDATE public.mechanisms SET name = '%s', description = '%s' WHERE id = %d", mechInfo.Name, mechInfo.Description, mechInfo.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()

	return err
}

func (d *Database) GetMechAllConfigs(mech_id int) ([]*models.MechConfig, error) {
	var configs []*models.MechConfig
	request := fmt.Sprintf("SELECT id, name, description, config, type, current FROM public.mech_configs WHERE mech_id = %d ORDER BY id;", mech_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return configs, err
	} else {
		for rows.Next() {
			var id int
			var name, description, configStr, configType string
			var current bool
			var config []*models.MechConfigParam

			rows.Scan(&id, &name, &description, &configStr, &configType, &current)
			json.Unmarshal([]byte(configStr), &config)
			var mech_config models.MechConfig
			mech_config.Id = id
			mech_config.Mech_id = mech_id
			mech_config.Name = name
			mech_config.Description = description
			mech_config.ConfigParams = config
			mech_config.Type = configType
			mech_config.Current = current
			configs = append(configs, &mech_config)
		}
	}
	return configs, nil
}

func (d *Database) GetMechConfig(mech_id int, config_id int) (*models.MechConfig, error) {
	request := fmt.Sprintf("SELECT name, description, config, type, current FROM public.mech_configs WHERE mech_id = %d AND id = %d;", mech_id, config_id)
	rows := d.dbDriver.QueryRow(request)

	var mechConfig models.MechConfig
	var name, description, configStr, config_type string
	var current bool
	var config []*models.MechConfigParam

	err := rows.Scan(&name, &description, &configStr, &config_type, &current)
	json.Unmarshal([]byte(configStr), &config)

	mechConfig.Id = config_id
	mechConfig.Mech_id = mech_id
	mechConfig.Name = name
	mechConfig.Description = description
	mechConfig.ConfigParams = config
	mechConfig.Type = config_type
	mechConfig.Current = current
	return &mechConfig, err
}

func (d *Database) GetCurrentMechConfig(mech_id int) (*models.MechConfig, error) {
	request := fmt.Sprintf("SELECT id, name, description, config, type FROM public.mech_configs WHERE mech_id = %d AND current = true;", mech_id)
	rows := d.dbDriver.QueryRow(request)

	var mechConfig models.MechConfig

	var id int
	var name, description, configStr, config_type string
	var config []*models.MechConfigParam

	err := rows.Scan(&id, &name, &description, &configStr, &config_type)
	json.Unmarshal([]byte(configStr), &config)

	mechConfig.Id = id
	mechConfig.Mech_id = mech_id
	mechConfig.Name = name
	mechConfig.Description = description
	mechConfig.ConfigParams = config
	mechConfig.Type = config_type
	mechConfig.Current = true
	return &mechConfig, err
}

func (d *Database) GetStandardMechConfig(mech_id int) (*models.MechConfig, error) {
	request := fmt.Sprintf("SELECT id, name, description, config, current FROM public.mech_configs WHERE mech_id = %d AND type = 'standard';", mech_id)
	rows := d.dbDriver.QueryRow(request)

	var mechConfig models.MechConfig

	var id int
	var name, description, configStr string
	var current bool
	var config []*models.MechConfigParam

	err := rows.Scan(&id, &name, &description, &configStr, &current)
	json.Unmarshal([]byte(configStr), &config)

	mechConfig.Id = id
	mechConfig.Mech_id = mech_id
	mechConfig.Name = name
	mechConfig.Description = description
	mechConfig.ConfigParams = config
	mechConfig.Type = "standard"
	mechConfig.Current = current
	return &mechConfig, err
}

func (d *Database) SetActiveNechConfig(mech_id int, config_id int) error {
	tx, err := d.dbDriver.Begin()
	if err != nil {
		return err
	}
	{
		request := fmt.Sprintf("UPDATE public.mech_configs SET current = false WHERE mech_id = %d;", mech_id)
		_, err := tx.Exec(request)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		request := fmt.Sprintf("UPDATE public.mech_configs SET current = true WHERE mech_id = %d AND id = %d;", mech_id, config_id)
		_, err := tx.Exec(request)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (d *Database) AddMechConfig(mechConfig *models.MechConfigAddRequest) (int, error) {
	configStr, _ := json.Marshal(mechConfig.ConfigParams)
	request := fmt.Sprintf("INSERT INTO public.mech_configs(id, mech_id, name, description, config, type, current) VALUES(DEFAULT, %d, '%s', '%s', '%s', '%s', 'false') RETURNING id;", mechConfig.Mech_id, mechConfig.Name, mechConfig.Description, configStr, mechConfig.Type)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

func (d *Database) EditMechConfig(mechConfig *models.MechConfigEditRequest) error {
	configStr, _ := json.Marshal(mechConfig.ConfigParams)

	request := fmt.Sprintf("UPDATE public.mech_configs SET name = '%s', description = '%s', config = '%s', type = '%s' WHERE id = %d AND mech_id = %d;", mechConfig.Name, mechConfig.Description, configStr, mechConfig.Type, mechConfig.Id, mechConfig.Mech_id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) GetMechMotors(mech_id int) ([]*models.MechMotor, error) {
	var mechMotors []*models.MechMotor
	request := fmt.Sprintf("SELECT  public.mech_motors.id, public.mech_motors.name, public.motors.id, public.motors.name, public.motors.description FROM public.mech_motors LEFT JOIN public.motors ON public.mech_motors.motor_id = public.motors.id WHERE public.mech_motors.mech_id = %d ORDER BY public.mech_motors.id; ", mech_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return mechMotors, err
	} else {
		for rows.Next() {
			var mech_motor_id, motor_id int
			var name, standard_name, description string
			rows.Scan(&mech_motor_id, &name, &motor_id, &standard_name, &description)
			var mechMotor models.MechMotor
			mechMotor.Id = mech_motor_id
			mechMotor.Name = name
			var standardMotor models.Motor
			standardMotor.Id = motor_id
			standardMotor.Name = standard_name
			standardMotor.Description = description
			mechMotor.MotorData = standardMotor
			if encoders, err := d.GetMechMotorEncoders(mech_motor_id); err == nil {
				mechMotor.Encoders = encoders.Encoders
			}

			if params, err := d.GetMechMotorParams(mech_motor_id); err == nil {
				mechMotor.Params = params
			}
			mechMotors = append(mechMotors, &mechMotor)
		}
	}
	if mechMotors == nil {
		return mechMotors, errors.New("Механизм не имеет приводов в базе данных.")
	}
	return mechMotors, nil
}

func (d *Database) GetMechMotor(mech_id int, mech_motor_id int) (*models.MechMotor, error) {
	var mechMotor models.MechMotor
	request := fmt.Sprintf("SELECT  public.mech_motors.name, public.motors.id, public.motors.name, public.motors.description FROM public.mech_motors LEFT JOIN public.motors ON public.mech_motors.motor_id = public.motors.id WHERE public.mech_motors.mech_id = %d AND public.mech_motors.id = %d; ", mech_id, mech_motor_id)
	rows := d.dbDriver.QueryRow(request)

	var motor_id int
	var name, standard_name, description string
	err := rows.Scan(&name, &motor_id, &standard_name, &description)
	mechMotor.Id = mech_motor_id
	mechMotor.Name = name
	var standardMotor models.Motor
	standardMotor.Id = motor_id
	standardMotor.Name = standard_name
	standardMotor.Description = description
	mechMotor.MotorData = standardMotor
	if encoders, err := d.GetMechMotorEncoders(mech_motor_id); err == nil {
		mechMotor.Encoders = encoders.Encoders
	}
	if params, err := d.GetMechMotorParams(mech_motor_id); err == nil {
		mechMotor.Params = params
	}
	return &mechMotor, err
}

func (d *Database) EditMechMotor(mechMotor *models.MechMotorEditRequest) error {
	request := fmt.Sprintf("UPDATE public.mech_motors SET motor_id = %d, name = '%s' WHERE id = %d", mechMotor.MotorId, mechMotor.Name, mechMotor.MechMotorId)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) AddMechMotor(mechMotor *models.MechMotorAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.mech_motors(id, mech_id, motor_id, name) VALUES(DEFAULT, %d, %d, '%s') RETURNING id;", mechMotor.MechId, mechMotor.MotorId, mechMotor.Name)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

func (d *Database) DeleteMechMotor(mech_motor_id int) error {
	ctx := context.Background()
	tx, err := d.dbDriver.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	request_DelParams := fmt.Sprintf("DELETE FROM public.mech_motors_params WHERE public.mech_motors_params.mech_motor_id = %d;", mech_motor_id)
	_, err = tx.ExecContext(ctx, request_DelParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	request_DelEncoders := fmt.Sprintf("DELETE FROM public.mech_motors_encoders WHERE public.mech_motors_encoders.mech_motor_id = %d;", mech_motor_id)
	_, err = tx.ExecContext(ctx, request_DelEncoders)
	if err != nil {
		tx.Rollback()
		return err
	}

	request_DelMotor := fmt.Sprintf("DELETE FROM public.mech_motors WHERE public.mech_motors.id = %d", mech_motor_id)
	_, err = tx.ExecContext(ctx, request_DelMotor)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (d *Database) GetMechMotorsParamsAll(mech_id int) ([]*models.MechMotorParams, error) {
	var motorParams []*models.MechMotorParams
	request_MotorsId := fmt.Sprintf("SELECT public.mech_motors.name, public.mech_motors.id FROM public.mech_motors LEFT JOIN public.motors ON public.mech_motors.motor_id = public.motors.id WHERE public.mech_motors.mech_id = %d; ", mech_id)
	rows, err := d.dbDriver.Query(request_MotorsId)
	if err != nil {
		return motorParams, err
	} else {
		for rows.Next() {
			var mech_motor_id int
			var mech_motor_name string
			rows.Scan(&mech_motor_name, &mech_motor_id)
			var motorParam models.MechMotorParams
			motorParam.MechMotorId = mech_motor_id
			motorParam.MechMotorName = mech_motor_name
			if params, err := d.GetMechMotorParams(mech_motor_id); err == nil {
				motorParam.Params = params
			}
			motorParams = append(motorParams, &motorParam)
		}
	}
	return motorParams, nil
}

func (d *Database) GetMechMotorParams(mech_motor_id int) ([]*models.MechMotorParam, error) {
	var params []*models.MechMotorParam
	request := fmt.Sprintf("SELECT public.mech_motors_params.id, public.mech_motors_params.param_id, public.params.name, public.params.tag, public.params.description FROM public.params LEFT JOIN public.mech_motors_params ON public.params.id = public.mech_motors_params.param_id WHERE public.mech_motors_params.mech_motor_id = %d;", mech_motor_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return params, err
	} else {
		for rows.Next() {
			var mech_motor_param_id, param_id int
			var name, tag, description string
			rows.Scan(&mech_motor_param_id, &param_id, &name, &tag, &description)
			var param models.MechMotorParam
			param.MechMotorParamId = mech_motor_param_id
			var paramData models.Param
			paramData.Id = param_id
			paramData.Name = name
			paramData.Tag = tag
			paramData.Description = description
			param.ParamData = paramData
			params = append(params, &param)
		}
	}
	if params == nil {
		return params, errors.New("Привод не имеет загруженных параметров.")
	}
	return params, nil
}

func (d *Database) AddMechMotorParam(mech_motor_id, param_id int) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.mech_motors_params(id, mech_motor_id, param_id) VALUES(DEFAULT, %d, %d) RETURNING id;", mech_motor_id, param_id)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}
func (d *Database) DeleteMechMotorParam(mech_motor_param_id int) error {
	request := fmt.Sprintf("DELETE FROM public.mech_motors_params WHERE public.mech_motors_params.id = %d;", mech_motor_param_id)
	_, err := d.dbDriver.Exec(request)
	return err
}

func (d *Database) GetMechMotorEncodersAll(mech_id int) ([]*models.MechMotorEncoders, error) {
	var allMechMotorEncoders []*models.MechMotorEncoders
	request_MotorsId := fmt.Sprintf("SELECT public.mech_motors.id FROM public.mech_motors LEFT JOIN public.motors ON public.mech_motors.motor_id = public.motors.id WHERE public.mech_motors.mech_id = %d; ", mech_id)
	rows, err := d.dbDriver.Query(request_MotorsId)
	if err != nil {
		return allMechMotorEncoders, err
	} else {
		for rows.Next() {
			var mech_motor_id int
			rows.Scan(&mech_motor_id)
			if mechMotorEncoders, err := d.GetMechMotorEncoders(mech_motor_id); err == nil {
				allMechMotorEncoders = append(allMechMotorEncoders, mechMotorEncoders)
			}
		}
	}
	return allMechMotorEncoders, nil
}

func (d *Database) GetMechMotorEncoders(mech_motor_id int) (*models.MechMotorEncoders, error) {
	var mechMotorEncoders models.MechMotorEncoders
	request := fmt.Sprintf("SELECT public.mech_motors_encoders.id, public.mech_motors_encoders.enc_id, public.encoders.name, public.encoders.description FROM public.mech_motors_encoders LEFT JOIN public.encoders ON public.mech_motors_encoders.enc_id = public.encoders.id WHERE mech_motor_id = %d ORDER BY public.mech_motors_encoders.enc_id;", mech_motor_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return &mechMotorEncoders, err
	} else {
		mechMotorEncoders.MechMotorId = mech_motor_id

		for rows.Next() {
			var mech_motor_enc_id, enc_id int
			var name, description string
			rows.Scan(&mech_motor_enc_id, &enc_id, &name, &description)
			var mechMotorEncoder models.MechMotorEncoder
			mechMotorEncoder.MechMotorEncId = mech_motor_enc_id
			var encoder models.Encoder
			encoder.Id = enc_id
			encoder.Name = name
			encoder.Description = description
			mechMotorEncoder.EncoderData = encoder
			mechMotorEncoders.Encoders = append(mechMotorEncoders.Encoders, &mechMotorEncoder)
		}
	}
	return &mechMotorEncoders, nil
}

func (d *Database) AddMechMotorEncoder(mech_motor_id int, enc_id int) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.mech_motors_encoders(id, mech_motor_id, enc_id) VALUES(DEFAULT, %d, %d) RETURNING id;", mech_motor_id, enc_id)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

func (d *Database) DeleteMechMotorEncoder(mech_motor_enc_id int) error {
	request := fmt.Sprintf("DELETE FROM public.mech_motors_encoders WHERE public.mech_motors_encoders.id = %d;", mech_motor_enc_id)
	_, err := d.dbDriver.Exec(request)
	return err
}

func (d *Database) GetMechHistorySeans(mech_id int) ([]*models.HistorySeans, error) {
	var histories []*models.HistorySeans
	request := fmt.Sprintf("SELECT id, time, name, description, params FROM public.history_mech_testing WHERE mech_id = %d;", mech_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return histories, err
	} else {
		for rows.Next() {
			var id int
			var time, name, description, params string
			rows.Scan(&id, &time, &name, &description, &params)
			var history models.HistorySeans
			history.Id = id
			history.MechId = mech_id
			history.Time = time
			history.Name = name
			history.Description = description
			history.Params = params
			histories = append(histories, &history)
		}
	}
	return histories, nil
}

func (d *Database) GetMechTrajectories(mech_id int) ([]*models.Trajectory, error) {
	var trajectories []*models.Trajectory
	request := fmt.Sprintf("SELECT id, mech_config_id, time, name, description, dkt_point, dkt_polinoms, ikt_point, ikt_polinoms FROM public.trajectories WHERE mech_id = %d", mech_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return trajectories, err
	} else {
		for rows.Next() {
			var id, mech_config_id int
			var time, name, description, dkt_point, dkt_polinoms, ikt_point, ikt_polinoms string
			rows.Scan(&id, &mech_config_id, &time, &name, &description, &dkt_point, &dkt_polinoms, &ikt_point, &ikt_polinoms)
			var trajectory models.Trajectory
			trajectory.Id = id
			trajectory.Mech_id = mech_id
			trajectory.MechConfigId = mech_config_id
			trajectory.Time = time
			trajectory.Name = name
			trajectory.Description = description
			trajectory.DKTpoint = dkt_point
			trajectory.DKTpolinoms = dkt_polinoms
			trajectory.IKTpoint = ikt_point
			trajectory.IKTpolinoms = ikt_polinoms
			trajectories = append(trajectories, &trajectory)
		}
	}
	return trajectories, nil
}

func (d *Database) GetMechTrajectoryById(mech_id int, traj_id int) (*models.Trajectory, error) {
	var trajectory models.Trajectory

	request := fmt.Sprintf("SELECT mech_config_id, time, name, description, dkt_point, dkt_polinoms, ikt_point, ikt_polinoms FROM public.trajectories WHERE id = %d AND mech_id = %d;", traj_id, mech_id)
	rows := d.dbDriver.QueryRow(request)

	var mech_config_id int
	var time, name, description, dkt_point, dkt_polinoms, ikt_point, ikt_polinoms string
	err := rows.Scan(&mech_config_id, &time, &name, &description, &dkt_point, &dkt_polinoms, &ikt_point, &ikt_polinoms)

	trajectory.Id = traj_id
	trajectory.Mech_id = mech_id
	trajectory.MechConfigId = mech_config_id
	trajectory.Time = time
	trajectory.Name = name
	trajectory.Description = description
	trajectory.DKTpoint = dkt_point
	trajectory.DKTpolinoms = dkt_polinoms
	trajectory.IKTpoint = ikt_point
	trajectory.IKTpolinoms = ikt_polinoms

	return &trajectory, err
}

func (d *Database) GetMechTrajectoryDKT(mech_id int, traj_id int) (*models.DKT, error) {
	var dkt models.DKT

	request := fmt.Sprintf("SELECT mech_config_id, dkt_point, dkt_polinoms FROM public.trajectories WHERE id = %d AND mech_id = %d;", traj_id, mech_id)
	rows := d.dbDriver.QueryRow(request)

	var mech_config_id int
	var dkt_point, dkt_polinoms string
	err := rows.Scan(&mech_config_id, &dkt_point, &dkt_polinoms)

	dkt.Id = traj_id
	dkt.Mech_id = mech_id
	dkt.MechConfigId = mech_config_id
	dkt.DKTpoint = dkt_point
	dkt.DKTpolinoms = dkt_polinoms

	return &dkt, err
}

func (d *Database) GetMechTrajectoryIKT(mech_id int, traj_id int) (*models.IKT, error) {
	var ikt models.IKT

	request := fmt.Sprintf("SELECT mech_config_id, ikt_point, ikt_polinoms FROM public.trajectories WHERE id = %d AND mech_id = %d;", traj_id, mech_id)
	rows := d.dbDriver.QueryRow(request)

	var mech_config_id int
	var ikt_point, ikt_polinoms string
	err := rows.Scan(&mech_config_id, &ikt_point, &ikt_polinoms)

	ikt.Id = traj_id
	ikt.Mech_id = mech_id
	ikt.MechConfigId = mech_config_id
	ikt.IKTpoint = ikt_point
	ikt.IKTpolinoms = ikt_polinoms

	return &ikt, err
}

func (d *Database) AddMechTrajectory(mech_config_id int, traj *models.TrajectoryAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.trajectories(id, mech_id, mech_config_id, name, description, dkt_point, dkt_polinoms, ikt_point, ikt_polinoms) VALUES(DEFAULT, %d, %d, '%s', '%s', '%s', '%s', '%s', '%s') RETURNING id;", traj.Mech_id, mech_config_id, traj.Name, traj.Description, traj.DKTpoint, traj.DKTpolinoms, traj.IKTpoint, traj.IKTpolinoms)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}
