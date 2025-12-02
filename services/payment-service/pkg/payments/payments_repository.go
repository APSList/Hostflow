package payments

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentsRepository struct {
	db *pgxpool.Pool
}

func GetPaymentsRepository(db *pgxpool.Pool) *PaymentsRepository {
	return &PaymentsRepository{
		db: db,
	}
}

func (r *PaymentsRepository) GetPayments() ([]Payment, error) {
	rows, err := r.db.Query(context.Background(), `
        SELECT *
        FROM payment
        ORDER BY "CreatedAt" DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments, err := pgx.CollectRows(rows, pgx.RowToStructByName[Payment])
	if err != nil {
		return nil, err
	}

	return payments, nil
}
