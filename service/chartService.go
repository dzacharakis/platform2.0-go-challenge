package service

import (
	"bitbucket.org/go-webservice/model"
	"bitbucket.org/go-webservice/repository"
)

type ChartService struct {
	ChartRepository     repository.ChartRepository
	ChartDataRepository repository.ChartDataRepository
}

func (service *ChartService) GetAll() []model.Chart {
	return service.ChartRepository.FindAll()
}

func (service *ChartService) Get(id int) (model.Chart, error) {
	var chart model.Chart
	chart, err := service.ChartRepository.FindByID(id)

	if err != nil {
		return chart, err
	}

	chartDataCollection, _ := service.ChartDataRepository.FindDataByChartID(chart.AssetID)
	chart.Data = chartDataCollection

	return chart, err
}
