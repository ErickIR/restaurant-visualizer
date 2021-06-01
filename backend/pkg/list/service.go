package list

import "restaurant-visualizer/pkg/models"

type DataEnlister interface {
	GetAllBuyers(page, size int) ([]models.Buyer, error)
	// GetAllBuyers(buyerId string) (interface{}, error)
	GetBuyersCount() (int, error)
}

type ListService struct {
	repo ListRepo
}

func NewService(r ListRepo) *ListService {
	return &ListService{repo: r}
}

func (s *ListService) GetAllBuyers(page, size int) ([]models.Buyer, error) {
	resp, err := s.repo.GetAllBuyers(page, size)

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
