package controllers

import (
	"github.com/indeedhat/potato/internal/store"
)

type Controllers struct {
	repo store.TheoryRepository
}

func New(repo store.TheoryRepository) *Controllers {
	return &Controllers{
		repo: repo,
	}
}
