package models

import "encoding/json"

func UnmarshalHardware(data []byte) (Hardware, error) {
	var r Hardware
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Hardware) Marshal()([]byte, error){
	return json.Marshal(r)
}

type Hardware struct {
	Status  int     `json:"status"`
	Chain   int     `json:"chain"`
	Tmst    int     `json:"tmst"`
	SNR     float64 `json:"snr"`
	RSSI    int     `json:"rssi"`
	Channel int     `json:"channel"`
	GPS     GPS     `json:"gps"`
}