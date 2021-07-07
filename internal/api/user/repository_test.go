package user

import (
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
		testRepo := NewPostgresRepository(gdb)

		// run
		users, err := testRepo.FindAll()

		// check
		mock.ExpectationsWereMet()
		assert.Len(t, users, 1)
		assert.NoError(t, err)
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
			AddRow(
				mockUser.ID,
				mockUser.Email,
				mockUser.Password,
				mockUser.FirstName,
				mockUser.LastName,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
			)
		query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE (email = $1) ORDER BY "users"."id" ASC LIMIT 1`)

		mock.ExpectQuery(query).WillReturnRows(rows)
		testRepo := NewPostgresRepository(gdb)

		// run
		user, err := testRepo.FindByEmail("tester@pook.com")

		// check
		mock.ExpectationsWereMet()
		assert.Equal(t, mockUser.ID, user.ID)
		assert.Equal(t, mockUser.Email, user.Email)
		assert.Equal(t, mockUser.FirstName, user.FirstName)
		assert.Equal(t, mockUser.LastName, user.LastName)
		assert.NoError(t, err)
	})
}

func TestRepo_Save(t *testing.T) {}

func TestRepo_Migrate(t *testing.T) {}
