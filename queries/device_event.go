package queries

import (
	"fmt"
	"go-cqrs-api/domain"
)

func ProcessDeviceEvent(event domain.DeviceEvent) error {
	fmt.Printf("Processing device event: %+v\n", event)
	return nil
}
