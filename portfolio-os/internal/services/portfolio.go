package services

import (
	"encoding/json"
	"fmt"
	"os"

	"portfolio-os/internal/models"
)

type PortfolioService struct {
	portfolio *models.Portfolio
}

func NewPortfolioService(path string) (*PortfolioService, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open portfolio data: %w", err)
	}
	defer file.Close()

	var portfolio models.Portfolio

	if err := json.NewDecoder(file).Decode(&portfolio); err != nil {
		return nil, fmt.Errorf("decode portfolio data: %w", err)
	}

	return &PortfolioService{
		portfolio: &portfolio,
	}, nil
}

func (s *PortfolioService) GetPortfolio() *models.Portfolio {
	return s.portfolio
}