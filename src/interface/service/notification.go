package service

import (
	"context"

	"github.com/faujiahmat/zentra-shipping-service/src/model/entity"
)

type Notification interface {
	Shipper(ctx context.Context, data *entity.Shipper) error
}
