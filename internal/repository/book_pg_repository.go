package repository

import (
	"errors"

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
	return []entity.Book{}, errors.New("not yet implemented")
}

func (b *bookRepository) FindAllByUserID(uint) ([]entity.Book, error) {
	return []entity.Book{}, errors.New("not yet implemented")
}

func (b *bookRepository) FindByID(uint) (*entity.Book, error) {
	return &entity.Book{}, errors.New("not yet implemented")
}

func (b *bookRepository) Create(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookRepository) Save(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookRepository) Delete(*entity.Book) error {
	return errors.New("not yet implemented")
}

func (b *bookRepository) Migrate() error {
	return b.conn.AutoMigrate(&entity.Book{}).Error // TODO: missing tests
}
