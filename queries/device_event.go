package queries

import (
	"fmt"
	"go-cqrs-api/domain"
)

func ProcessDeviceEvent(event domain.DeviceEvent) error {
	// Your logic here, e.g., insert event into DB
	fmt.Printf("Processing device event: %+v\n", event)
	return nil
}
