package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/chumnend/pook/internal/book/repository"
	"github.com/chumnend/pook/internal/domain"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestRepo_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "test book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "test book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "title", "created_at", "updated_at", "user_id"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockBooks[0].ID, mockBooks[0].Title, mockBooks[0].CreatedAt, mockBooks[0].UpdatedAt, mockBooks[0].UserID).
			AddRow(mockBooks[1].ID, mockBooks[1].Title, mockBooks[1].CreatedAt, mockBooks[1].UpdatedAt, mockBooks[1].UserID)
		query := regexp.QuoteMeta(`SELECT * FROM "books"`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := repository.NewPostgresRepository(gdb)

		// run
		books, err := repo.FindAll()

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, books, 2)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "books"`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := repository.NewPostgresRepository(gdb)

		// run
		books, err := repo.FindAll()

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, books, 0)
		assert.Error(t, err)
	})
}

func TestRepo_FindAllByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	mockBooks := []domain.Book{
		domain.Book{
			ID:        1,
			Title:     "test book 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
		domain.Book{
			ID:        2,
			Title:     "test book 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "title", "created_at", "updated_at", "user_id"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockBooks[0].ID, mockBooks[0].Title, mockBooks[0].CreatedAt, mockBooks[0].UpdatedAt, mockBooks[0].UserID).
			AddRow(mockBooks[1].ID, mockBooks[1].Title, mockBooks[1].CreatedAt, mockBooks[1].UpdatedAt, mockBooks[1].UserID)
		query := regexp.QuoteMeta(`SELECT * FROM "books" WHERE (user_id = $1)`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := repository.NewPostgresRepository(gdb)

		// run
		books, err := repo.FindAllByUserID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, books, 2)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "books" WHERE (user_id = $1)`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := repository.NewPostgresRepository(gdb)

		// run
		books, err := repo.FindAllByUserID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, books, 0)
		assert.Error(t, err)
	})
}

func TestRepo_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	mockBook := domain.Book{
		ID:        1,
		Title:     "test book 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "title", "created_at", "updated_at", "user_id"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockBook.ID, mockBook.Title, mockBook.CreatedAt, mockBook.UpdatedAt, mockBook.UserID)
		query := regexp.QuoteMeta(`SELECT * FROM "books" WHERE ("books"."id" = 1) ORDER BY "books"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := repository.NewPostgresRepository(gdb)

		// run
		book, err := repo.FindByID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, mockBook.ID, book.ID)
		assert.Equal(t, mockBook.Title, book.Title)
		assert.Equal(t, mockBook.UserID, book.UserID)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "books" WHERE ("books"."id" = 1) ORDER BY "books"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := repository.NewPostgresRepository(gdb)

		// run
		book, err := repo.FindByID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, domain.Book{}.ID, book.ID)
		assert.Equal(t, domain.Book{}.Title, book.Title)
		assert.Equal(t, domain.Book{}.UserID, book.UserID)
		assert.Error(t, err)
	})
}

func TestRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	book := domain.Book{
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id"}
		rows := sqlmock.NewRows(headers).AddRow(1)
		query := regexp.QuoteMeta(`INSERT INTO "books" ("title","created_at","updated_at","user_id") VALUES ($1,$2,$3,$4) RETURNING "books"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(book.Title, book.CreatedAt, book.UpdatedAt, book.UserID).
			WillReturnRows(rows)
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Create(&book)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`INSERT INTO "books" ("title","created_at","updated_at","user_id") VALUES ($1,$2,$3,$4) RETURNING "books"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(book.Title, book.CreatedAt, book.UpdatedAt, book.UserID).
			WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Create(&book)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}

func TestRepo_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	book := domain.Book{
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id"}
		rows := sqlmock.NewRows(headers).AddRow(1)
		query := regexp.QuoteMeta(`INSERT INTO "books" ("title","created_at","updated_at","user_id") VALUES ($1,$2,$3,$4) RETURNING "books"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(book.Title, book.CreatedAt, book.UpdatedAt, book.UserID).
			WillReturnRows(rows)
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Save(&book)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`INSERT INTO "books" ("title","created_at","updated_at","user_id") VALUES ($1,$2,$3,$4) RETURNING "books"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(book.Title, book.CreatedAt, book.UpdatedAt, book.UserID).
			WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Save(&book)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}

func TestRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	book := domain.Book{
		ID:        1,
		Title:     "test book",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`DELETE FROM "books" WHERE "books"."id" = $1`)
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Delete(&book)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`DELETE FROM "books" WHERE "books"."id" = $1`)
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(1).WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit()
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Delete(&book)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}
