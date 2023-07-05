package store

import (
	"errors"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

type Theory struct {
	gorm.Model

	Text string
}

type TheoryRepository interface {
	Create(text string) error
	GetRandom() (*Theory, error)
}

type TheorySqlRepo struct {
	db *gorm.DB
}

func NewTheorySqlRepo(db *gorm.DB) TheoryRepository {
	return &TheorySqlRepo{db: db}
}

// Create implements TheoryRepository
func (r *TheorySqlRepo) Create(text string) error {
	theory := Theory{
		Text: text,
	}

	return r.db.Create(&theory).Error
}

// GetRandom implements TheoryRepository
func (r *TheorySqlRepo) GetRandom() (*Theory, error) {
	var (
		count  int64
		theory Theory
	)
	r.db.Model(&theory).Count(&count)

	if count == 0 {
		return nil, errors.New("no theories")
	}

	offset := rand.Intn(int(count))
	log.Printf("%d/%d", offset, count)

	if tx := r.db.Model(theory).Offset(offset).First(&theory); tx.Error != nil {
		return nil, tx.Error
	}

	return &theory, nil
}

var _ TheoryRepository = (*TheorySqlRepo)(nil)
