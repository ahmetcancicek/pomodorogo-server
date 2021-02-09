package postgresql_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/repository/postgresql"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository auth.Repository
	user       *model.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = postgresql.CreateRepository(s.DB)

}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestFindByID() {
	var (
		id        int64  = 123456
		firstName string = "FirstName"
	)

	mockedRow := sqlmock.NewRows([]string{"id", "first_name"}).AddRow(id, firstName)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByID(id)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), firstName, res.FirstName, "They should be equal!")

	assert.Equal(s.T(), id, res.ID, "They should be equal!")
}

func (s *Suite) TestUpdate() {
	var (
		id        int64     = 123456
		uuid                = uuid.NewV4()
		firstName string    = "FirstName"
		lastName  string    = "LastName"
		username  string    = "username"
		email     string    = "email@email.com"
		password  string    = "password"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`UPDATE "users" SET "first_name" = $1,"last_name" = $2,"username" = $3,"email" = $4,"password" = $5,"updated_at" = $6 WHERE "id" = $7 RETURNING "id"`)).
		WithArgs(firstName, lastName, username, email, password, updatedAt, id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	// New value
	firstName = "Update FirstName"

	user := &model.User{
		ID:        id,
		UUID:      uuid,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err := s.repository.Update(user)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), firstName, user.FirstName, "They should be equal!")

}

func (s *Suite) TestSave() {
	var (
		id        int64     = 123456
		uuid                = uuid.NewV4()
		firstName string    = "FirstName"
		lastName  string    = "LastName"
		username  string    = "username"
		email     string    = "email@email.com"
		password  string    = "password"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("uuid","first_name","last_name","username","email","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`)).
		WithArgs(uuid, firstName, lastName, username, email, password, createdAt, updatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	user := &model.User{
		UUID:      uuid,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err := s.repository.Save(user)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), id, user.ID, "They should be equal!")

}

func (s *Suite) TestDelete() {
	var (
		id int64 = 123456
	)

	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(id)
	require.NoError(s.T(), err)

}
