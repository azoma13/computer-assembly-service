package models

import "errors"

// cancel order
var (
	ErrCancelOrderNotFound          = errors.New("order not found when canceling the order")
	ErrCancelOrderAlreadyPaid       = errors.New("order already paid, cannot be cancelled")
	ErrCancelOrderAlreadyInProgress = errors.New("order already in progress, cannot be cancelled")
	ErrCancelOrderAlreadyCompleted  = errors.New("order already completed, cannot be cancelled")
	ErrCancelOrderAlreadyCancelled  = errors.New("order already cancelled, cannot be cancelled again")
	ErrCancelOrderInternalService   = errors.New("internal error when canceled order")
)

// create order
var (
	ErrCreateOrderBadRequest = errors.New("bad request when creating the order")
	ErrCreateOrderNotFound   = errors.New("hardwares not found")
	ErrCreateOrderBadGateway = errors.New("bad gateway hardware")
)

// get order
var (
	ErrGetOrderNotFound      = errors.New("order not found when getting order")
	ErrGetOrderInternalError = errors.New("internal error when getting order")
)

// payment order
var (
	ErrPaymentOrderNotFound      = errors.New("payment not found")
	ErrPaymentOrderConflict      = errors.New("payment conflict")
	ErrPaymentOrderInternalError = errors.New("internal error while processing payment")
)
