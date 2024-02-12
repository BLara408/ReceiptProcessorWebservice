package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptIDResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

var receipts = make(map[string]Receipt)

func main() {
	http.HandleFunc("receipts/process", processReceipts)
	http.HandleFunc("receipts/{id}/points", getPoints)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func processReceipts(writer http.ResponseWriter, request *http.Request) {
	var receipt Receipt
	error := json.NewDecoder(request.Body).Decode(&receipts)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusBadRequest)
	}

	id := generateReceiptUUID()

	receipts[id] = receipt

	response := ReceiptIDResponse{ID: id}
	json.NewEncoder(writer).Encode(response)

}

func generateReceiptUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func getPoints(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	id := parts[len(parts)-2]

	receipt, found := receipts[id]
	if !found {
		http.Error(writer, "Receipt not found", http.StatusNotFound)
		return
	}

	points := calculatePointsForReceipt(receipt)

	response := PointsResponse{Points: points}

	json.NewEncoder(writer).Encode(response)

}

func calculatePointsForReceipt(receipt Receipt) int {

	return 0
}
