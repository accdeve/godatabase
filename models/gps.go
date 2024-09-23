package models

import "encoding/json"

func UnmarshalGps(data []byte) (GPS, error) {
	var r GPS
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GPS) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type GPS struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Alt int     `json:"alt"`
}