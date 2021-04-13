package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// чтение аргументов
	var datafile *string

	datafile = flag.String("datafile", "./indications.json", `Path to datafile. Default: "./indications.json"`)

	flag.Parse()

	log.Printf("Datafile: %s\n", *datafile)

	// открытие файла
	f, err := os.Open(*datafile)
	if err != nil {
		log.Printf("ERROR os.Open, %s\n", err)
		return
	}

	// преобразование в срез
	data, err := parseFile(f)
	if err != nil {
		log.Printf("ERROR parseFile, %s\n", err)
		return
	}

	// сортировка данных
	err = sortByDate(data)
	if err != nil {
		log.Printf("ERROR sortByDate, %s\n", err)
		return
	}

	// форматированный вывод
	print(data)
}

// Indication структура показания
type Indication struct {
	Indicator string    `json:"indicator"`
	Value     int       `json:"value"`
	Date      time.Time `json:"date"`
}

// преобразование файла в срез показаний приборов
func parseFile(file *os.File) (data []Indication, err error) {
	var (
		b []byte
	)

	// чтение содержимого файла
	b, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// преобразование из json в срез структур
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return
}

// функция сортировки данных по дате. Заменяет ссылку входного среза на отсортированный срез
func sortByDate(data []Indication) (err error) {

	var (
		sortedData []Indication
		saveObject Indication
	)

	sortedData = data

	for i := 0; i < len(sortedData); i++ {
		for j := i; j < len(sortedData); j++ {
			if sortedData[i].Date.Unix() < sortedData[j].Date.Unix() {
				saveObject = sortedData[i]
				sortedData[i] = sortedData[j]
				sortedData[j] = saveObject
			}
		}
	}

	//data = sortedData

	return
}

// функция форматированного вывода среза показаний
func print(data []Indication) {
	var (
		res string
	)

	for _, i := range data {
		res += fmt.Sprintf("value: %v indicator: %s, data: %v \n", i.Value, i.Indicator, i.Date)
	}

	log.Printf(res)
}
