package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	authRepository "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/repository"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "SELECT id, uuid, first_name, last_name, username, password, email, updated_at, created_at FROM user WHERE ID = \\?"
	rows := sqlmock.NewRows([]string{"id", "uuid", "first_name", "last_name", "username", "password", "email", "updated_at", "created_at"}).AddRow(1, "1820208c-f225-4777-96a5-95e24ecf7b25", "First Name", "Last Name", "username", "password", "email", time.Now(), time.Now())

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := authRepository.NewPostgreSQLUserRepository(db)

	id := 1
	user, err := a.FindByID(context.TODO(), int64(id))
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUpdate(t *testing.T) {
	user := &model.User{
		UUID:      uuid.NewV1(),
		FirstName: "First Name",
		LastName:  "Last Name",
		Username:  "username",
		Email:     "email@email.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE user SET first_name=\\?, last_name=\\?, username=\\?, password=\\?, email=\\?, updated_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := authRepository.NewPostgreSQLUserRepository(db)

	err = a.Update(context.TODO(), user)
	assert.NoError(t, err)
}

func TestSave(t *testing.T) {
	user := &model.User{
		UUID:      uuid.NewV1(),
		FirstName: "First Name",
		LastName:  "Last Name",
		Username:  "username",
		Email:     "email@email.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT user SET uuid=\\?, first_name=\\?, last_name=\\?, username=\\?, email=\\?, password=\\?, created_at=\\?, updated_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.UUID, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := authRepository.NewPostgreSQLUserRepository(db)
	err = a.Save(context.TODO(), user)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), user.ID)
	assert.Equal(t, "username", user.Username)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM user WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 12))

	a := authRepository.NewPostgreSQLUserRepository(db)

	num := int64(12)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}
