package healthCheckGrpc

import (
	"context"

	"google.golang.org/grpc/codes"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	healthCheckDomain "golang-oauth2-server/internal/modules/health_check/domain"
)

type controller struct {
	healthCheckUseCase    healthCheckDomain.HealthCheckUseCase
	postgresHealthCheckUc healthCheckDomain.PostgresHealthCheckUseCase
	kafkaHealthCheckUc    healthCheckDomain.KafkaHealthCheckUseCase
	tmpDirHealthCheckUc   healthCheckDomain.TmpDirHealthCheckUseCase
}

func NewController(
	healthCheckUc healthCheckDomain.HealthCheckUseCase,
	postgresHealthCheckUc healthCheckDomain.PostgresHealthCheckUseCase,
	kafkaHealthCheckUc healthCheckDomain.KafkaHealthCheckUseCase,
	tmpDirHealthCheckUc healthCheckDomain.TmpDirHealthCheckUseCase) healthCheckDomain.GrpcController {
	return &controller{
		healthCheckUseCase:    healthCheckUc,
		postgresHealthCheckUc: postgresHealthCheckUc,
		kafkaHealthCheckUc:    kafkaHealthCheckUc,
		tmpDirHealthCheckUc:   tmpDirHealthCheckUc,
	}
}

func (c *controller) Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error) {
	var healthStatus bool

	switch request.Service {
	case "", "all":
		healthStatus = c.healthCheckUseCase.Check().Status
	case "kafka":
		healthStatus = c.kafkaHealthCheckUc.Check()
	case "postgres":
		healthStatus = c.postgresHealthCheckUc.Check()
	case "writable-tmp-dir":
		healthStatus = c.tmpDirHealthCheckUc.Check()
	default:
		return &grpcHealthV1.HealthCheckResponse{
			Status: grpcHealthV1.HealthCheckResponse_UNKNOWN,
		}, nil
	}

	grpcStatus := grpcHealthV1.HealthCheckResponse_SERVING

	if !healthStatus {
		grpcStatus = grpcHealthV1.HealthCheckResponse_NOT_SERVING
	}

	return &grpcHealthV1.HealthCheckResponse{
		Status: grpcStatus,
	}, nil
}

func (c *controller) Watch(request *grpcHealthV1.HealthCheckRequest, server grpcHealthV1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
