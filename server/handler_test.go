package server

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)


func TestCalculateScore(t *testing.T) {

	testTime, err := time.Parse("2006-01-02", "2025-05-13")
	if err != nil {
		panic("Error parsing time")
		return
	}
	// Test cases
	testCases := []struct {
		company CompanyResponse
		expected   float64
	}{
		{CompanyResponse{"Company A", "EMPA", "X", "upgraded by", "Neutral", "Buy", "$100.00", "$120.00", "EMPA",testTime, 0}, 3.0},
		{CompanyResponse{"Company B", "EMPB", "Y", "target lowered by", "Buy", "Buy", "$50.00", "$45.00", "EMPB", testTime,0}, -1.0},
		{CompanyResponse{"Company C", "EMPC", "Z", "reiterated by", "Market Perform", "Buy", "$20.00", "$22.00", "EMPC", testTime, 0}, 1.0},
	}

	for _, tc := range testCases {
		CalculateRecommendationScore(&tc.company)
		assert.Equal(t, tc.expected, tc.company.Score, "Expected score for %s to be %f, but got %f", tc.company.Company, tc.expected, tc.company.Score)	
	}
}

func TestCalculateScoreFail(t *testing.T) {
	testTime, err := time.Parse("2006-01-02", "2025-05-13")
	if err != nil {
		panic("Error parsing time")
		return
	}
	// Test cases
	testCases := []struct {
		company CompanyResponse
		expected   float64
	}{
		{CompanyResponse{"Company A", "EMPA", "X", "upgraded by", "Neutral", "Buy", "$100.00", "$120.00", "EMPA",testTime, 0}, 1.5},
		{CompanyResponse{"Company B", "EMPB", "Y", "target lowered by", "Buy", "Buy", "$50.00", "$45.00", "EMPB", testTime,0}, 2.0},
		{CompanyResponse{"Company C", "EMPC", "Z", "reiterated by", "Market Perform", "Buy", "$20.00", "$22.00", "EMPC", testTime, 0}, 1.0},
	}

	for _, tc := range testCases {
		CalculateRecommendationScore(&tc.company)
		// Intentionally set an incorrect expected value
		assert.Equal(t, tc.expected, tc.company.Score, "Expected score for %s to be %f, but got %f", tc.company.Company, tc.expected, tc.company.Score)	
	}
}

func TestGetCompanies(t *testing.T) {
	expectedResponse := []CompanyResponse{
	{Id: "07be24f2-9931-482d-9277-27a13ce2b2fc", Action: "target lowered by", Brokerage: "Morgan Stanley", Company: "Lamar Advertising", Rating_from: "Equal Weight", Rating_to: "Equal Weight", Target_from: "$135.00", Target_to: "$125.00", Ticker: "LAMR", Time: toTime("2025-05-02T00:30:08.201845"), Score: -1.25},
	{Id: "19c6219d-6687-4e8a-b610-23a2d67c996b", Action: "target lowered by", Brokerage: "Keefe, Bruyette & Woods", Company: "BCB Bancorp", Rating_from: "Market Perform", Rating_to: "Market Perform", Target_from: "$12.50", Target_to: "$10.50", Ticker: "BCBP", Time: toTime("2025-04-25T00:30:07.385412"), Score: -1.25},
	{Id: "55cddc96-4264-4c7f-bf35-ddd1b06cfc6b", Action: "upgraded by", Brokerage: "HC Wainwright", Company: "CECO Environmental", Rating_from: "Neutral", Rating_to: "Buy", Target_from: "$33.00", Target_to: "$33.00", Ticker: "CECO", Time: toTime("2025-05-01T00:30:06.015697"), Score: 2.00},
	{Id: "789180b5-6e68-4de5-b0ae-9cdb5ba8783e", Action: "initiated by", Brokerage: "Leerink Partners", Company: "Akebia Therapeutics", Rating_from: "Outperform", Rating_to: "Outperform", Target_from: "$7.00", Target_to: "$7.00", Ticker: "AKBA", Time: toTime("2025-04-29T00:30:06.253902"), Score: 1.25},
	{Id: "80e9a0c6-8791-4b9a-85dd-875faa32dc7b", Action: "target lowered by", Brokerage: "Wells Fargo & Company", Company: "Blend Labs", Rating_from: "Overweight", Rating_to: "Overweight", Target_from: "$6.00", Target_to: "$5.00", Ticker: "BLND", Time: toTime("2025-04-23T00:30:08.235831"), Score: -1.25},
	{Id: "8d9c8688-7594-4aef-b1af-853c87a971a3", Action: "reiterated by", Brokerage: "Truist Financial", Company: "REX American Resources", Rating_from: "Buy", Rating_to: "Buy", Target_from: "$55.00", Target_to: "$50.00", Ticker: "REX", Time: toTime("2025-03-28T00:30:05.10188"), Score: 0.5},
	{Id: "95421ded-986a-4aeb-ad0f-8429bc797036", Action: "reiterated by", Brokerage: "HC Wainwright", Company: "Voyager Therapeutics", Rating_from: "Buy", Rating_to: "Buy", Target_from: "$30.00", Target_to: "$30.00", Ticker: "VYGR", Time: toTime("2025-04-09T00:30:09.332772"), Score: 0.75},
	{Id: "b1caa227-43f9-4539-9a96-57db4403fac5", Action: "reiterated by", Brokerage: "Benchmark", Company: "Hello Group", Rating_from: "Buy", Rating_to: "Buy", Target_from: "$13.00", Target_to: "$13.00", Ticker: "MOMO", Time: toTime("2025-03-14T00:30:05.974622"), Score: 0.75},
	{Id: "bb35bc71-deb6-43f1-913a-e1a8e6668bd2", Action: "target lowered by", Brokerage: "HC Wainwright", Company: "Rockwell Medical", Rating_from: "Buy", Rating_to: "Buy", Target_from: "$7.00", Target_to: "$3.00", Ticker: "RMTI", Time: toTime("2025-03-25T00:30:06.000668"), Score: -1.00},
	{Id: "fe18f879-d5cb-41cc-9a09-586842e6a5d2", Action: "upgraded by", Brokerage: "The Goldman Sachs Group", Company: "Banco Santander (Brasil)", Rating_from: "Sell", Rating_to: "Neutral", Target_from: "$4.20", Target_to: "$4.70", Ticker: "BSRR", Time: toTime("2025-02-13T00:30:05.813548"), Score: 3.00},
}
	companies := GetCompaniesToTest()

	assert.Equal(t, expectedResponse, companies, "Expected companies to be equal")
}

func toTime(date string) time.Time {
	layout := "2006-01-02T15:04:05.999999999"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Error parsing time")
	}
	return t
}