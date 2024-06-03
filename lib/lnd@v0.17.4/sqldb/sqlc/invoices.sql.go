// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: invoices.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const deleteInvoice = `-- name: DeleteInvoice :exec
DELETE 
FROM invoices 
WHERE (
    id = $1 OR 
    $1 IS NULL
) AND (
    hash = $2 OR 
    $2 IS NULL
) AND (
    preimage = $3 OR 
    $3 IS NULL
) AND (
    payment_addr = $4 OR
    $4 IS NULL
)
`

type DeleteInvoiceParams struct {
	AddIndex    sql.NullInt32
	Hash        []byte
	Preimage    []byte
	PaymentAddr []byte
}

func (q *Queries) DeleteInvoice(ctx context.Context, arg DeleteInvoiceParams) error {
	_, err := q.db.ExecContext(ctx, deleteInvoice,
		arg.AddIndex,
		arg.Hash,
		arg.Preimage,
		arg.PaymentAddr,
	)
	return err
}

const deleteInvoiceFeatures = `-- name: DeleteInvoiceFeatures :exec
DELETE 
FROM invoice_features
WHERE invoice_id = $1
`

func (q *Queries) DeleteInvoiceFeatures(ctx context.Context, invoiceID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceFeatures, invoiceID)
	return err
}

const deleteInvoiceHTLC = `-- name: DeleteInvoiceHTLC :exec
DELETE
FROM invoice_htlcs
WHERE htlc_id = $1
`

func (q *Queries) DeleteInvoiceHTLC(ctx context.Context, htlcID int64) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceHTLC, htlcID)
	return err
}

const deleteInvoiceHTLCCustomRecords = `-- name: DeleteInvoiceHTLCCustomRecords :exec
WITH htlc_ids AS (
    SELECT ih.id
    FROM invoice_htlcs ih JOIN invoice_htlc_custom_records ihcr ON ih.id=ihcr.htlc_id 
    WHERE ih.invoice_id = $1
)
DELETE
FROM invoice_htlc_custom_records
WHERE htlc_id IN (SELECT id FROM htlc_ids)
`

func (q *Queries) DeleteInvoiceHTLCCustomRecords(ctx context.Context, invoiceID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceHTLCCustomRecords, invoiceID)
	return err
}

const deleteInvoiceHTLCs = `-- name: DeleteInvoiceHTLCs :exec
DELETE
FROM invoice_htlcs
WHERE invoice_id = $1
`

func (q *Queries) DeleteInvoiceHTLCs(ctx context.Context, invoiceID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceHTLCs, invoiceID)
	return err
}

const filterInvoicePayments = `-- name: FilterInvoicePayments :many
SELECT  
    ip.id AS settle_index, ip.amount_paid_msat, ip.settled_at AS settle_date, 
    i.id, i.hash, i.preimage, i.memo, i.amount_msat, i.cltv_delta, i.expiry, i.payment_addr, i.payment_request, i.state, i.amount_paid_msat, i.is_amp, i.is_hodl, i.is_keysend, i.created_at 
FROM invoice_payments ip JOIN invoices i ON ip.invoice_id = i.id
WHERE (
    ip.id >= $1 OR 
    $1 IS NULL
) AND (
    ip.settled_at >= $2 OR 
    $2 IS NULL
)
ORDER BY
    CASE
        WHEN $3 = FALSE THEN ip.id  
        ELSE NULL
    END ASC,
    CASE
        WHEN $3 = TRUE THEN ip.id  
        ELSE NULL
    END DESC
