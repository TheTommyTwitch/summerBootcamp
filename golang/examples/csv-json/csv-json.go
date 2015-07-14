package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
	// High, Low, Close
}

func main() {
	//open our source file
	src, err := os.Open("table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	//create our json file
	dst, err := os.Create("table.json")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	//read our source file
	rows, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	//create slice and append info to our
	records := make([]Record, 0, len(rows))
	for _, row := range rows {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	err = json.NewEncoder(dst).Encode(records)
	if err != nil {
		panic(err)
	}

}
