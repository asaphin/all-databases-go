package postgres

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresSQLXAddressRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXAddressRepository() (*PostgresSQLXAddressRepository, error) {
	db, err := NewSqlx(sqlxDatabaseName)
	if err != nil {
		return nil, err
	}

	return &PostgresSQLXAddressRepository{
		db: db,
	}, nil
}

func (repo *PostgresSQLXAddressRepository) Create(ctx context.Context, address domain.Address) (string, error) {
	query := `INSERT INTO addresses (id, street, city, state, zip) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := repo.db.QueryRowContext(ctx, query, address.ID, address.Street, address.City, address.State, address.Zip).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (repo *PostgresSQLXAddressRepository) List(ctx context.Context, limit, offset int) ([]*domain.Address, error) {
	query := `SELECT * FROM addresses LIMIT $1 OFFSET $2`
	var addresses []*domain.Address
	err := repo.db.SelectContext(ctx, &addresses, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (repo *PostgresSQLXAddressRepository) GetByID(ctx context.Context, addressID string) (*domain.Address, error) {
	query := `SELECT * FROM addresses WHERE id = $1`
	var address domain.Address
	err := repo.db.GetContext(ctx, &address, query, addressID)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repo *PostgresSQLXAddressRepository) Update(ctx context.Context, address domain.Address) error {
	query := `UPDATE addresses SET street=$1, city=$2, state=$3, zip=$4 WHERE id=$5`
	_, err := repo.db.ExecContext(ctx, query, address.Street, address.City, address.State, address.Zip, address.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresSQLXAddressRepository) Delete(ctx context.Context, addressID string) error {
	query := `DELETE FROM addresses WHERE id=$1`
	_, err := repo.db.ExecContext(ctx, query, addressID)
	if err != nil {
		return err
	}
	return nil
}