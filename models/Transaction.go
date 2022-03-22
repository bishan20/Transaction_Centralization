package models

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Citizenship string    `gorm:"size:20;not null" json:"citizenship"`
	From        string    `gorm:"size:30;not null" json:"from"`
	Type        string    `gorm:"size:30;not null" json:"type"`
	Author      User      `json:"author"`
	Amount      int64     `gorm:"not null" json:"amount"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (p *Transaction) Prepare() {
	p.ID = 0
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
func (p *Transaction) Validate() error {
	if p.Citizenship == "" {
		return errors.New("required required")
	}
	if p.From == "" {
		return errors.New("required transaction from")
	}
	if p.Type == "" {
		return errors.New("required transaction type")
	}
	if p.Amount < 1 {
		return errors.New("required amount")
	}
	return nil
}
