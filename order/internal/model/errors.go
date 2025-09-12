package model

import "errors"

var (
	ErrOrderNotFound          = errors.New("order not found")
	ErrOrderConflict          = errors.New("order conflict")
	ErrOrderUndefinedStatus   = errors.New("order status undefined")
	ErrOrderAlreadyPaid       = errors.New("order already paid")
	ErrOrderAlreadyInProgress = errors.New("order already in progress")
	ErrOrderAlreadyCompleted  = errors.New("order already completed")
	ErrOrderAlreadyCancelled  = errors.New("order already cancelled")
)
