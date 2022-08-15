package controller

import (
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

}