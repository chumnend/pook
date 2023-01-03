package repository

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/jinzhu/gorm"
)

type pageRepository struct {
	conn *gorm.DB
}

// NewPagePostgresRepository returns a PageRepository struct utilizing PostgreSQL
func NewPagePostgresRepository(conn *gorm.DB) entity.PageRepository {
	return &pageRepository{conn: conn}
}

func (repo *pageRepository) FindAllByBookID(id uint) ([]entity.Page, error) {
	var pages []entity.Page
	result := repo.conn.Where("book_id = ?", id).Find(&pages)
	if result.Error != nil {
		return pages, result.Error
	}
	return pages, nil
}

func (repo *pageRepository) FindByID(id uint) (*entity.Page, error) {
	var page entity.Page
	result := repo.conn.First(&page, id)
	if result.Error != nil {
		return &page, result.Error
	}
	return &page, nil
}

func (repo *pageRepository) Create(page *entity.Page) error {
	result := repo.conn.Create(page)
	return result.Error
}

func (repo *pageRepository) Update(page *entity.Page) error {
	result := repo.conn.Save(page)
	return result.Error
}

func (repo *pageRepository) Delete(page *entity.Page) error {
	result := repo.conn.Delete(page)
	return result.Error
}

func (repo *pageRepository) Migrate() error {
	return repo.conn.AutoMigrate(&entity.Page{}).Error
}
