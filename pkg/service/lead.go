package service

import (
	"fmt"
	utils "leadwebhook/pkg"
	"leadwebhook/pkg/domain"
	"leadwebhook/pkg/repository"
)

type PaginatedLeads struct {
	Data  []domain.Lead
	Total int
}

type LeadService struct {
	leadRepository repository.LeadRepository
}

func NewLeadService(repo repository.LeadRepository) *LeadService {
	return &LeadService{repo}
}

func (s *LeadService) ListLeads(userID string, offset, limit int) (PaginatedLeads, error) {
	webhookID := utils.GetWebhookID(userID)
	leads, err := s.leadRepository.ListLeads(webhookID, offset, limit)
	if err != nil {
		return PaginatedLeads{}, err
	}
	total, err := s.leadRepository.TotalLeads(webhookID)
	if err != nil {
		return PaginatedLeads{}, err
	}
	return PaginatedLeads{Data: leads, Total: total}, nil
}

func (s *LeadService) AddLead(data domain.Lead) (domain.Lead, error) {
	return s.leadRepository.AddLead(data)
}

func (s *LeadService) GetWebhookID(userID string) string {
	return fmt.Sprintf("/webhook/%s/lead", utils.GetWebhookID(userID))
}
