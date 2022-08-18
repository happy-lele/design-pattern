package adapter

import "fmt"

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct{}

func (c *AWSClient) RunInstance(cpu, mum float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f", cpu, mum)
	return nil
}

type AWSClientAdapter struct {
	Client AWSClient
}

func (a *AWSClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

type AliyunClient struct{}

func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
