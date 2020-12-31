package service

import (
	"context"
	"github.com/pkg/errors"
	"log"

	"github.com/longjoy/micro-go-course/register/discovery"
)

type Service interface {
	HealthCheck() string
	DiscoveryService(ctx context.Context, serviceName string) ([]*discovery.InstanceInfo, error)
}

type RegisterServiceImpl struct {
	discoveryClient *discovery.DiscoveryClient
}

var ErrNotServiceInstances = errors.New("instances are not existed")

func NewRegisterServiceImpl(discoveryClient *discovery.DiscoveryClient) Service  {
	return &RegisterServiceImpl{
		discoveryClient:discoveryClient,
	}
}

func (service RegisterServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]*discovery.InstanceInfo, error) {
	instances, err := service.discoveryClient.DiscoverServices(ctx, serviceName)
	if err != nil{
		log.Printf("get service info err: %s", err)
	}
	if instances == nil || len(instances) == 0{
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

func (*RegisterServiceImpl) HealthCheck() string {
	return "OK"
}