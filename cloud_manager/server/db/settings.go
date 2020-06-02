package db

import "time"

const (
	defaultOfflineTimeout = 5 * time.Second
	defaultRuntimeWriteTimeout = 1000 * time.Millisecond
	defaultRuntimeReadTimeout  = 500 * time.Millisecond
)
