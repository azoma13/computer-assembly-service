package v1

import (
	"github.com/azoma13/computer-assembly-service/order/internal/client/grpc"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
)

var _ grpc.HardwareClient = (*client)(nil)

type client struct {
	generatedClient hardware_v1.HardwareServiceClient
}

func NewHardwareClient(generatedClient hardware_v1.HardwareServiceClient) *client {
	return &client{
		generatedClient: generatedClient,
	}
}
