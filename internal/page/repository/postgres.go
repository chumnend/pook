package repository

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/jinzhu/gorm"
)

type pageRepo struct {
	db *gorm.DB
}

// NewPostgresRepository returns a PageRepository struct utilizing PostgreSQL
func NewPostgresRepository(db *gorm.DB) domain.PageRepository {
	return &pageRepo{db: db}
}

func (repo *pageRepo) FindAllByBookID(id uint) ([]domain.Page, error) {
	var pages []domain.Page
	result := repo.db.Where("book_id = ?", id).Find(&pages)
	if result.Error != nil {
		return pages, result.Error
	}
	return pages, nil
}

func (repo *pageRepo) FindByID(id uint) (*domain.Page, error) {
	var page domain.Page
	result := repo.db.First(&page, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &page, nil
}

func (repo *pageRepo) Save(page *domain.Page) error {
	result := repo.db.Create(page)
	return result.Error
}

func (repo *pageRepo) Delete(page *domain.Page) error {
	result := repo.db.Delete(page)
	return result.Error
}

func (repo *pageRepo) Migrate() error {
	return repo.db.AutoMigrate(&domain.Page{}).Error
}