LIMIT $5 OFFSET $4
`

type FilterInvoicePaymentsParams struct {
	SettleIndexGet sql.NullInt32
	SettledAfter   sql.NullTime
	Reverse        interface{}
	NumOffset      int32
	NumLimit       int32
}

type FilterInvoicePaymentsRow struct {
	SettleIndex      int32
	AmountPaidMsat   int64
	SettleDate       time.Time
	ID               int32
	Hash             []byte
	Preimage         []byte
	Memo             sql.NullString
	AmountMsat       int64
	CltvDelta        sql.NullInt32
	Expiry           int32
	PaymentAddr      []byte
	PaymentRequest   sql.NullString
	State            int16
	AmountPaidMsat_2 int64
	IsAmp            bool
	IsHodl           bool
	IsKeysend        bool
	CreatedAt        time.Time
}

func (q *Queries) FilterInvoicePayments(ctx context.Context, arg FilterInvoicePaymentsParams) ([]FilterInvoicePaymentsRow, error) {
	rows, err := q.db.QueryContext(ctx, filterInvoicePayments,
		arg.SettleIndexGet,
		arg.SettledAfter,
		arg.Reverse,
		arg.NumOffset,
		arg.NumLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterInvoicePaymentsRow
	for rows.Next() {
		var i FilterInvoicePaymentsRow
		if err := rows.Scan(
			&i.SettleIndex,
			&i.AmountPaidMsat,
			&i.SettleDate,
			&i.ID,
			&i.Hash,
			&i.Preimage,
			&i.Memo,
			&i.AmountMsat,
			&i.CltvDelta,
			&i.Expiry,
			&i.PaymentAddr,
			&i.PaymentRequest,
			&i.State,
			&i.AmountPaidMsat_2,
			&i.IsAmp,
			&i.IsHodl,
			&i.IsKeysend,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const filterInvoices = `-- name: FilterInvoices :many
SELECT id, hash, preimage, memo, amount_msat, cltv_delta, expiry, payment_addr, payment_request, state, amount_paid_msat, is_amp, is_hodl, is_keysend, created_at 
FROM invoices
WHERE (
    id >= $1 OR 
    $1 IS NULL
) AND (
    id <= $2 OR 
    $2 IS NULL
) AND (
    state = $3 OR 
    $3 IS NULL
) AND (
    created_at >= $4 OR
    $4 IS NULL
) AND (
    created_at <= $5 OR 
    $5 IS NULL
) AND (
    CASE
        WHEN $6=TRUE THEN (state = 0 OR state = 3)
        ELSE TRUE 
    END
)
ORDER BY
    CASE
        WHEN $7 = FALSE THEN id  
        ELSE NULL
    END ASC,
    CASE
        WHEN $7 = TRUE  THEN id  
        ELSE NULL
    END DESC
LIMIT $9 OFFSET $8
`

type FilterInvoicesParams struct {
	AddIndexGet   sql.NullInt32
	AddIndexLet   sql.NullInt32
	State         sql.NullInt16
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	PendingOnly   interface{}
	Reverse       interface{}
	NumOffset     int32
	NumLimit      int32
}

func (q *Queries) FilterInvoices(ctx context.Context, arg FilterInvoicesParams) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, filterInvoices,
		arg.AddIndexGet,
		arg.AddIndexLet,
		arg.State,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.PendingOnly,
		arg.Reverse,
		arg.NumOffset,
		arg.NumLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.Hash,
			&i.Preimage,
			&i.Memo,
			&i.AmountMsat,
			&i.CltvDelta,
			&i.Expiry,
			&i.PaymentAddr,
			&i.PaymentRequest,
			&i.State,
			&i.AmountPaidMsat,
			&i.IsAmp,
			&i.IsHodl,
			&i.IsKeysend,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoice = `-- name: GetInvoice :many

SELECT id, hash, preimage, memo, amount_msat, cltv_delta, expiry, payment_addr, payment_request, state, amount_paid_msat, is_amp, is_hodl, is_keysend, created_at
FROM invoices
WHERE (
    id = $1 OR 
    $1 IS NULL
) AND (
    hash = $2 OR 
    $2 IS NULL
) AND (
    preimage = $3 OR 
    $3 IS NULL
) AND (
    payment_addr = $4 OR 
    $4 IS NULL
)
LIMIT 2
`

type GetInvoiceParams struct {
	AddIndex    sql.NullInt32
	Hash        []byte
	Preimage    []byte
	PaymentAddr []byte
}

// This method may return more than one invoice if filter using multiple fields
// from different invoices. It is the caller's responsibility to ensure that
// we bubble up an error in those cases.
func (q *Queries) GetInvoice(ctx context.Context, arg GetInvoiceParams) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getInvoice,
		arg.AddIndex,
		arg.Hash,
		arg.Preimage,
		arg.PaymentAddr,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.Hash,
			&i.Preimage,
			&i.Memo,
			&i.AmountMsat,
			&i.CltvDelta,
			&i.Expiry,
			&i.PaymentAddr,
			&i.PaymentRequest,
			&i.State,
			&i.AmountPaidMsat,
			&i.IsAmp,
			&i.IsHodl,
			&i.IsKeysend,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoiceFeatures = `-- name: GetInvoiceFeatures :many
