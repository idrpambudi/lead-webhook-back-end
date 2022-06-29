package domain

import (
	"time"

	"gorm.io/datatypes"
)

type Lead struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	WebhookID string `json:"-"`
	Name      string
	Email     string
	Phone     string
	Others    datatypes.JSON
}
