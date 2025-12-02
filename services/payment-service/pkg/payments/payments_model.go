package payments

import "time"

// Payment represents a payment record
type Payment struct {
	Id        string    `json:"id" db:"Id"`
	UserId    int32     `json:"user_id" db:"UserId"`
	CreatedAt time.Time `json:"created_at" db:"CreatedAt"`
}
