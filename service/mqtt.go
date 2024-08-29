package service

import (
	"database/controllers"
	"database/models"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messagePubHandler(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("Received message : %s from topic %s\n", m.Payload(), m.Topic())

	parsedData, err := models.UnmarshalEmpty(m.Payload())
	if err != nil {
		log.Printf("Error parsing: %v", err)
		return
	}

	if err := controllers.InsertDeviceData(parsedData); err != nil {
		log.Printf("Error saving data to the database: %v", err)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %s", err)
}

func ConnectMQTT() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.telkomiot.id:1883").SetUsername("18f81b89c2a31a3e").SetPassword("18f81b89c2aaeaa0")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker %s", token.Error())
	}

	subcribe(client)
}

func subcribe(client mqtt.Client) {
	token := client.Subscribe("v2.0/subs/APP65d8663a863e793932/DEV66460ff8d029595565", 1, nil)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %s", "v2.0/subs/APP65d8663a863e793932/DEV66460ff8d029595565", token.Error())
	} else {
		fmt.Printf("Subscribed to topic: %s\n", "v2.0/subs/APP65d8663a863e793932/DEV66460ff8d029595565")
	}
	select {}
}

func Disconnect(client mqtt.Client) {
	client.Disconnect(250)
	log.Printf("disconnected from mqtt broker")
}
