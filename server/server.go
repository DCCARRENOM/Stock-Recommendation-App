package server

import (
	"net/http"
	"time"
)

type Company struct {
	Company string `json:"company"`
	Brokerage string `json:"brokerage"`
	Action string `json:"action"`
	Rating_from string `json:"rating_from"`
	Rating_to string `json:"rating_to"`
	Target_from string `json:"target_from"`
	Target_to string `json:"target_to"`
	Ticker string `json:"ticker"`
	Time time.Time `json:"time"`
}

type CompanyList struct {
	Items []Company `json:"items"`
	Next_page string `json:"next_page"`
}

type CompanyResponse struct{
	Id string `json:"id"`
	Company string `json:"company"`
	Brokerage string `json:"brokerage"`
	Action string `json:"action"`
	Rating_from string `json:"rating_from"`
	Rating_to string `json:"rating_to"`
	Target_from string `json:"target_from"`
	Target_to string `json:"target_to"`
	Ticker string `json:"ticker"`
	Time time.Time `json:"time"`
	Score float64 `json:"score"`
}

func New(addr string) *http.Server{
	initRoutes()
	return &http.Server{
		Addr:addr,
	}
}