package book

import (
	"github.com/chumnend/pook/internal/domain"
	"github.com/jinzhu/gorm"
)

type bookRepo struct {
	conn *gorm.DB
}

// NewPostgresRepository returns a BookRepository struct utilizing PostgreSQL
func NewPostgresRepository(conn *gorm.DB) domain.BookRepository {
	return &bookRepo{conn: conn}
}

func (repo *bookRepo) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	result := repo.conn.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (repo *bookRepo) FindAllByUserID(id uint) ([]domain.Book, error) {
	var books []domain.Book
	result := repo.conn.Where("user_id = ?", id).Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (repo *bookRepo) FindByID(id uint) (*domain.Book, error) {
	var book domain.Book
	result := repo.conn.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (repo *bookRepo) Save(book *domain.Book) error {
	result := repo.conn.Create(book)
	return result.Error
}

func (repo *bookRepo) Delete(book *domain.Book) error {
	result := repo.conn.Delete(book)
	return result.Error
}

func (repo *bookRepo) Migrate() error {
	return repo.conn.AutoMigrate(&domain.Book{}).Error
}
