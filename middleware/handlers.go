package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id, omitempty"`
	MESSAGE string `json:"message, omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the db")
	return db
}
func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stock, err := getAllStocks()
	if err != nil {
		log.Fatalf("Unable to get all stocks %v ", err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	prams := mux.Vars(r)
	id, err := strconv.Atoi(prams["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stock %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		MESSAGE: "stock created successfullt",
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}
	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records effected %v", updatedRows)
	res := response{
		ID:      int64(id),
		MESSAGE: msg,
	}
	json.NewEncoder(w).Encode(res)

}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}
	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records effected %v", deletedRows)
	res := response{
		ID:      int64(id),
		MESSAGE: msg,
	}
	json.NewEncoder(w).Encode(res)

}

func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	statement := `INSERT INTO stocks(name, price, company) values($1, $2, $3) RETURNING stockid`
	var id int64
	err := db.QueryRow(statement, stock.NAME, stock.PRICE, stock.COMPANY).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id

}
func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stock models.Stock
	statement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(statement, id)
	err := row.Scan(&stock.STOCKID, &stock.NAME, &stock.PRICE, &stock.COMPANY)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		fmt.Printf("Unable to scan the row %v", err)
	}
	return stock, err
}
func getAllStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stocks []models.Stock
	statement := `SELECT * FROM stocks`
	rows, err := db.Query(statement)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.STOCKID, &stock.NAME, &stock.PRICE, &stock.COMPANY)
		if err != nil {
			log.Fatalf("Unable to execute the query %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}
func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()

	statement := `UPDATE stocks SET name=$2,price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(statement, id, stock.NAME, stock.PRICE, stock.COMPANY)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the effected rows %v", err)
	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}
func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()

	statement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(statement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the effected rows %v", err)
	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}
