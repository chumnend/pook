package book

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/pook/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo domain.BookRepository
}

func (s *Suite) SetupTest() {
	db, mock, err := sqlmock.New()
	require.NoError(s.T(), err)

	gdb, err := gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	repo := NewPostgresRepository(gdb)

	s.DB = gdb
	s.mock = mock
	s.repo = repo
}

func (s *Suite) TestFindAll() {
	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "Book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "Book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at"}).
		AddRow(mockBooks[0].ID, mockBooks[0].Title, mockBooks[0].CreatedAt, mockBooks[0].UpdatedAt).
		AddRow(mockBooks[1].ID, mockBooks[1].Title, mockBooks[1].CreatedAt, mockBooks[1].UpdatedAt)

	query := regexp.QuoteMeta(`SELECT * FROM "books"`)

	s.Run("success", func() {
		// setup
		s.mock.ExpectQuery(query).WillReturnRows(rows)

		// run
		books, err := s.repo.FindAll()

		// check
		s.NoError(err)
		s.Len(books, 2)
		s.mock.ExpectationsWereMet()
	})

	s.Run("fail", func() {
		// setup
		s.mock.ExpectQuery(query).WillReturnError(errors.New("Unexpected error"))

		// run
		books, err := s.repo.FindAll()

		// check
		s.Error(err)
		s.Len(books, 0)
		s.mock.ExpectationsWereMet()
	})
}

func (s *Suite) TestFindAllByUserID() {
	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "Book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "Book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at"}).
		AddRow(mockBooks[0].ID, mockBooks[0].Title, mockBooks[0].CreatedAt, mockBooks[0].UpdatedAt).
		AddRow(mockBooks[1].ID, mockBooks[1].Title, mockBooks[1].CreatedAt, mockBooks[1].UpdatedAt)

	query := regexp.QuoteMeta(`SELECT * FROM "books" WHERE (user_id = $1)`)

	s.Run("success", func() {
		// setup
		s.mock.ExpectQuery(query).WillReturnRows(rows)

		// run
		books, err := s.repo.FindAllByUserID(1)

		// check
		s.NoError(err)
		s.Len(books, 2)
		s.mock.ExpectationsWereMet()
	})

	s.Run("fail", func() {
		// setup
		s.mock.ExpectQuery(query).WillReturnError(errors.New("Unexpected error"))

		// run
		books, err := s.repo.FindAllByUserID(1)

		// check
		s.Error(err)
		s.Len(books, 0)
		s.mock.ExpectationsWereMet()
	})
}

func TestBookRepositorySuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
