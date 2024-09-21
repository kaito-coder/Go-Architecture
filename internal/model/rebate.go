package model

import (
	"time"
)

// Rebate represents the rebate structure
type Rebate struct {
	ID                    int        `json:"id"`
	ProductManufacturerID int        `json:"product_manufacturer_id"`
	Name                  string     `json:"name"`
	Code                  string     `json:"code"`
	Description           string     `json:"description"`
	StartDate             time.Time  `json:"start_date"`
	EndDate               time.Time  `json:"end_date"`
	MaxUses               *int       `json:"max_uses,omitempty"` // nullable int
	MixedAvailability     bool       `json:"mixed_availability"`
	State                 int        `json:"state"` // 0: inactive, 1: expired, 2: active
	ActiveTime            *time.Time `json:"active_time,omitempty"`
	IsCombinable          int        `json:"is_combinable"` // 0: no, 1: yes
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

