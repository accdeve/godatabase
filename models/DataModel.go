package models

import "encoding/json"

func UnmarshalDataModel(data []byte) (DataModel, error) {
	var r DataModel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DataModel) Marshal() ([] byte, error){
	return json.Marshal(r)
}

type DataModel struct {
	Modulation Modulation `json:"modulation"`
	Hardware   Hardware   `json:"hardware"`
	Freq       int        `json:"freq"`
	DataRate   int        `json:"datarate"`
	Time       float64    `json:"time"`
}