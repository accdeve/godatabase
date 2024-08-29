package models

import "encoding/json"

func UnmarshalEmpty(data []byte) (Device, error) {
	var r Device
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Device) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Device struct {
	Type    string `json:"type"`
	Port    int64  `json:"port"`
	Data    string `json:"data"`
	Counter int64  `json:"counter"`
	DevEui  string `json:"devEui"`
	Radio   string `json:"radio"`
}
