package broker

import (
	"github.com/faujiahmat/zentra-shipping-service/src/core/broker/consumer"
	"github.com/faujiahmat/zentra-shipping-service/src/core/broker/handler"
	"github.com/faujiahmat/zentra-shipping-service/src/interface/service"
)

func InitShipperConsumer(ns service.Notification) *consumer.ShipperKafka {
	shipperHandler := handler.NewShipperKafka(ns)
	shipperConsumer := consumer.NewShipperKafka(shipperHandler)

	return shipperConsumer
}
