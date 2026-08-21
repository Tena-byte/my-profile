package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"portfolio-os/internal/models"
)

type PortfolioService struct {
	mu        sync.RWMutex
	portfolio *models.Portfolio
	path      string
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
		path:      path,
	}, nil
}

func (s *PortfolioService) GetPortfolio() *models.Portfolio {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.portfolio
}

func (s *PortfolioService) UpdateProfile(
	profile models.Profile,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Profile = profile

	return s.save()
}

func (s *PortfolioService) save() error {
	file, err := os.Create(s.path)
	if err != nil {
		return fmt.Errorf("create portfolio data: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(s.portfolio); err != nil {
		return fmt.Errorf("encode portfolio data: %w", err)
	}

	return nil
}

func (s *PortfolioService) AddCertificate(
	certificate models.Certificate,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Certificates = append(
		s.portfolio.Certificates,
		certificate,
	)

	return s.save()
}

func (s *PortfolioService) UpdateCertificate(
	index int,
	certificate models.Certificate,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Certificates) {
		return fmt.Errorf("certificate not found")
	}

	s.portfolio.Certificates[index] = certificate

	return s.save()
}

func (s *PortfolioService) DeleteCertificate(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Certificates) {
		return fmt.Errorf("certificate not found")
	}

	s.portfolio.Certificates = append(
		s.portfolio.Certificates[:index],
		s.portfolio.Certificates[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) AddSkill(
	skill models.Skill,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Skills = append(
		s.portfolio.Skills,
		skill,
	)

	return s.save()
}

func (s *PortfolioService) UpdateSkill(
	index int,
	skill models.Skill,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Skills) {
		return fmt.Errorf("skill not found")
	}

	s.portfolio.Skills[index] = skill

	return s.save()
}

func (s *PortfolioService) DeleteSkill(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Skills) {
		return fmt.Errorf("skill not found")
	}

	s.portfolio.Skills = append(
		s.portfolio.Skills[:index],
		s.portfolio.Skills[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) AddStatistic(
	statistic models.Statistic,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Statistics = append(
		s.portfolio.Statistics,
		statistic,
	)

	return s.save()
}

func (s *PortfolioService) UpdateStatistic(
	index int,
	statistic models.Statistic,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Statistics) {
		return fmt.Errorf("statistic not found")
	}

	s.portfolio.Statistics[index] = statistic

	return s.save()
}

func (s *PortfolioService) DeleteStatistic(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Statistics) {
		return fmt.Errorf("statistic not found")
	}

	s.portfolio.Statistics = append(
		s.portfolio.Statistics[:index],
		s.portfolio.Statistics[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) AddService(
	service models.ServiceOffering,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Services = append(
		s.portfolio.Services,
		service,
	)

	return s.save()
}

func (s *PortfolioService) UpdateService(
	index int,
	service models.ServiceOffering,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Services) {
		return fmt.Errorf("service not found")
	}

	s.portfolio.Services[index] = service

	return s.save()
}

func (s *PortfolioService) DeleteService(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Services) {
		return fmt.Errorf("service not found")
	}

	s.portfolio.Services = append(
		s.portfolio.Services[:index],
		s.portfolio.Services[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) AddEducation(
	entry models.EducationEntry,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Education = append(
		s.portfolio.Education,
		entry,
	)

	return s.save()
}

func (s *PortfolioService) UpdateEducation(
	index int,
	entry models.EducationEntry,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Education) {
		return fmt.Errorf("education entry not found")
	}

	s.portfolio.Education[index] = entry

	return s.save()
}

func (s *PortfolioService) DeleteEducation(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Education) {
		return fmt.Errorf("education entry not found")
	}

	s.portfolio.Education = append(
		s.portfolio.Education[:index],
		s.portfolio.Education[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) AddProject(
	project models.Project,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Projects = append(
		s.portfolio.Projects,
		project,
	)

	return s.save()
}

func (s *PortfolioService) UpdateProject(
	index int,
	project models.Project,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Projects) {
		return fmt.Errorf("project not found")
	}

	s.portfolio.Projects[index] = project

	return s.save()
}

func (s *PortfolioService) DeleteProject(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.Projects) {
		return fmt.Errorf("project not found")
	}

	s.portfolio.Projects = append(
		s.portfolio.Projects[:index],
		s.portfolio.Projects[index+1:]...,
	)

	return s.save()
}

func (s *PortfolioService) UpdateContact(
	contact models.ContactInfo,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.Contact = contact

	return s.save()
}

func (s *PortfolioService) AddSocialLink(
	link models.SocialLink,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.portfolio.SocialLinks = append(
		s.portfolio.SocialLinks,
		link,
	)

	return s.save()
}

func (s *PortfolioService) UpdateSocialLink(
	index int,
	link models.SocialLink,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.SocialLinks) {
		return fmt.Errorf("social link not found")
	}

	s.portfolio.SocialLinks[index] = link

	return s.save()
}

func (s *PortfolioService) DeleteSocialLink(
	index int,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.portfolio.SocialLinks) {
		return fmt.Errorf("social link not found")
	}

	s.portfolio.SocialLinks = append(
		s.portfolio.SocialLinks[:index],
		s.portfolio.SocialLinks[index+1:]...,
	)

	return s.save()
}
