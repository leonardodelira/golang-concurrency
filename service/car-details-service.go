package service

type CarDetailsService interface {
	GetDetails()
}

type service struct {}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() {}