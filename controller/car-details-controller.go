package controller

import (
	"encoding/json"
	"go_concurrency/service"
	"net/http"
)

type controller struct {}

type CarDetailsControler interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}

var carDetailService service.CarDetailsService

func NewCarDetailsControler(s service.CarDetailsService) CarDetailsControler {
	carDetailService = s
	return &controller{}
}

func (*controller) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	respApi := carDetailService.GetDetails()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(respApi)
}