SELECT feature, invoice_id
FROM invoice_features
WHERE invoice_id = $1
`

func (q *Queries) GetInvoiceFeatures(ctx context.Context, invoiceID int32) ([]InvoiceFeature, error) {
	rows, err := q.db.QueryContext(ctx, getInvoiceFeatures, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceFeature
	for rows.Next() {
		var i InvoiceFeature
		if err := rows.Scan(&i.Feature, &i.InvoiceID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoiceHTLCCustomRecords = `-- name: GetInvoiceHTLCCustomRecords :many
SELECT ihcr.htlc_id, key, value
FROM invoice_htlcs ih JOIN invoice_htlc_custom_records ihcr ON ih.id=ihcr.htlc_id 
WHERE ih.invoice_id = $1
`

type GetInvoiceHTLCCustomRecordsRow struct {
	HtlcID int64
	Key    int64
	Value  []byte
}

func (q *Queries) GetInvoiceHTLCCustomRecords(ctx context.Context, invoiceID int32) ([]GetInvoiceHTLCCustomRecordsRow, error) {
	rows, err := q.db.QueryContext(ctx, getInvoiceHTLCCustomRecords, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInvoiceHTLCCustomRecordsRow
	for rows.Next() {
		var i GetInvoiceHTLCCustomRecordsRow
		if err := rows.Scan(&i.HtlcID, &i.Key, &i.Value); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoiceHTLCs = `-- name: GetInvoiceHTLCs :many
SELECT id, htlc_id, chan_id, amount_msat, total_mpp_msat, accept_height, accept_time, expiry_height, state, resolve_time, invoice_id
FROM invoice_htlcs
WHERE invoice_id = $1
`

func (q *Queries) GetInvoiceHTLCs(ctx context.Context, invoiceID int32) ([]InvoiceHtlc, error) {
	rows, err := q.db.QueryContext(ctx, getInvoiceHTLCs, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceHtlc
	for rows.Next() {
		var i InvoiceHtlc
		if err := rows.Scan(
			&i.ID,
			&i.HtlcID,
			&i.ChanID,
			&i.AmountMsat,
			&i.TotalMppMsat,
			&i.AcceptHeight,
			&i.AcceptTime,
			&i.ExpiryHeight,
			&i.State,
			&i.ResolveTime,
			&i.InvoiceID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoicePayments = `-- name: GetInvoicePayments :many
SELECT id, settled_at, amount_paid_msat, invoice_id
FROM invoice_payments
WHERE invoice_id = $1
`

func (q *Queries) GetInvoicePayments(ctx context.Context, invoiceID int32) ([]InvoicePayment, error) {
	rows, err := q.db.QueryContext(ctx, getInvoicePayments, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoicePayment
	for rows.Next() {
		var i InvoicePayment
		if err := rows.Scan(
			&i.ID,
			&i.SettledAt,
			&i.AmountPaidMsat,
			&i.InvoiceID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertInvoice = `-- name: InsertInvoice :one
