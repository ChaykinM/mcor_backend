package database

import (
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetAllHistorySeans() ([]*models.HistorySeans, error) {
	var histories []*models.HistorySeans
	request := fmt.Sprintf("SELECT public.history_mech_testing.id, public.history_mech_testing.mech_id, public.mechanisms.name, public.history_mech_testing.time, public.history_mech_testing.name, public.history_mech_testing.description, public.history_mech_testing.params FROM public.history_mech_testing LEFT JOIN public.mechanisms ON public.mechanisms.id = public.history_mech_testing.mech_id ORDER BY time;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return histories, err
	} else {
		for rows.Next() {
			var id, mech_id int
			var mech_name, time, name, description, params string
			rows.Scan(&id, &mech_id, &mech_name, &time, &name, &description, &params)
			var history models.HistorySeans
			history.Id = id
			history.MechId = mech_id
			history.MechName = mech_name
			history.Time = time
			history.Name = name
			history.Description = description
			history.Params = params
			histories = append(histories, &history)
		}
	}
	return histories, nil
}

func (d *Database) GetHistorySeansById(id int) (*models.HistorySeans, error) {
	var history models.HistorySeans
	request := fmt.Sprintf("SELECT mech_id, time, name, description, params FROM public.history_mech_testing WHERE id = %d;", id)
	rows := d.dbDriver.QueryRow(request)

	var mech_id int
	var time, name, description, params string

	err := rows.Scan(&mech_id, &time, &name, &description, &params)
	history.Id = id
	history.MechId = mech_id
	history.Name = name
	history.Description = description
	history.Params = params

	return &history, err
}

func (d *Database) EditHistorySeans(history *models.HistorySeans) error {
	request := fmt.Sprintf("UPDATE public.history_mech_testing SET time = '%s', name = '%s', description = '%s', params = '%s' WHERE id = %d AND mech_id = %d", history.Time, history.Name, history.Description, history.Params, history.Id, history.MechId)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) AddNewHistorySeans(history *models.HistorySeansAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.history_mech_testing(id,mech_id, time, name, description, params) VALUES(DEFAULT, '%d', '%s', '%s', '%s', '%s') RETURNING id;", history.MechId, history.Time, history.Name, history.Description, history.Params)
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

func (d *Database) DeleteHistorySeans(id int) error {
	request := fmt.Sprintf("DELETE FROM public.history_mech_testing WHERE id = %d", id)
	_, err := d.dbDriver.Exec(request)
	return err
}
