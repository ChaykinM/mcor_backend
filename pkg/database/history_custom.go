package database

import (
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetAllHistoryCustomParams() ([]*models.HistoryCustomParams, error) {
	var histories []*models.HistoryCustomParams
	request := fmt.Sprintf("SELECT id, time, name, description, params FROM public.history_custom_params ORDER BY time;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return histories, err
	} else {
		for rows.Next() {
			var id int
			var time, name, description, params string
			rows.Scan(&id, &time, &name, &description, &params)
			var history models.HistoryCustomParams
			history.Id = id
			history.Time = time
			history.Name = name
			history.Description = description
			history.Params = params
			histories = append(histories, &history)
		}
	}
	return histories, nil
}

func (d *Database) GetHistoryCustomParamsById(id int) (*models.HistoryCustomParams, error) {
	var history models.HistoryCustomParams
	request := fmt.Sprintf("SELECT time, name, description, params FROM public.history_custom_params WHERE id = %d;", id)
	rows := d.dbDriver.QueryRow(request)

	var time, name, description, params string

	err := rows.Scan(&time, &name, &description, &params)
	history.Id = id
	history.Name = name
	history.Description = description
	history.Params = params

	return &history, err
}

func (d *Database) EditHistoryCustomParams(history *models.HistoryCustomParams) error {
	request := fmt.Sprintf("UPDATE public.history_custom_params SET time = '%s', name = '%s', description = '%s', params = '%s' WHERE id = %d", history.Time, history.Name, history.Description, history.Params, history.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) AddNewHistoryCustomParams(history *models.HistoryCustomParamsAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.history_custom_params(id, time, name, description, params) VALUES(DEFAULT, '%s', '%s', '%s', '%s') RETURNING id;", history.Time, history.Name, history.Description, history.Params)
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

func (d *Database) DeleteHistoryCustomParams(id int) error {
	request := fmt.Sprintf("DELETE FROM public.history_custom_params WHERE id = %d", id)
	_, err := d.dbDriver.Exec(request)
	return err
}
