package promotion_test

import (
	"github.com/innotechdevops/gomvc/database"
	"github.com/innotechdevops/gomvc/promotion"
	"testing"
)

var repo promotion.Repository

func init() {
	db := database.Connection("test.db")
	repo = promotion.NewRepository(db)
}

func TestAdd(t *testing.T) {
	promo := promotion.Promotion{Id: 2, Name: "Promotion 1"}
	err := repo.Add(promo)
	if err != nil {
		t.Error("Cannot add, error:", err)
	}
}

func TestUpdate(t *testing.T) {
	promo := promotion.Promotion{Id: 2, Name: "Promotion 1 Update"}
	err := repo.Update(1, promo)
	if err != nil {
		t.Error("Cannot update, error:", err)
	}
}

func TestDelete(t *testing.T) {
	err := repo.Delete(1)
	if err != nil {
		t.Error("Cannot delete, error:", err)
	}
}

func TestGetAll(t *testing.T) {
	promos := repo.GetAll()
	if len(promos) == 0 {
		t.Error("Cannot get all, error: is zero")
	}
}

func TestGetById(t *testing.T) {
	promo := repo.GetById(2)
	if promo.Id == 0 {
		t.Error("Cannot get by id, error: is zero")
	}
}