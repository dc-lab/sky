package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type TransactionStatus int

const (
	TxStatusUnknown TransactionStatus = iota
	TxStatusActive
	TxStatusCompleted
	TxStatusExpired
	TxStatusCancelled
)

func (ts TransactionStatus) String() string {
	return txStatusToString[ts]
}

var txStatusToString = map[TransactionStatus]string{
	TxStatusUnknown:   "unknown",
	TxStatusActive:    "active",
	TxStatusCompleted: "completed",
	TxStatusExpired:   "expired",
	TxStatusCancelled: "cancelled",
}

var txStatusToID = map[string]TransactionStatus{
	"unknown":   TxStatusUnknown,
	"active":    TxStatusActive,
	"completed": TxStatusCompleted,
	"expired":   TxStatusExpired,
	"cancelled": TxStatusCancelled,
}

func (ts TransactionStatus) MarshalJSON() ([]byte, error) {
	if ts == TxStatusUnknown {
		return nil, fmt.Errorf("cannot serialize unknown transaction status")
	}
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(txStatusToString[ts])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (ts *TransactionStatus) UnmarshalJSON(b []byte) error {
	var idx string
	err := json.Unmarshal(b, &idx)
	if err != nil {
		return err
	}
	// NOTE: return Unknown for unrecognized id
	*ts = txStatusToID[idx]
	return nil
}
