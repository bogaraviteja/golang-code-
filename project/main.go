package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"project/database"
	"project/models"
	"strconv"

	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	Db = database.Db()
	csvFile, _ := os.Open("info.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []*models.Employee
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		id, _ := strconv.ParseInt(line[0], 10, 64)
		age, _ := strconv.ParseInt(line[3], 10, 64)
		Db.Create(&models.Employee{
			Id:       id,
			Name:     line[1],
			Email:    line[2],
			Age:      age,
			Location: line[4],
		})
	}
	dataJson, _ := json.Marshal(data)
	fmt.Println(dataJson)
}
