package database

import (
	"errors"
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetMotors() ([]*models.Motor, error) {
	var motors []*models.Motor
	request := fmt.Sprintf("SELECT id, name, description, imgurl FROM public.motors;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return motors, err
	} else {
		for rows.Next() {
			var id int
			var name, description, imgurl string
			rows.Scan(&id, &name, &description, &imgurl)
			var motor models.Motor
			motor.Id = id
			motor.Name = name
			motor.Description = description
			motor.ImgUrl = imgurl
			motors = append(motors, &motor)
		}
	}
	if motors == nil {
		return motors, errors.New("Нет загруженных в базу данных приводов.")
	}
	return motors, nil
}

func (d *Database) GetMotorById(motor_id int) (*models.Motor, error) {
	var motor models.Motor
	request := fmt.Sprintf("SELECT name, description, imgurl FROM public.motors WHERE id = %d;", motor_id)
	rows := d.dbDriver.QueryRow(request)

	var name, description, imgurl string
	err := rows.Scan(&name, &description, &imgurl)
	motor.Id = motor_id
	motor.Name = name
	motor.Description = description
	motor.ImgUrl = imgurl

	return &motor, err
}

func (d *Database) EditMotor(motor *models.Motor) error {
	request := fmt.Sprintf("UPDATE public.motors SET name = '%s', description = '%s' WHERE id = %d", motor.Name, motor.Description, motor.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) UploadImgMotor(motor_id int, imgUrl string) error {
	request := fmt.Sprintf("UPDATE public.motors SET imgurl = '%s' WHERE id = %d", imgUrl, motor_id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) AddMotor(motor *models.MotorAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.motors(name, description) VALUES('%s', '%s') RETURNING id;", motor.Name, motor.Description)
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
