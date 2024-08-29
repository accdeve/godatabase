package controllers

import (
	"database/models"
	"database/sql"
	"log"
)

var DB *sql.DB

func InsertDeviceData(device models.Device) error {
	// query := `INSERT INTO device(type, port, data, counter, deveui, radio) VALUES (?, ?, ?, ?, ?, ?)`

	// _, err := DB.Exec(query, device.Type, device.Port, device.Data, device.Counter, device.DevEui, device.Radio)

	// if err != nil{
	// 	log.Printf("error di insert database")
	// 	return err
	// }

	// log.Printf("insert data berhasil")
	log.Printf(device.Data)
	log.Printf(device.DevEui)
	log.Printf(device.Radio)	
	log.Printf(device.Type)	
	log.Printf("%v",device.Counter)	
	log.Printf("%v",device.Port)	

	return nil
}