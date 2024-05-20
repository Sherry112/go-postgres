package models

type Stock struct {
	STOCKID int64  `json:"stockId"`
	NAME    string `json:"name"`
	PRICE   int64  `json:"price"`
	COMPANY string `json:"company"`
}
