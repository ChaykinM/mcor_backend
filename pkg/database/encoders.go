package database

import (
	"errors"
	"fmt"
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetEncoders() ([]*models.Encoder, error) {
	var encoders []*models.Encoder
	request := fmt.Sprintf("SELECT id, name, description, imgurl FROM public.encoders;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return encoders, err
	} else {
		for rows.Next() {
			var id int
			var name, description, imgurl string
			rows.Scan(&id, &name, &description, &imgurl)
			var encoder models.Encoder
			encoder.Id = id
			encoder.Name = name
			encoder.Description = description
			encoder.ImgUrl = imgurl
			encoders = append(encoders, &encoder)
		}
	}
	if encoders == nil {
		return encoders, errors.New("Нет загруженных в базу данных энкодеров.")
	}
	return encoders, nil
}

func (d *Database) GetEncoder(enc_id int) (*models.Encoder, error) {
	var encoder models.Encoder
	request := fmt.Sprintf("SELECT name, description, imgurl FROM public.encoders WHERE id = %d;", enc_id)
	rows := d.dbDriver.QueryRow(request)

	var name, description, imgurl string
	err := rows.Scan(&name, &description, &imgurl)
	encoder.Id = enc_id
	encoder.Name = name
	encoder.Description = description
	encoder.ImgUrl = imgurl

	return &encoder, err
}

func (d *Database) EditEncoder(encoder *models.EncoderEditRequest) error {
	request := fmt.Sprintf("UPDATE public.encoders SET name = '%s', description = '%s' WHERE id = %d", encoder.Name, encoder.Description, encoder.Id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) UploadImgEncoder(enc_id int, imgUrl string) error {
	request := fmt.Sprintf("UPDATE public.encoders SET imgurl = '%s' WHERE id = %d", imgUrl, enc_id)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(err)
	return err
}

func (d *Database) AddEncoder(encoder *models.EncoderAddRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.encoders(name, description) VALUES('%s', '%s') RETURNING id;", encoder.Name, encoder.Description)
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
