package fakestore

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/gocarina/gocsv"
)

type Store struct {
}

//go:embed ueba.csv
var uebaDataRaw []byte
var uebaDataPrepared map[string]any

func New() Store {
	uebaDataPrepared = prepareData(uebaDataRaw)
	return Store{}
}

func (s Store) Close() {
	uebaDataPrepared = nil
}

func (s Store) GetInfoByIds(ids []string) ([]Item, error) {
	var res []Item
	for i := 0; i < len(ids); i++ {
		data, ok := uebaDataPrepared[ids[i]]
		if ok {
			res = append(res, Item{
				Id:   ids[i],
				Data: data,
			})
		}
	}
	if len(res) == 0 {
		return nil, errors.New("GetInfoByIds error: nil result")
	}
	return res, nil
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
