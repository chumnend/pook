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

func (b *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	result := b.conn.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (b *bookRepository) FindAllByUserID(id uint) ([]entity.Book, error) {
	var books []entity.Book
	result := b.conn.Where("user_id = ?", id).Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (b *bookRepository) FindByID(id uint) (*entity.Book, error) {
	var book entity.Book
	result := b.conn.First(&book, id)
	if result.Error != nil {
		return &book, result.Error
	}
	return &book, nil
}

func (b *bookRepository) Create(book *entity.Book) error {
	result := b.conn.Create(book)
	return result.Error
}

func (b *bookRepository) Save(book *entity.Book) error {
	result := b.conn.Save(book)
	return result.Error
}

func (b *bookRepository) Delete(book *entity.Book) error {
	result := b.conn.Delete(book)
	return result.Error
}

func (b *bookRepository) Migrate() error {
	return b.conn.AutoMigrate(&entity.Book{}).Error // TODO: missing tests
}
