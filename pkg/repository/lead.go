package repository

import (
	"leadwebhook/pkg/domain"
)

type LeadRepository interface {
	ListLeads(webhookID string, offset, limit int) ([]domain.Lead, error)
	TotalLeads(webhookID string) (int, error)
	AddLead(data domain.Lead) (domain.Lead, error)
}
