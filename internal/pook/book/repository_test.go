package book

import (
	"regexp"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindAll(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockBooks := []domain.Book{
		domain.Book{
			ID: 1, Title: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: 1,
		},
		domain.Book{
			ID: 2, Title: "Book 2", CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: 2,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at"}).
		AddRow(mockBooks[0].ID, mockBooks[0].Title, mockBooks[0].CreatedAt, mockBooks[0].UpdatedAt).
		AddRow(mockBooks[1].ID, mockBooks[1].Title, mockBooks[1].CreatedAt, mockBooks[1].UpdatedAt)
	query := regexp.QuoteMeta(`SELECT * FROM "books"`)
	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := NewPostgresRepository(mockDB)

	// run
	books, err := repo.FindAll()

	// check
	assert.NoError(t, err)
	assert.Len(t, books, 2)

	mock.ExpectationsWereMet()
}
