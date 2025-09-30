package delivery

import (
	"context"

	pb "github.com/faujiahmat/zentra-proto/protogen/order"
	"github.com/faujiahmat/zentra-shipping-service/src/common/log"
	"github.com/faujiahmat/zentra-shipping-service/src/core/grpc/interceptor"
	"github.com/faujiahmat/zentra-shipping-service/src/infrastructure/cbreaker"
	"github.com/faujiahmat/zentra-shipping-service/src/infrastructure/config"
	"github.com/faujiahmat/zentra-shipping-service/src/interface/delivery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderGrpcImpl struct {
	client pb.OrderServiceClient
}

func NewOrderGrpc(unaryRequest *interceptor.UnaryRequest) (delivery.OrderGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryRequest.AddBasicAuth),
	)

	conn, err := grpc.NewClient(config.Conf.ApiGateway.BaseUrl, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.NewOrderGrpc", "section": "grpc.NewClient"}).Fatal(err)
	}

	client := pb.NewOrderServiceClient(conn)

	return &OrderGrpcImpl{
		client: client,
	}, conn
}

func (o *OrderGrpcImpl) AddShippingId(ctx context.Context, data *pb.AddShippingIdReq) error {
	_, err := cbreaker.OrderGrpc.Execute(func() (any, error) {
		_, err := o.client.AddShippingId(ctx, data)
		return nil, err
	})

	return err
}

func (o *OrderGrpcImpl) UpdateStatus(ctx context.Context, data *pb.UpdateStatusReq) error {
	_, err := cbreaker.OrderGrpc.Execute(func() (any, error) {
		_, err := o.client.UpdateStatus(ctx, data)
		return nil, err
	})

	return err
}
