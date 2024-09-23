package models

type Modulation struct {
	Bandwidth int    `json:"bandwidth"`
	Type      string `json:"type"`
	Spreading int    `json:"spreading"`
	CodeRate  string `json:"coderate"`
}