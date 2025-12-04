package invoice

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InvoiceRepository struct {
	db *pgxpool.Pool
}

func GetInvoiceRepository(db *pgxpool.Pool) *InvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

// Get all invoices
func (r *InvoiceRepository) GetInvoices() ([]Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
        SELECT *
        FROM invoice
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices, err := pgx.CollectRows(rows, pgx.RowToStructByName[Invoice])
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// Get invoice by ID
func (r *InvoiceRepository) GetInvoiceById(id string) (*Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
        SELECT *
        FROM invoice
        WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inv, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Invoice])
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

// Create invoice
func (r *InvoiceRepository) CreateInvoice(inv Invoice) (*Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
        INSERT INTO invoice (
            organization_id,
            payment_id,
            invoice_number,
            customer_id,
            issue_date,
            due_date,
            amount,
            txt_amount
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
        RETURNING *
    `,
		inv.OrganizationId,
		inv.PaymentId,
		inv.InvoiceNumber,
		inv.CustomerId,
		inv.IssueDate,
		inv.DueDate,
		inv.Amount,
		inv.TxtAmount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	created, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Invoice])
	if err != nil {
		return nil, err
	}

	return &created, nil
}

// Update invoice
func (r *InvoiceRepository) UpdateInvoice(inv *Invoice) (*Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
        UPDATE invoice
        SET 
            organization_id = $1,
            payment_id = $2,
            invoice_number = $3,
            customer_id = $4,
            issue_date = $5,
            due_date = $6,
            amount = $7,
            txt_amount = $8
        WHERE id = $9
        RETURNING *
    `,
		inv.OrganizationId,
		inv.PaymentId,
		inv.InvoiceNumber,
		inv.CustomerId,
		inv.IssueDate,
		inv.DueDate,
		inv.Amount,
		inv.TxtAmount,
		inv.Id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Invoice])
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete invoice
func (r *InvoiceRepository) DeleteInvoice(id string) error {
	_, err := r.db.Exec(context.Background(), `
        DELETE FROM invoice
        WHERE id = $1
    `, id)

	return err
}
