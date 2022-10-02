package service

import (
	"errors"
	"rafit/models"
	"rafit/repository"
)

type Service interface {
	CreateIP(req models.ReqIpModel) (*models.ResIpModel, error)
	GetAllIp() ([]*models.ResIpModel, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) CreateIP(req models.ReqIpModel) (*models.ResIpModel, error) {
	data := models.ReqFormatterIP(req)

	duplicate, err := s.repository.FindByIp(req.IP)
	if err != nil {
		return nil, err
	}

	if len(duplicate.IP) > 0 {
		return nil, errors.New("duplicate ip")
	}

	result, err := s.repository.CreateIp(data)
	if err != nil {
		return nil, err
	}

	response := models.ResFormatterIP(*result)
	return &response, nil
}
func (s *service) GetAllIp() ([]*models.ResIpModel, error) {
	data, err := s.repository.GetAllIp()
	if err != nil {
		return nil, err
	}

	results := models.GetAllFormatterIP(data)
	return results, nil
}
