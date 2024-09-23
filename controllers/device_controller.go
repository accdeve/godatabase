package controllers

import (
	"database/database"
	"database/models"
	"fmt"
	"log"
)

func InsertDeviceData(device models.Device) error {
	database := database.GetDB()

	if database == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	query := `INSERT INTO device(type, port, data, counter, deveui, radio) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := database.Exec(query, device.Type, device.Port, device.Data, device.Counter, device.DevEui, device.Radio)

	if err != nil {
		log.Printf("error di insert database")
		return err
	}

	return nil
}
