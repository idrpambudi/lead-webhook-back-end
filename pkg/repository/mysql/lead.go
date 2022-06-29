package mysqlrepo

import (
	"leadwebhook/pkg/domain"
	"leadwebhook/pkg/repository"

	"gorm.io/gorm"
)

type MySqlLeadRepository struct {
	client *gorm.DB
}

func NewMySqlLeadRepository(client *gorm.DB) repository.LeadRepository {
	return &MySqlLeadRepository{client}
}

func (r *MySqlLeadRepository) AddLead(data domain.Lead) (domain.Lead, error) {
	if err := r.client.Create(&data).Error; err != nil {
		return domain.Lead{}, err
	}
	return data, nil
}

func (r *MySqlLeadRepository) ListLeads(webhookID string, offset, limit int) ([]domain.Lead, error) {
	var leads []domain.Lead
	if err := r.client.Limit(limit).Offset(offset).Where("webhook_id = ?", webhookID).Find(&leads).Error; err != nil {
		return leads, err
	}
	return leads, nil
}

func (r *MySqlLeadRepository) TotalLeads(webhookID string) (int, error) {
	var count int64

	if err := r.client.Model(&domain.Lead{}).Where("webhook_id = ?", webhookID).Count(&count).Error; err != nil {
		return int(count), err
	}
	return int(count), nil
}
