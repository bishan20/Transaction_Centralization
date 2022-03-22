package repository

import "Centralized_transaction/models"

type TransactionRepository interface {
	Save(models.Transaction) (models.Transaction, error)
	FindAll(models.Transaction) ([]models.Transaction, error)
	// FindById(uint64) (models.Transaction, error)
	// Update(uint64, models.Transaction) (int64, error)
	// Delete(transaction_id uint64, user_id uint32) (int64, error)
}
