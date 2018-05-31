package pg

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kirill-a-belov/test-sg/models"
)

func NewLocalStorage() (*PGStorage, error) {
	ConnString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"127.0.0.1",
		5432,
		"postgres",
		"postgres",
		"test",
	)

	db, err := gorm.Open("postgres", ConnString)

	if err != nil {
		return nil, fmt.Errorf("DB opening error: %v", err)
	}

	return &PGStorage{db: db}, nil
}

var testData = []*models.Product{
	{
		PID:   "00000000-0000-0000-0000-000000000001",
		Title: "Test product 1",
		Price: "100",
		Image: "http://some-url.com/1.jpg",
	},
	{
		PID:   "00000000-0000-0000-0000-000000000002",
		Title: "Test product 2",
		Price: "200",
		Image: "http://some-url.com/2.jpg",
	},
}

func TestGetProduct(t *testing.T) {
	s, err := NewLocalStorage()
	if err != nil {
		t.Fatal(err)
	}

	for _, td := range testData {
		res, err := s.GetProduct(td.PID)
		if err != nil {
			t.Error(
				"For input: ", td,
				"got error", err,
			)
			continue
		}

		if res.Title != td.Title ||
			res.Price != td.Price ||
			res.Image != td.Image ||
			res.PID != td.PID {

			t.Error(
				"For input: ", td,
				"got not equal", res,
			)
		}
	}
}

func TestSaveProduct(t *testing.T) {
	s, err := NewLocalStorage()
	if err != nil {
		t.Fatal(err)
	}

	s.db.AutoMigrate(&models.Product{})

	for _, td := range testData {
		if err := s.SaveProduct(td); err != nil {
			t.Error(
				"For input: ", td,
				"got error", err,
			)
			continue
		}
	}

	for _, td := range testData {
		res, err := s.GetProduct(td.PID)
		if err != nil {
			t.Error(
				"For input: ", td,
				"got error", err,
			)
			continue
		}

		if res.Title != td.Title ||
			res.Price != td.Price ||
			res.Image != td.Image ||
			res.PID != td.PID {

			t.Error(
				"For input: ", td,
				"got not equal", res,
			)
		}
	}
}