INSERT INTO invoices (
    hash, preimage, memo, amount_msat, cltv_delta, expiry, payment_addr, 
    payment_request, state, amount_paid_msat, is_amp, is_hodl, is_keysend, 
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING id
`

type InsertInvoiceParams struct {
	Hash           []byte
	Preimage       []byte
	Memo           sql.NullString
	AmountMsat     int64
	CltvDelta      sql.NullInt32
	Expiry         int32
	PaymentAddr    []byte
	PaymentRequest sql.NullString
	State          int16
	AmountPaidMsat int64
	IsAmp          bool
	IsHodl         bool
	IsKeysend      bool
	CreatedAt      time.Time
}

func (q *Queries) InsertInvoice(ctx context.Context, arg InsertInvoiceParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertInvoice,
		arg.Hash,
		arg.Preimage,
		arg.Memo,
		arg.AmountMsat,
		arg.CltvDelta,
		arg.Expiry,
		arg.PaymentAddr,
		arg.PaymentRequest,
		arg.State,
		arg.AmountPaidMsat,
		arg.IsAmp,
		arg.IsHodl,
		arg.IsKeysend,
		arg.CreatedAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertInvoiceFeature = `-- name: InsertInvoiceFeature :exec
INSERT INTO invoice_features (
    invoice_id, feature
) VALUES (
    $1, $2
)
`

type InsertInvoiceFeatureParams struct {
	InvoiceID int32
	Feature   int32
}

func (q *Queries) InsertInvoiceFeature(ctx context.Context, arg InsertInvoiceFeatureParams) error {
	_, err := q.db.ExecContext(ctx, insertInvoiceFeature, arg.InvoiceID, arg.Feature)
	return err
}

const insertInvoiceHTLC = `-- name: InsertInvoiceHTLC :exec
INSERT INTO invoice_htlcs (
    htlc_id, chan_id, amount_msat, total_mpp_msat, accept_height, accept_time,
    expiry_height, state, resolve_time, invoice_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
`

type InsertInvoiceHTLCParams struct {
	HtlcID       int64
	ChanID       string
	AmountMsat   int64
	TotalMppMsat sql.NullInt64
	AcceptHeight int32
	AcceptTime   time.Time
	ExpiryHeight int32
	State        int16
	ResolveTime  sql.NullTime
	InvoiceID    int32
}

func (q *Queries) InsertInvoiceHTLC(ctx context.Context, arg InsertInvoiceHTLCParams) error {
	_, err := q.db.ExecContext(ctx, insertInvoiceHTLC,
		arg.HtlcID,
		arg.ChanID,
		arg.AmountMsat,
		arg.TotalMppMsat,
		arg.AcceptHeight,
		arg.AcceptTime,
		arg.ExpiryHeight,
		arg.State,
		arg.ResolveTime,
		arg.InvoiceID,
	)
	return err
}

const insertInvoiceHTLCCustomRecord = `-- name: InsertInvoiceHTLCCustomRecord :exec
INSERT INTO invoice_htlc_custom_records (
    key, value, htlc_id 
) VALUES (
    $1, $2, $3
)
`

type InsertInvoiceHTLCCustomRecordParams struct {
	Key    int64
	Value  []byte
	HtlcID int64
}

func (q *Queries) InsertInvoiceHTLCCustomRecord(ctx context.Context, arg InsertInvoiceHTLCCustomRecordParams) error {
	_, err := q.db.ExecContext(ctx, insertInvoiceHTLCCustomRecord, arg.Key, arg.Value, arg.HtlcID)
	return err
}

const insertInvoicePayment = `-- name: InsertInvoicePayment :one
INSERT INTO invoice_payments (
    invoice_id, amount_paid_msat, settled_at 
) VALUES (
    $1, $2, $3
) RETURNING id
`

type InsertInvoicePaymentParams struct {
	InvoiceID      int32
	AmountPaidMsat int64
	SettledAt      time.Time
}

func (q *Queries) InsertInvoicePayment(ctx context.Context, arg InsertInvoicePaymentParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertInvoicePayment, arg.InvoiceID, arg.AmountPaidMsat, arg.SettledAt)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateInvoice = `-- name: UpdateInvoice :exec
UPDATE invoices 
SET preimage=$2, state=$3, amount_paid_msat=$4
WHERE id=$1
`

type UpdateInvoiceParams struct {
	ID             int32
	Preimage       []byte
	State          int16
	AmountPaidMsat int64
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) error {
	_, err := q.db.ExecContext(ctx, updateInvoice,
		arg.ID,
		arg.Preimage,
		arg.State,
		arg.AmountPaidMsat,
	)
	return err
}

const updateInvoiceHTLC = `-- name: UpdateInvoiceHTLC :exec
UPDATE invoice_htlcs 
SET state=$2, resolve_time=$3
WHERE id = $1
`

type UpdateInvoiceHTLCParams struct {
	ID          int32
	State       int16
	ResolveTime sql.NullTime
}

func (q *Queries) UpdateInvoiceHTLC(ctx context.Context, arg UpdateInvoiceHTLCParams) error {
	_, err := q.db.ExecContext(ctx, updateInvoiceHTLC, arg.ID, arg.State, arg.ResolveTime)
	return err
}

const updateInvoiceHTLCs = `-- name: UpdateInvoiceHTLCs :exec
UPDATE invoice_htlcs 
SET state=$2, resolve_time=$3
WHERE invoice_id = $1 AND resolve_time IS NULL
`

type UpdateInvoiceHTLCsParams struct {
	InvoiceID   int32
	State       int16
	ResolveTime sql.NullTime
}

func (q *Queries) UpdateInvoiceHTLCs(ctx context.Context, arg UpdateInvoiceHTLCsParams) error {
	_, err := q.db.ExecContext(ctx, updateInvoiceHTLCs, arg.InvoiceID, arg.State, arg.ResolveTime)
	return err
}
