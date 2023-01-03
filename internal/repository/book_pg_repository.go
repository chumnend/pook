package repository

import (
	"github.com/chumnend/pook/internal/entity"
	"github.com/jinzhu/gorm"
)

type bookRepository struct {
	conn *gorm.DB
}

// NewBookPostgresRepository returns a BookRepository struct utilizing PostgreSQL
func NewBookPostgresRepository(conn *gorm.DB) entity.BookRepository {
	return &bookRepository{conn: conn}
}

func (repo *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	result := repo.conn.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (repo *bookRepository) FindAllByUserID(id uint) ([]entity.Book, error) {
	var books []entity.Book
	result := repo.conn.Where("user_id = ?", id).Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (repo *bookRepository) FindByID(id uint) (*entity.Book, error) {
	var book entity.Book
	result := repo.conn.First(&book, id)
	if result.Error != nil {
		return &book, result.Error
	}
	return &book, nil
}

func (repo *bookRepository) Create(book *entity.Book) error {
	result := repo.conn.Create(book)
	return result.Error
}

func (repo *bookRepository) Save(book *entity.Book) error {
	result := repo.conn.Save(book)
	return result.Error
}

func (repo *bookRepository) Delete(book *entity.Book) error {
	result := repo.conn.Delete(book)
	return result.Error
}

func (repo *bookRepository) Migrate() error {
	return repo.conn.AutoMigrate(&entity.Book{}).Error // TODO: missing tests
}
