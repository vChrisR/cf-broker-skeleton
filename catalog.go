package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pivotal-cf/brokerapi"
)

func CatalogLoad() ([]brokerapi.Service, error) {
	var services []brokerapi.Service

	inBuf, err := ioutil.ReadFile("./catalog.json")
	if err != nil {
		return []brokerapi.Service{}, err
	}

	err = json.Unmarshal(inBuf, &services)
	if err != nil {
		return []brokerapi.Service{}, err
	}
	return services, nil
}
