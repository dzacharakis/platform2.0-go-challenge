package service

import (
	"bitbucket.org/go-webservice/model"
	"bitbucket.org/go-webservice/repository"
)

type AudienceService struct {
	AudienceRepository repository.AudienceRepository
}

func (service *AudienceService) GetAll() []model.Variable {
	return service.AudienceRepository.FindAll()
}

func (service *AudienceService) Get(id int) (model.Variable, error) {
	return service.AudienceRepository.FindByID(id)
}
