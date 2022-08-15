package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct {}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %v", ownerServiceURL)
	resp, _ := client.Get(ownerServiceURL)
	ownerDataChannel <- resp
}