package service

import (
	"bitbucket.org/go-webservice/dto"
	"bitbucket.org/go-webservice/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (service *UserService) GetAll() []dto.UserResponseBody {
	var collection []dto.UserResponseBody
	for _, user := range service.UserRepository.FindAll() {
		userDTO := new(dto.UserResponseBody)
		userDTO = userDTO.ConvertUserEntityToDTO(user)
		collection = append(collection, *userDTO)
	}
	return collection
}

func (service *UserService) Get(id int) (dto.UserResponseBody, error) {
	userDTO := new(dto.UserResponseBody)
	user, err := service.UserRepository.FindByID(id)
	userDTO = userDTO.ConvertUserEntityToDTO(user)

	return *userDTO, err
}
