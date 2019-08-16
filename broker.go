package main

import (
	"context"
	"errors"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

type broker struct {
	services []brokerapi.Service
	logger   lager.Logger
	env      BrokerConfig
}

func (b *broker) Services(context context.Context) ([]brokerapi.Service, error) {
	//fmt.Println(b.services)
	return b.services, nil
}

func (b *broker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (b *broker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (b *broker) GetInstance(context context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	return brokerapi.GetInstanceDetailsSpec{}, nil
}

func (b *broker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool ) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, errors.New("service does not support bind")
}

func (b *broker) GetBinding(context context.Context, instanceID string, bindingID string) (brokerapi.GetBindingSpec, error) {
	return brokerapi.GetBindingSpec{}, nil
}

func (b *broker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool)  (brokerapi.UnbindSpec, error) {
	return brokerapi.UnbindSpec{}, errors.New("service does not support bind")
}

func (b *broker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (b *broker) LastOperation(context context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

func (b *broker) LastBindingOperation(context context.Context,  instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}