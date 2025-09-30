package helper

import "github.com/faujiahmat/zentra-shipping-service/src/model/entity"

func FormatPickupReq(shippingIds []string) *entity.Pickup {
	return &entity.Pickup{
		Data: entity.PickupData{
			OrderActivation: entity.OrderActivation{
				OrderId: shippingIds,
			},
		},
	}
}
