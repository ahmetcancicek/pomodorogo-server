package postgresql_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	user2 "github.com/ahmetcancicek/pomodorogo-server/internal/app/user"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/user/repository/postgresql"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/util"
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

	repository user2.Repository
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

var (
	mockID        int64 = 123456
	mockUUID            = uuid.NewV4()
	mockFirstName       = "FirstName"
	mockLastName        = "LastName"
	mockUsername        = "username"
	mockEmail           = "email@email.com"
	mockTokenHash       = util.GenerateRandomString(15)
	mockPassword        = "password"
	mockCreatedAt       = time.Now()
	mockUpdatedAt       = time.Now()
)

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestFindByID() {
	mockedRow := sqlmock.NewRows([]string{"id", "first_name"}).AddRow(mockID, mockFirstName)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(mockID).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByID(mockID)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockFirstName, res.FirstName, "They should be equal!")

	assert.Equal(s.T(), mockID, res.ID, "They should be equal!")
}

func (s *Suite) TestFindByUUID() {
	mockedRow := sqlmock.NewRows([]string{"id", "uuid", "first_name"}).AddRow(mockID, mockUUID, mockFirstName)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE uuid = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(mockUUID).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByUUID(mockUUID.String())

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockFirstName, res.FirstName, "They should be equal!")

	assert.Equal(s.T(), mockID, res.ID, "They should be equal!")

}

func (s *Suite) TestFindByEmail() {
	mockedRow := sqlmock.NewRows([]string{"id", "first_name", "email"}).AddRow(mockID, mockFirstName, mockEmail)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(mockEmail).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByEmail(mockEmail)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockFirstName, res.FirstName, "They should be equal!")

	assert.Equal(s.T(), mockEmail, res.Email, "They should be equal!")
}

func (s *Suite) TestFindByCredentials() {

	mockedRow := sqlmock.NewRows([]string{"id", "email", "username", "password"}).AddRow(mockID, mockEmail, mockUsername, mockPassword)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(mockEmail).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByCredentials(mockEmail, mockPassword)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockPassword, res.Password, "They should be equal!")

	assert.Equal(s.T(), mockEmail, res.Email, "They should be equal!")

}

func (s *Suite) TestUpdate() {

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`UPDATE "users" SET "first_name" = $1,"last_name" = $2,"username" = $3,"email" = $4,"password" = $5,"updated_at" = $6 WHERE "id" = $7 RETURNING "id"`)).
		WithArgs(mockFirstName, mockLastName, mockUsername, mockEmail, mockPassword, mockUpdatedAt, mockID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockID))

	// New value
	mockFirstName = "Update FirstName"

	user := &model.User{
		ID:        mockID,
		UUID:      mockUUID,
		FirstName: mockFirstName,
		LastName:  mockLastName,
		Username:  mockUsername,
		Email:     mockEmail,
		Password:  mockPassword,
		CreatedAt: mockCreatedAt,
		UpdatedAt: mockUpdatedAt,
	}

	err := s.repository.Update(user)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), mockFirstName, user.FirstName, "They should be equal!")

}

func (s *Suite) TestSave() {

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("uuid","first_name","last_name","username","token_hash","email","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`)).
		WithArgs(mockUUID, mockFirstName, mockLastName, mockUsername, mockTokenHash, mockEmail, mockPassword, mockCreatedAt, mockUpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockID))

	user := &model.User{
		UUID:      mockUUID,
		FirstName: mockFirstName,
		LastName:  mockLastName,
		Username:  mockUsername,
		Email:     mockEmail,
		TokenHash: mockTokenHash,
		Password:  mockPassword,
		CreatedAt: mockCreatedAt,
		UpdatedAt: mockUpdatedAt,
	}

	err := s.repository.Save(user)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), mockID, user.ID, "They should be equal!")

}

func (s *Suite) TestDelete() {
	var (
		id int64 = 123456
	)

	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(id)
	require.NoError(s.T(), err)

}
