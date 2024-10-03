package sampleExtServiceUseCase

import (
	"context"

	articleV1 "golang-oauth2-server/api/article/v1"

	sampleExtServiceDomain "golang-oauth2-server/external/sample_ext_service/domain"
	grpcError "golang-oauth2-server/pkg/error/grpc"
	"golang-oauth2-server/pkg/grpc"
)

type sampleExtServiceUseCase struct {
	grpcClient grpc.Client
}

func NewSampleExtServiceUseCase(grpcClient grpc.Client) sampleExtServiceDomain.SampleExtServiceUseCase {
	return &sampleExtServiceUseCase{
		grpcClient: grpcClient,
	}
}

func (esu *sampleExtServiceUseCase) CreateUsers(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	articleGrpcClient := articleV1.NewArticleServiceClient(esu.grpcClient.GetGrpcConnection())

	res, err := articleGrpcClient.CreateArticle(ctx, req)
	if err != nil {
		return nil, grpcError.ParseExternalGrpcErr(err)
	}

	return res, nil
}
