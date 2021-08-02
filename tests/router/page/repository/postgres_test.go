package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/router/page/repository"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestRepo_FindAllByBookID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	mockPages := []domain.Page{
		domain.Page{
			ID:        1,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
		domain.Page{
			ID:        2,
			Content:   "page content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			BookID:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "content", "created_at", "updated_at", "book_id"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockPages[0].ID, mockPages[0].Content, mockPages[0].CreatedAt, mockPages[0].UpdatedAt, mockPages[0].BookID).
			AddRow(mockPages[1].ID, mockPages[1].Content, mockPages[1].CreatedAt, mockPages[1].UpdatedAt, mockPages[1].BookID)
		query := regexp.QuoteMeta(`SELECT * FROM "pages" WHERE (book_id = $1)`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := repository.NewPostgresRepository(gdb)

		// run
		pages, err := repo.FindAllByBookID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, pages, 2)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "pages" WHERE (book_id = $1)`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := repository.NewPostgresRepository(gdb)

		// run
		books, err := repo.FindAllByBookID(1)

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

	mockPage := domain.Page{
		ID:        1,
		Content:   "page content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		headers := []string{"id", "content", "created_at", "updated_at", "book_id"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockPage.ID, mockPage.Content, mockPage.CreatedAt, mockPage.UpdatedAt, mockPage.BookID)
		query := regexp.QuoteMeta(`SELECT * FROM "pages" WHERE ("pages"."id" = 1) ORDER BY "pages"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := repository.NewPostgresRepository(gdb)

		// run
		page, err := repo.FindByID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, mockPage.ID, page.ID)
		assert.Equal(t, mockPage.Content, page.Content)
		assert.Equal(t, mockPage.BookID, page.BookID)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "pages" WHERE ("pages"."id" = 1) ORDER BY "pages"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := repository.NewPostgresRepository(gdb)

		// run
		page, err := repo.FindByID(1)

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, domain.Page{}.ID, page.ID)
		assert.Equal(t, domain.Page{}.Content, page.Content)
		assert.Equal(t, domain.Page{}.BookID, page.BookID)
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

	page := domain.Page{
		Content:   "test page",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id"}
		rows := sqlmock.NewRows(headers).AddRow(1)
		query := regexp.QuoteMeta(`INSERT INTO "pages" ("content","created_at","updated_at","book_id") VALUES ($1,$2,$3,$4) RETURNING "pages"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(page.Content, page.CreatedAt, page.UpdatedAt, page.BookID).
			WillReturnRows(rows)
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Create(&page)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`INSERT INTO "pages" ("content","created_at","updated_at","book_id") VALUES ($1,$2,$3,$4) RETURNING "pages"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(page.Content, page.CreatedAt, page.UpdatedAt, page.BookID).
			WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Create(&page)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}

func TestRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	page := domain.Page{
		Content:   "test page",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id"}
		rows := sqlmock.NewRows(headers).AddRow(1)
		query := regexp.QuoteMeta(`INSERT INTO "pages" ("content","created_at","updated_at","book_id") VALUES ($1,$2,$3,$4) RETURNING "pages"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(page.Content, page.CreatedAt, page.UpdatedAt, page.BookID).
			WillReturnRows(rows)
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Update(&page)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`INSERT INTO "pages" ("content","created_at","updated_at","book_id") VALUES ($1,$2,$3,$4) RETURNING "pages"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(page.Content, page.CreatedAt, page.UpdatedAt, page.BookID).
			WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit() // commit transaction
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Update(&page)

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

	page := domain.Page{
		ID:        1,
		Content:   "test page",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		BookID:    1,
	}

	t.Run("success", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`DELETE FROM "pages" WHERE "pages"."id" = $1`)
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Delete(&page)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`DELETE FROM "pages" WHERE "pages"."id" = $1`)
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(1).WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit()
		repo := repository.NewPostgresRepository(gdb)

		// run
		err := repo.Delete(&page)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}
