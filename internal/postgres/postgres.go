package postgres

import (
	"fmt"
	"log"

	"github.com/newcrab/MCT-2025-containers.git/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Visit struct {
	ID    uint   `gorm:"primaryKey"`
	IP    string `gorm:"not null"`
	Count int64  `gorm:"not null"`
}

type PostgresDB struct {
	pg *gorm.DB
}

func NewPostgresDB(cfg *config.Config) *PostgresDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s pgname=%s port=%s",
		cfg.PG.Host, cfg.PG.User, cfg.PG.Password, cfg.PG.Name, cfg.PG.Port)

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("pg connection error: %v", err)
	}

	log.Println("pg connection successful")

	return &PostgresDB{pg: pg}
}

func (s *PostgresDB) AutoMigrate() {
	s.pg.AutoMigrate(&Visit{})
}

func (s *PostgresDB) SaveVisit(ip string) error {
	var visit Visit
	result := s.pg.Where("ip = ?", ip).First(&visit)

	if result.Error == nil {
		return s.pg.Model(&Visit{}).Where("ip = ?", ip).Update("count", gorm.Expr("count + 1")).Error
	} else {
		newVisit := Visit{IP: ip, Count: 1}
		return s.pg.Create(&newVisit).Error
	}
}

func (s *PostgresDB) GetVisitsCount(ip string) (int64, error) {
	var visit Visit
	result := s.pg.Where("ip = ?", ip).First(&visit)
	if result.Error != nil {
		return 0, result.Error
	}
	return visit.Count, nil
}
