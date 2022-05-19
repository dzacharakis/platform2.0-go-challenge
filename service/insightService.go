package service

import (
	"bitbucket.org/go-webservice/model"
	"bitbucket.org/go-webservice/repository"
)

type InsightService struct {
	InsightRepository repository.InsightRepository
}

func (service *InsightService) GetAll() []model.Insight {
	return service.InsightRepository.FindAll()
}

func (service *InsightService) Get(id int) (model.Insight, error) {
	return service.InsightRepository.FindByID(id)
}
