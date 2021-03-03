package postgresql_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/repository/postgresql"
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

	repository tag.Repository
	tag        *model.Tag
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
	mockID        int64 = 12345
	mockUserId    int64 = 12345
	mockName            = "Name"
	mockColour          = "#000000"
	mockCreatedAt       = time.Now()
	mockUpdatedAt       = time.Now()
	mockDeletedAt       = time.Now()
)

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestFindByID() {
	mockedRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(mockID, mockName)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tags" WHERE id = $1 ORDER BY "tags"."id" LIMIT 1`)).
		WithArgs(mockID).
		WillReturnRows(mockedRow)

	res, err := s.repository.FindByID(mockID)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockName, res.Name, "They should be equal!")

	assert.Equal(s.T(), mockID, res.ID, "They should be equal!")
}

func (s *Suite) TestSave() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "tags" ("user_id","name","colour","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(mockUserId, mockName, mockColour, mockCreatedAt, mockUpdatedAt, mockDeletedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockID))

	tag := &model.Tag{
		UserID:    mockUserId,
		Name:      mockName,
		Colour:    mockColour,
		CreatedAt: mockCreatedAt,
		DeletedAt: &mockDeletedAt,
		UpdatedAt: mockUpdatedAt,
	}

	tag, err := s.repository.Save(tag)

	require.NoError(s.T(), err)

	assert.Equal(s.T(), mockID, tag.ID, "They should be equal!")
}

func (s *Suite) TestDelete() {
	var (
		id int64 = 123456
	)

	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tags" WHERE "tags"."id" = $1`)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(id)
	require.NoError(s.T(), err)
}

func (s *Suite) TestUpdate() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`UPDATE "tags" SET "name" = $1,"colour" = $2,"updated_at" = $3,"deleted_at" = $4 WHERE "id" = $5 RETURNING "id"`)).
		WithArgs(mockName, mockColour, mockUpdatedAt, mockDeletedAt, mockID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockID))

	// New value
	mockName = "Update Name"

	tag := &model.Tag{
		ID:        mockID,
		UserID:    mockUserId,
		Name:      mockName,
		Colour:    mockColour,
		CreatedAt: mockCreatedAt,
		DeletedAt: &mockDeletedAt,
		UpdatedAt: mockUpdatedAt,
	}

	tag, err := s.repository.Update(tag)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), mockName, tag.Name, "They should be equal!")

}
