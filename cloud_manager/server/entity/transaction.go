package entity

type Transaction struct {
	Id         string            `json:"id"`
	ExternalId *string           `json:"external_id,omitempty"`
	ExpireAt   int64             `json:"expire_at"`
	Status     TransactionStatus `json:"status"`
}
