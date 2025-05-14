package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	//"server/server"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	//"github.com/google/uuid"
)
var DB *pgx.Conn

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo no permitido")
		return
	}
}

func initCompanies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	req,err := http.NewRequest("GET", "",nil)
	if (err != nil){}
	req.Header.Add("Authorization", "")

	client:= &http.Client{}
	response,err := client.Do(req)

	defer response.Body.Close()
	body,err := io.ReadAll(response.Body)

	var records CompanyList;

	json.Unmarshal(body,&records)

	
	/*for _, record := range records.Items {
		fmt.Fprintf(w, "%v\n", record.Company)
	}*/
	err = crdbpgx.ExecuteTx(context.Background(), DB, pgx.TxOptions{}, func(tx pgx.Tx) error {
		fmt.Fprintf(w, "%v\n", initTable(context.Background(), tx))
		return nil
    })

	err = crdbpgx.ExecuteTx(context.Background(), DB, pgx.TxOptions{}, func(tx pgx.Tx) error {
		for _, record := range records.Items {
			fmt.Fprintf(w, "%v\n", insertRows(context.Background(), tx, record))	
		}
		return nil
	})
	
}

func initTable(ctx context.Context, tx pgx.Tx) error {
    // Dropping existing table if it exists
    if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS companies"); err != nil {
        return err
    }

    // Create the accounts table
    if _, err := tx.Exec(ctx,"CREATE TABLE companies (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),action VARCHAR(100),brokerage VARCHAR(100),company VARCHAR(100),rating_from VARCHAR(100),rating_to VARCHAR(100),target_from VARCHAR(100),target_to VARCHAR(100),ticker VARCHAR(100),time TIMESTAMP)"); err != nil {
        return err
    }
    return nil
}

func insertRows(ctx context.Context, tx pgx.Tx, record Company) error {
	// Insert the record into the table
	if _, err := tx.Exec(ctx, "INSERT INTO companies (action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		record.Action,
		record.Brokerage,
		record.Company,
		record.Rating_from,
		record.Rating_to,
		record.Target_from,
		record.Target_to,
		record.Ticker,
		record.Time); err != nil {
		return err
	}
	return nil
}

func getCompanies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")

	rows,err := DB.Query(context.Background(), "SELECT * FROM companies;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer rows.Close()
	companies := []CompanyResponse{}
	for rows.Next() {
		var company CompanyResponse
		err := rows.Scan(&company.Id, &company.Action, &company.Brokerage, &company.Company, &company.Rating_from, &company.Rating_to, &company.Target_from, &company.Target_to, &company.Ticker, &company.Time)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err)
			return
		}
		CalculateRecommendationScore(&company)
		companies = append(companies, company)
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

func CalculateRecommendationScore(company *CompanyResponse){
	switch strings.ToLower(company.Action) {
	case "upgraded by":
		company.Score += 2
	case "initiated by":
		company.Score += 1
	case "reiterated by":
		lowerRatingTo := strings.ToLower(company.Rating_to)
		if strings.Contains(lowerRatingTo, "buy") || strings.Contains(lowerRatingTo, "outperform") {
			company.Score += 0.5
		} else if strings.Contains(lowerRatingTo, "sell") {
			company.Score -= 0.5
		}
	case "target lowered by":
		company.Score -= 1
	}

	// 2. Consistencia de la Calificación
	if strings.ToLower(company.Rating_from) == strings.ToLower(company.Rating_to) {
		lowerRatingTo := strings.ToLower(company.Rating_to)
		if strings.Contains(lowerRatingTo, "buy") || strings.Contains(lowerRatingTo, "outperform") {
			company.Score += 0.25
		} else if strings.Contains(lowerRatingTo, "sell") {
			company.Score -= 0.25
		}
	}

	// 3. Potencial de Crecimiento del Precio Objetivo
	targetFrom, errFrom := strconv.ParseFloat(strings.TrimPrefix(company.Target_from, "$"), 64)
	targetTo, errTo := strconv.ParseFloat(strings.TrimPrefix(company.Target_to, "$"), 64)

	if errFrom == nil && errTo == nil && targetFrom != 0 {
		growthPercentage := ((targetTo - targetFrom) / targetFrom) * 100
		if growthPercentage > 10 {
			company.Score += 1
		} else if growthPercentage > 5 {
			company.Score += 0.5
		} else if growthPercentage > 0 {
			company.Score += 0.25
		} else if growthPercentage < 0 {
			company.Score -= 0.25
		}
	}
}

func GetCompaniesToTest() []CompanyResponse{
	rows,err := DB.Query(context.Background(), "SELECT * FROM companies;")
	if err != nil {
		return nil
	}
	defer rows.Close()
	companies := []CompanyResponse{}
	for rows.Next() {
		var company CompanyResponse
		err := rows.Scan(&company.Id, &company.Action, &company.Brokerage, &company.Company, &company.Rating_from, &company.Rating_to, &company.Target_from, &company.Target_to, &company.Ticker, &company.Time)
		if err != nil {
			return nil
		}
		CalculateRecommendationScore(&company)
		companies = append(companies, company)
	}
	return companies
}
