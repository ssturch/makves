package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"net/http"
)

//go:embed ueba.csv
var uebaDataRaw []byte
var uebaDataPrepared map[string]any

type Item struct {
	Id   string `json:"id"`
	Data any    `json:"data"`
}

func main() {
	var portFlag string
	flag.StringVar(&portFlag, "port", "8000", "port number")
	flag.Parse()
	uebaDataPrepared = prepareData(uebaDataRaw)
	http.HandleFunc("/get-items", getItemsHandler)
	fmt.Println("service listen and serve port", portFlag)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", portFlag), nil))
}

func prepareData(data []byte) map[string]any {
	rawData, _ := gocsv.CSVToMaps(bytes.NewReader(data))
	csvData := make(map[string]any, len(rawData))
	for _, d := range rawData {
		id := d["id"]
		delete(d, "id")
		delete(d, "#")
		csvData[id] = d
	}
	return csvData
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryValues := r.URL.Query()
		var respRaw []Item
		for k, value := range queryValues {
			if k == "id" {
				for _, v := range value {
					data, ok := uebaDataPrepared[v]
					if ok {
						respRaw = append(respRaw, Item{Id: v, Data: data})
					}
				}
			}
		}
		resp, _ := json.Marshal(respRaw)
		w.WriteHeader(200)
		w.Write(resp)
	} else {
		w.WriteHeader(405)
	}
}
