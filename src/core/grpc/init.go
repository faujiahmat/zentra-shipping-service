package grpc

import (
	"github.com/faujiahmat/zentra-shipping-service/src/core/grpc/client"
	"github.com/faujiahmat/zentra-shipping-service/src/core/grpc/delivery"
	"github.com/faujiahmat/zentra-shipping-service/src/core/grpc/interceptor"
)

func InitClient() *client.Grpc {
	unaryRequestInterceptor := interceptor.NewUnaryRequest()

	orderDelivery, orderConn := delivery.NewOrderGrpc(unaryRequestInterceptor)

	grpcClient := client.NewGrpc(orderDelivery, orderConn)
	return grpcClient
}
