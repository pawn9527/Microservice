package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/longjoy/micro-go-course/register/discovery"
	"github.com/longjoy/micro-go-course/register/endpoint"
	"github.com/longjoy/micro-go-course/register/service"
	"github.com/longjoy/micro-go-course/register/transport"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	consulAddr := flag.String("consul.addr", "localhost", "consul address")	
	consulPort  := flag.Int("consul.port", 8500, "consul port")
	serviceName := flag.String("service.name", "register", "service name")
	serviceAddr := flag.String("service.addr", "localhost", "service addr")
	servicePort := flag.Int("service.port", 12312, "service port")
	flag.Parse()
	client := discovery.NewDiscoveryClient(*consulAddr, *consulPort)
	errChan := make(chan error)
	srv := service.NewRegisterServiceImpl(client)
	endpoints := endpoint.RegisterEndpoints{
		DiscoveryEndpoint: endpoint.MakeDiscoveryEndpoint(srv),
		HealthCheckEndpoint: endpoint.MakeDiscoveryEndpoint(srv),
	}
	handler := transport.MakeHttpHandler(context.Background(), &endpoints)
	go func() {
		errChan <- http.ListenAndServe(":" + strconv.Itoa(*servicePort), handler)
	}()
	go func() {
		// 监控信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	instanceId := *serviceName + "-" + uuid.New().String()
}
