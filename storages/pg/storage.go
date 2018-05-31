package pg

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kirill-a-belov/test-sg/models"
)

type PGStorage struct {
	db *gorm.DB
}

func NewPGStorage(config *PostgresConfig) (*PGStorage, error) {
	ConnString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName)

	db, err := gorm.Open("postgres", ConnString)

	if err != nil {
		return nil, fmt.Errorf("DB opening error: %v", err)
	}

	return &PGStorage{db: db}, nil
}

func (pgs *PGStorage) Close() {
	pgs.db.Close()
}

func (pgs *PGStorage) GetDB() *sql.DB {
	return pgs.db.DB()
}

func (pgs *PGStorage) GetProduct(PID string) (*models.Product, error) {
	var result = &models.Product{}
	if res := pgs.db.First(result, "p_id = ?", PID); res.Error != nil {
		return nil, fmt.Errorf("error in method GetDP: %v", res.Error)
	}

	return result, nil
}

func (pgs *PGStorage) SaveProduct(product *models.Product) error {
	if res := pgs.db.Create(product); res.Error != nil {
		return fmt.Errorf("error in method GetDP: %v", res.Error)
	}

	return nil
}
