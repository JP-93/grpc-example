package health

import (
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Check struct {
	Serving    string `json:"serving"`
	NotServing string `json:"not_serving"`
	Unknown    string `json:"unknown"`
}

type Health interface {
	Check(ctx context.Context, req *healthpb.HealthCheckResponse) (*Check, error)
}

type health struct {
}

func (h *health) Check(ctx context.Context, req *healthpb.HealthCheckResponse) (*Check, error) {

	var res *Check

	if req.Status == healthpb.HealthCheckResponse_SERVING {
		res.Serving = string(healthpb.HealthCheckResponse_SERVING)
	} else if req.Status == healthpb.HealthCheckResponse_NOT_SERVING {
		res.NotServing = string(healthpb.HealthCheckResponse_NOT_SERVING)
	}

	return res, nil
}

func NewHealth() Health {
	return &health{}
}
