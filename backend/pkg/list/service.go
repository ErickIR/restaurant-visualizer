package list

import (
	"errors"
	dtos "restaurant-visualizer/pkg/dtos/out"
	"restaurant-visualizer/pkg/models"
)

type DataEnlister interface {
	GetAllBuyers(page, size int) ([]models.Buyer, error)
	GetBuyerInformation(buyerId string) (*dtos.BuyerInfo, error)
	GetBuyersCount() (int, error)
}

type ListService struct {
	repo ListRepo
}

func NewService(r ListRepo) *ListService {
	return &ListService{repo: r}
}

func (s *ListService) GetAllBuyers(page, size int) ([]models.Buyer, error) {

	offset := ((page - 1) * size) + 1

	resp, err := s.repo.GetAllBuyers(offset, size)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ListService) GetBuyersCount() (int, error) {
	resp, err := s.repo.GetTotalBuyersCount()

	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (s *ListService) GetBuyerInformation(buyerId string) (*dtos.BuyerInfo, error) {
	resp, err := s.repo.GetBuyerInformation(buyerId)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("buyer not found")
	}

	return resp, nil
}
