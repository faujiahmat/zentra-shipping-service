package restful

import (
	"github.com/faujiahmat/zentra-shipping-service/src/core/restful/client"
	"github.com/faujiahmat/zentra-shipping-service/src/core/restful/delivery"
	"github.com/faujiahmat/zentra-shipping-service/src/core/restful/handler"
	"github.com/faujiahmat/zentra-shipping-service/src/core/restful/middleware"
	"github.com/faujiahmat/zentra-shipping-service/src/core/restful/server"
	"github.com/faujiahmat/zentra-shipping-service/src/interface/service"
)

func InitServer(ss service.Shipping) *server.Restful {
	shippingHandler := handler.NewShipping(ss)
	middleware := middleware.New()

	restfulServer := server.NewRestful(shippingHandler, middleware)
	return restfulServer
}

func InitClient() *client.Restful {
	shipperDelivery := delivery.NewShipper()
	restfulClient := client.NewRestful(shipperDelivery)

	return restfulClient
}
