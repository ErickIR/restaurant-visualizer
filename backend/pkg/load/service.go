package load

import (
	"encoding/json"
	"errors"
	"fmt"
	"restaurant-visualizer/pkg/integration"
	"time"
)

type DataLoader interface {
	LoadData(date string) error
}

type LoadService struct {
	repo            LoadRepo
	externalService integration.ExternalGetter
}

func NewService(r LoadRepo, es integration.ExternalGetter) *LoadService {
	return &LoadService{repo: r, externalService: es}
}

func (s *LoadService) LoadData(date string) error {
	if date == "" {
		date = fmt.Sprint(time.Now().Local().Unix())
	}

	isLoaded, err := s.repo.IsDateLoaded(date)

	if err != nil {
		return err
	}

	if isLoaded {
		return errors.New("date is loaded")
	}

	err = loadBuyers(s.repo, s.externalService, date)

	if err != nil {
		return err
	}

	err = loadProducts(s.repo, s.externalService, date)

	if err != nil {
		return err
	}

	err = loadTransactions(s.repo, s.externalService, date)

	if err != nil {
		return err
	}

	return nil
}

func loadBuyers(repo LoadRepo, externalService integration.ExternalGetter, date string) error {
	buyers, err := externalService.GetBuyers(date)

	if err != nil {
		return err
	}

	filteredBuyersList, err := repo.FilterDuplicateBuyers(buyers)

	if err != nil {
		return err
	}

	if filteredBuyersList != nil {
		json, err := json.Marshal(filteredBuyersList)

		if err != nil {
			return err
		}

		err = repo.SaveData(json)

		if err != nil {
			return err
		}
	}

	return nil
}

func loadProducts(repo LoadRepo, externalService integration.ExternalGetter, date string) error {
	products, err := externalService.GetProducts(date)

	if err != nil {
		return err
	}

	filteredProductsList, err := repo.FilterDuplicateProducts(products)

	if err != nil {
		return err
	}

	if filteredProductsList != nil {
		json, err := json.Marshal(filteredProductsList)

		if err != nil {
			return err
		}

		err = repo.SaveData(json)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadTransactions(repo LoadRepo, externalService integration.ExternalGetter, date string) error {
	transactions, err := externalService.GetTransactions(date)

	if err != nil {
		return err
	}

	filteredTransactions, err := repo.FilterDuplicateTransactions(transactions)

	if err != nil {
		return err
	}

	if filteredTransactions != nil {
		json, err := json.Marshal(filteredTransactions)

		if err != nil {
			return err
		}

		err = repo.SaveData(json)
		if err != nil {
			return err
		}

		err = repo.SetBuyersToTransactions(filteredTransactions)

		if err != nil {
			return err
		}

		err = repo.SetProductsToTransactions(filteredTransactions)

		if err != nil {
			return err
		}
	}

	return nil
}
