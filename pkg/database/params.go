package database

import (
	"context"
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetParams() ([]*models.Param, error) {
	var params []*models.Param
	request := fmt.Sprintf("SELECT id, name, tag, description FROM public.params;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return params, err
	} else {
		for rows.Next() {
			var id int
			var name, tag, description string
			rows.Scan(&id, &name, &tag, &description)
			var param models.Param
			param.Id = id
			param.Name = name
			param.Tag = tag
			param.Description = description
			params = append(params, &param)
		}
	}
	return params, nil
}

func (d *Database) GetParamById(param_id int) (*models.Param, error) {
	var param models.Param
	request := fmt.Sprintf("SELECT name, tag, description FROM public.params WHERE id = %d;", param_id)
	rows := d.dbDriver.QueryRow(request)

	var name, tag, description string
	err := rows.Scan(&name, &tag, &description)
	param.Id = param_id
	param.Name = name
	param.Tag = tag
	param.Description = description

	return &param, err
}

func (d *Database) AddParam(param *models.ParamAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.params(id, name, tag, description) VALUES(DEFAULT, '%s', '%s', '%s') RETURNING id;", param.Name, param.Tag, param.Description)
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

func (d *Database) EditParam(param *models.Param) error {
	request := fmt.Sprintf("UPDATE public.params SET name = '%s', tag = '%s', description = '%s' WHERE id = %d", param.Name, param.Tag, param.Description, param.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}

func (d *Database) DeleteParam(param_id int) error {
	ctx := context.Background()
	tx, err := d.dbDriver.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	delete_MechMotorParams := fmt.Sprintf("DELETE FROM public.mech_motors_params WHERE param_id = %d;", param_id)
	_, err = tx.ExecContext(ctx, delete_MechMotorParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	delete_Params := fmt.Sprintf("DELETE FROM public.params WHERE id = %d;", param_id)
	_, err = tx.ExecContext(ctx, delete_Params)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
