package service

import (
	"bitbucket.org/go-webservice/dto"
	"bitbucket.org/go-webservice/model"
	"bitbucket.org/go-webservice/repository"
)

type FavouriteService struct {
	FavouriteRepository repository.FavouriteRepository
	UserRepository      repository.UserRepository
}

func (service *FavouriteService) GetAll() []model.Favourite {
	return service.FavouriteRepository.FindAll()
}

func (service *FavouriteService) CreateFavourite(favouriteDTO model.Favourite) (int, error) {
	id, err := service.FavouriteRepository.CreateFavourite(favouriteDTO)
	return id, err
}

func (service *FavouriteService) GetByID(id int) (model.Favourite, error) {
	return service.FavouriteRepository.FindByID(id)
}

func (service *FavouriteService) GetByUserID(id int) (dto.UserFavouriteResponseBody, error) {

	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		return dto.UserFavouriteResponseBody{}, err
	}

	userDTO := &dto.UserFavouriteResponseBody{UserID: user.ID, Assets: []model.Asset{}}

	insCollection, _ := service.FavouriteRepository.FindFavouriteInsightsByUserID(id)
	varCollection, _ := service.FavouriteRepository.FindFavouriteAudiencesByUserID(id)
	chCollection, _ := service.FavouriteRepository.FindFavouriteChartsByUserID(id)

	for _, ins := range insCollection {
		userDTO.AppendChild(ins)
	}

	for _, v := range varCollection {
		userDTO.AppendChild(v)
	}

	for _, c := range chCollection {
		userDTO.AppendChild(c)
	}

	return *userDTO, nil
}

func (service *FavouriteService) DeleteByUserIDAndAssetID(userID, assetID int) (int, error) {
	return service.FavouriteRepository.DeleteFavouriteFromUser(userID, assetID)
}
