package v1

import (
	"context"
	"net/http"

	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
)

func (a *orderAPI) NewError(ctx context.Context, err error) *order_v1.UnexpectedErrorStatusCode {
	return &order_v1.UnexpectedErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: order_v1.UnexpectedError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}}
}
