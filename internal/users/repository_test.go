package users

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("Create", func(t *testing.T) {
		user := &User{
			ID:           uuid.New(),
			Username:     "testuser",
			Email:        "test@example.com",
			PasswordHash: "hashedpassword",
			CreatedAt:    time.Now(),
		}

		mock.ExpectExec("INSERT INTO users").
			WithArgs(user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(user)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("FindAll", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
			AddRow(uuid.New(), "user1", "user1@example.com", "hash1", time.Now()).
			AddRow(uuid.New(), "user2", "user2@example.com", "hash2", time.Now())

		mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users").
			WillReturnRows(rows)

		users, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("FindByID", func(t *testing.T) {
		id := uuid.New()
		mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE id = \\$1").
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
				AddRow(id, "user1", "user1@example.com", "hash1", time.Now()))

		user, err := repo.FindByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, id, user.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("FindByUsername", func(t *testing.T) {
		username := "user1"
		mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE username = \\$1").
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
				AddRow(uuid.New(), username, "user1@example.com", "hash1", time.Now()))

		user, err := repo.FindByUsername(username)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, username, user.Username)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("FindByEmail", func(t *testing.T) {
		email := "user1@example.com"
		mock.ExpectQuery("SELECT id, username, email, password_hash, created_at FROM users WHERE email = \\$1").
			WithArgs(email).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password_hash", "created_at"}).
				AddRow(uuid.New(), "user1", email, "hash1", time.Now()))

		user, err := repo.FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update", func(t *testing.T) {
		user := &User{
			ID:           uuid.New(),
			Username:     "updateduser",
			Email:        "updated@example.com",
			PasswordHash: "updatedhash",
		}

		mock.ExpectExec("UPDATE users SET username = \\$1, email = \\$2, password_hash = \\$3 WHERE id = \\$4").
			WithArgs(user.Username, user.Email, user.PasswordHash, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Update(user)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
