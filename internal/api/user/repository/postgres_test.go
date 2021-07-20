package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chumnend/pook/internal/api/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
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

	mockUsers := []domain.User{
		domain.User{
			ID:        1,
			Email:     "tester@pook.com",
			Password:  "123",
			FirstName: "tester",
			LastName:  "tester",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "email", "password", "first_name", "last_name", "created_at", "updated_at"}
		rows := sqlmock.NewRows(headers).
			AddRow(
				mockUsers[0].ID,
				mockUsers[0].Email,
				mockUsers[0].Password,
				mockUsers[0].FirstName,
				mockUsers[0].LastName,
				mockUsers[0].CreatedAt,
				mockUsers[0].UpdatedAt,
			)
		query := regexp.QuoteMeta(`SELECT * FROM "users"`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := NewPostgresRepository(gdb)

		// run
		users, err := repo.FindAll()

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, users, 1)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "users"`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := NewPostgresRepository(gdb)

		// run
		users, err := repo.FindAll()

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, users, 0)
		assert.Error(t, err)
	})
}

func TestRepo_FindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal("an error occured when opening stub database", err)
	}

	mockUser := domain.User{
		ID:        1,
		Email:     "tester@pook.com",
		Password:  "123",
		FirstName: "tester",
		LastName:  "tester",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id", "email", "password", "first_name", "last_name", "created_at", "updated_at"}
		rows := sqlmock.NewRows(headers).
			AddRow(mockUser.ID, mockUser.Email, mockUser.Password, mockUser.FirstName, mockUser.LastName, mockUser.CreatedAt, mockUser.UpdatedAt)
		query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE (email = $1) ORDER BY "users"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnRows(rows)
		repo := NewPostgresRepository(gdb)

		// run
		user, err := repo.FindByEmail("tester@pook.com")

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, mockUser.ID, user.ID)
		assert.Equal(t, mockUser.Email, user.Email)
		assert.Equal(t, mockUser.FirstName, user.FirstName)
		assert.Equal(t, mockUser.LastName, user.LastName)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE (email = $1) ORDER BY "users"."id" ASC LIMIT 1`)
		mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
		repo := NewPostgresRepository(gdb)

		// run
		user, err := repo.FindByEmail("tester@pook.com")

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, domain.User{}.ID, user.ID)
		assert.Equal(t, domain.User{}.Email, user.Email)
		assert.Equal(t, domain.User{}.FirstName, user.FirstName)
		assert.Equal(t, domain.User{}.LastName, user.LastName)
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

	user := domain.User{
		Email:     "tester@pook.com",
		Password:  "123",
		FirstName: "tester",
		LastName:  "tester",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		// setup
		headers := []string{"id"}
		rows := sqlmock.NewRows(headers).AddRow(1)
		query := regexp.QuoteMeta(`INSERT INTO "users" ("email","password","first_name","last_name","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "users"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(user.Email, user.Password, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).
			WillReturnRows(rows)
		mock.ExpectCommit() // commit transaction
		repo := NewPostgresRepository(gdb)

		// run
		err := repo.Save(&user)

		// check
		mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// setup
		query := regexp.QuoteMeta(`INSERT INTO "users" ("email","password","first_name","last_name","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "users"."id"`)
		mock.ExpectBegin() // begin transaction
		mock.ExpectQuery(query).
			WithArgs(user.Email, user.Password, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).
			WillReturnError(errors.New("unexpected error"))
		mock.ExpectCommit() // commit transaction
		repo := NewPostgresRepository(gdb)

		// run
		err := repo.Save(&user)

		// check
		mock.ExpectationsWereMet()
		assert.Error(t, err)
	})
}
