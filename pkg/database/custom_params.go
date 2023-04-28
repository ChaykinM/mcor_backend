package database

import (
	"fmt"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetCustomParams() ([]*models.CustomParam, error) {
	var params []*models.CustomParam
	request := fmt.Sprintf("SELECT id, name, tag, description FROM public.custom_params;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return params, err
	} else {
		for rows.Next() {
			var id int
			var name, tag, description string
			rows.Scan(&id, &name, &tag, &description)
			var param models.CustomParam
			param.Id = id
			param.Name = name
			param.Tag = tag
			param.Description = description
			params = append(params, &param)
		}
	}
	return params, nil
}

func (d *Database) GetCustomParamById(param_id int) (*models.CustomParam, error) {
	var param models.CustomParam
	request := fmt.Sprintf("SELECT name, tag, description FROM public.custom_params WHERE id = %d;", param_id)
	rows := d.dbDriver.QueryRow(request)

	var name, tag, description string
	err := rows.Scan(&name, &tag, &description)
	param.Id = param_id
	param.Name = name
	param.Tag = tag
	param.Description = description

	return &param, err
}

func (d *Database) AddCustomParam(param *models.CustomParamAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.custom_params(id, name, tag, description) VALUES(DEFAULT, '%s', '%s', '%s') RETURNING id;", param.Name, param.Tag, param.Description)
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

func (d *Database) EditCustomParam(param *models.CustomParam) error {
	request := fmt.Sprintf("UPDATE public.custom_params SET name = '%s', tag = '%s', description = '%s' WHERE id = %d", param.Name, param.Tag, param.Description, param.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}

func (d *Database) DeleteCustomParam(param_id int) error {
	request := fmt.Sprintf("DELETE FROM public.custom_params WHERE id = %d", param_id)
	_, err := d.dbDriver.Exec(request)
	return err
}
