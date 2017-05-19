package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

func main() {
	config, err := ConfigLoad()
	if err != nil {
		panic(err)
	}

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: config.BrokerUsername,
		Password: config.BrokerPassword,
	}

	services, err := CatalogLoad("./catalog.json")
	if err != nil {
		panic(err)
	}

	logger := lager.NewLogger("static-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	erviceBroker := &broker{services: services, logger: logger, env: config}
	brokerHandler := brokerapi.New(serviceBroker, logger, brokerCredentials)
	http.Handle("/", brokerHandler)
	http.ListenAndServe(":"+config.Port, nil)
}
