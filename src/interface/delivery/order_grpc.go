package delivery

import (
	"context"

	pb "github.com/faujiahmat/zentra-proto/protogen/order"
)

type OrderGrpc interface {
	AddShippingId(ctx context.Context, data *pb.AddShippingIdReq) error
	UpdateStatus(ctx context.Context, data *pb.UpdateStatusReq) error
}
