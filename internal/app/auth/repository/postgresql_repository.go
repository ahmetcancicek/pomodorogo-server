package repository

import (
	"context"
	"database/sql"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

type postgreSQLUserRepository struct {
	Conn *sql.DB
}

func NewPostgreSQLUserRepository(Conn *sql.DB) auth.Repository {
	return &postgreSQLUserRepository{Conn}
}

func (p postgreSQLUserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {

	query := `SELECT id, uuid, first_name, last_name, username, password, email, updated_at, created_at FROM user WHERE ID = ?`

	row, err := p.Conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	user := new(model.User)
	row.Scan(
		&user.ID,
		&user.UUID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)

	return user, nil
}

func (p postgreSQLUserRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE user set first_name=?, last_name=?, username=?, password=?, email=?, updated_at=? WHERE ID = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return nil
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return nil
	}

	if rowAffected != 1 {
		return err
	}

	return nil
}

func (p postgreSQLUserRepository) Save(ctx context.Context, user *model.User) error {
	query := `INSERT user SET uuid=?, first_name=?, last_name=?, username=?, email=?, password=?, created_at=?, updated_at=?`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(
		ctx,
		user.UUID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt)

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	user.ID = lastID
	return nil
}

func (p postgreSQLUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM user WHERE id = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return err
	}

	return nil
}
