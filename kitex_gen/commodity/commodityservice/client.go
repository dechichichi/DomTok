/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.12.1. DO NOT EDIT.

package commodityservice

import (
	"context"

	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"

	commodity "github.com/west2-online/DomTok/kitex_gen/commodity"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateCoupon(ctx context.Context, req *commodity.CreateCouponReq, callOptions ...callopt.Option) (r *commodity.CreateCouponResp, err error)
	DeleteCoupon(ctx context.Context, req *commodity.DeleteCouponReq, callOptions ...callopt.Option) (r *commodity.DeleteCouponResp, err error)
	CreateUserCoupon(ctx context.Context, req *commodity.CreateUserCouponReq, callOptions ...callopt.Option) (r *commodity.CreateUserCouponResp, err error)
	ViewCoupon(ctx context.Context, req *commodity.ViewCouponReq, callOptions ...callopt.Option) (r *commodity.ViewCouponResp, err error)
	ViewUserAllCoupon(ctx context.Context, req *commodity.ViewUserAllCouponReq, callOptions ...callopt.Option) (r *commodity.ViewUserAllCouponResp, err error)
	UseUserCoupon(ctx context.Context, req *commodity.UseUserCouponReq, callOptions ...callopt.Option) (r *commodity.UseUserCouponResp, err error)
	ViewSpu(ctx context.Context, req *commodity.ViewSpuReq, callOptions ...callopt.Option) (r *commodity.ViewSpuResp, err error)
	DeleteSpu(ctx context.Context, req *commodity.DeleteSpuReq, callOptions ...callopt.Option) (r *commodity.DeleteSpuResp, err error)
	ViewSpuImage(ctx context.Context, req *commodity.ViewSpuImageReq, callOptions ...callopt.Option) (r *commodity.ViewSpuImageResp, err error)
	DeleteSpuImage(ctx context.Context, req *commodity.DeleteSpuImageReq, callOptions ...callopt.Option) (r *commodity.DeleteSpuImageResp, err error)
	DeleteSku(ctx context.Context, req *commodity.DeleteSkuReq, callOptions ...callopt.Option) (r *commodity.DeleteSkuResp, err error)
	ViewSkuImage(ctx context.Context, req *commodity.ViewSkuImageReq, callOptions ...callopt.Option) (r *commodity.ViewSkuImageResp, err error)
	ViewSku(ctx context.Context, req *commodity.ViewSkuReq, callOptions ...callopt.Option) (r *commodity.ViewSkuResp, err error)
	UploadSkuAttr(ctx context.Context, req *commodity.UploadSkuAttrReq, callOptions ...callopt.Option) (r *commodity.UploadSkuAttrResp, err error)
	ListSkuInfo(ctx context.Context, req *commodity.ListSkuInfoReq, callOptions ...callopt.Option) (r *commodity.ListSkuInfoResp, err error)
	ViewHistory(ctx context.Context, req *commodity.ViewHistoryPriceReq, callOptions ...callopt.Option) (r *commodity.ViewHistoryPriceResp, err error)
	DeleteSkuImage(ctx context.Context, req *commodity.DeleteSkuImageReq, callOptions ...callopt.Option) (r *commodity.DeleteSkuImageResp, err error)
	DescSkuLockStock(ctx context.Context, req *commodity.DescSkuLockStockReq, callOptions ...callopt.Option) (r *commodity.DescSkuLockStockResp, err error)
	IncrSkuLockStock(ctx context.Context, req *commodity.IncrSkuLockStockReq, callOptions ...callopt.Option) (r *commodity.IncrSkuLockStockResp, err error)
	DescSkuStock(ctx context.Context, req *commodity.DescSkuStockReq, callOptions ...callopt.Option) (r *commodity.DescSkuStockResp, err error)
	ListSpuInfo(ctx context.Context, req *commodity.ListSpuInfoReq, callOptions ...callopt.Option) (r *commodity.ListSpuInfoResp, err error)
	CreateCategory(ctx context.Context, req *commodity.CreateCategoryReq, callOptions ...callopt.Option) (r *commodity.CreateCategoryResp, err error)
	DeleteCategory(ctx context.Context, req *commodity.DeleteCategoryReq, callOptions ...callopt.Option) (r *commodity.DeleteCategoryResp, err error)
	ViewCategory(ctx context.Context, req *commodity.ViewCategoryReq, callOptions ...callopt.Option) (r *commodity.ViewCategoryResp, err error)
	UpdateCategory(ctx context.Context, req *commodity.UpdateCategoryReq, callOptions ...callopt.Option) (r *commodity.UpdateCategoryResp, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	CreateSpu(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSpuClient, err error)
	UpdateSpu(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSpuClient, err error)
	CreateSpuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSpuImageClient, err error)
	UpdateSpuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSpuImageClient, err error)
	CreateSku(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSkuClient, err error)
	UpdateSku(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSkuClient, err error)
	CreateSkuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSkuImageClient, err error)
	UpdateSkuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSkuImageClient, err error)
}

type CommodityService_CreateSpuClient interface {
	streaming.Stream
	Send(*commodity.CreateSpuReq) error
	CloseAndRecv() (*commodity.CreateSpuResp, error)
}

type CommodityService_UpdateSpuClient interface {
	streaming.Stream
	Send(*commodity.UpdateSpuReq) error
	CloseAndRecv() (*commodity.UpdateSpuResp, error)
}

type CommodityService_CreateSpuImageClient interface {
	streaming.Stream
	Send(*commodity.CreateSpuImageReq) error
	CloseAndRecv() (*commodity.CreateSpuImageResp, error)
}

type CommodityService_UpdateSpuImageClient interface {
	streaming.Stream
	Send(*commodity.UpdateSpuImageReq) error
	CloseAndRecv() (*commodity.UpdateSpuImageResp, error)
}

type CommodityService_CreateSkuClient interface {
	streaming.Stream
	Send(*commodity.CreateSkuReq) error
	CloseAndRecv() (*commodity.CreateSkuResp, error)
}

type CommodityService_UpdateSkuClient interface {
	streaming.Stream
	Send(*commodity.UpdateSkuReq) error
	CloseAndRecv() (*commodity.UpdateSkuResp, error)
}

type CommodityService_CreateSkuImageClient interface {
	streaming.Stream
	Send(*commodity.CreateSkuImageReq) error
	CloseAndRecv() (*commodity.CreateSkuImageResp, error)
}

type CommodityService_UpdateSkuImageClient interface {
	streaming.Stream
	Send(*commodity.UpdateSkuImageReq) error
	CloseAndRecv() (*commodity.UpdateSkuImageResp, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommodityServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommodityServiceClient struct {
	*kClient
}

func (p *kCommodityServiceClient) CreateCoupon(ctx context.Context, req *commodity.CreateCouponReq, callOptions ...callopt.Option) (r *commodity.CreateCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCoupon(ctx, req)
}

func (p *kCommodityServiceClient) DeleteCoupon(ctx context.Context, req *commodity.DeleteCouponReq, callOptions ...callopt.Option) (r *commodity.DeleteCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCoupon(ctx, req)
}

func (p *kCommodityServiceClient) CreateUserCoupon(ctx context.Context, req *commodity.CreateUserCouponReq, callOptions ...callopt.Option) (r *commodity.CreateUserCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUserCoupon(ctx, req)
}

func (p *kCommodityServiceClient) ViewCoupon(ctx context.Context, req *commodity.ViewCouponReq, callOptions ...callopt.Option) (r *commodity.ViewCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewCoupon(ctx, req)
}

func (p *kCommodityServiceClient) ViewUserAllCoupon(ctx context.Context, req *commodity.ViewUserAllCouponReq, callOptions ...callopt.Option) (r *commodity.ViewUserAllCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewUserAllCoupon(ctx, req)
}

func (p *kCommodityServiceClient) UseUserCoupon(ctx context.Context, req *commodity.UseUserCouponReq, callOptions ...callopt.Option) (r *commodity.UseUserCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UseUserCoupon(ctx, req)
}

func (p *kCommodityServiceClient) ViewSpu(ctx context.Context, req *commodity.ViewSpuReq, callOptions ...callopt.Option) (r *commodity.ViewSpuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSpu(ctx, req)
}

func (p *kCommodityServiceClient) DeleteSpu(ctx context.Context, req *commodity.DeleteSpuReq, callOptions ...callopt.Option) (r *commodity.DeleteSpuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSpu(ctx, req)
}

func (p *kCommodityServiceClient) ViewSpuImage(ctx context.Context, req *commodity.ViewSpuImageReq, callOptions ...callopt.Option) (r *commodity.ViewSpuImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSpuImage(ctx, req)
}

func (p *kCommodityServiceClient) DeleteSpuImage(ctx context.Context, req *commodity.DeleteSpuImageReq, callOptions ...callopt.Option) (r *commodity.DeleteSpuImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSpuImage(ctx, req)
}

func (p *kCommodityServiceClient) DeleteSku(ctx context.Context, req *commodity.DeleteSkuReq, callOptions ...callopt.Option) (r *commodity.DeleteSkuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSku(ctx, req)
}

func (p *kCommodityServiceClient) ViewSkuImage(ctx context.Context, req *commodity.ViewSkuImageReq, callOptions ...callopt.Option) (r *commodity.ViewSkuImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSkuImage(ctx, req)
}

func (p *kCommodityServiceClient) ViewSku(ctx context.Context, req *commodity.ViewSkuReq, callOptions ...callopt.Option) (r *commodity.ViewSkuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSku(ctx, req)
}

func (p *kCommodityServiceClient) UploadSkuAttr(ctx context.Context, req *commodity.UploadSkuAttrReq, callOptions ...callopt.Option) (r *commodity.UploadSkuAttrResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadSkuAttr(ctx, req)
}

func (p *kCommodityServiceClient) ListSkuInfo(ctx context.Context, req *commodity.ListSkuInfoReq, callOptions ...callopt.Option) (r *commodity.ListSkuInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListSkuInfo(ctx, req)
}

func (p *kCommodityServiceClient) ViewHistory(ctx context.Context, req *commodity.ViewHistoryPriceReq, callOptions ...callopt.Option) (r *commodity.ViewHistoryPriceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewHistory(ctx, req)
}

func (p *kCommodityServiceClient) DeleteSkuImage(ctx context.Context, req *commodity.DeleteSkuImageReq, callOptions ...callopt.Option) (r *commodity.DeleteSkuImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSkuImage(ctx, req)
}

func (p *kCommodityServiceClient) DescSkuLockStock(ctx context.Context, req *commodity.DescSkuLockStockReq, callOptions ...callopt.Option) (r *commodity.DescSkuLockStockResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescSkuLockStock(ctx, req)
}

func (p *kCommodityServiceClient) IncrSkuLockStock(ctx context.Context, req *commodity.IncrSkuLockStockReq, callOptions ...callopt.Option) (r *commodity.IncrSkuLockStockResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncrSkuLockStock(ctx, req)
}

func (p *kCommodityServiceClient) DescSkuStock(ctx context.Context, req *commodity.DescSkuStockReq, callOptions ...callopt.Option) (r *commodity.DescSkuStockResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescSkuStock(ctx, req)
}

func (p *kCommodityServiceClient) ListSpuInfo(ctx context.Context, req *commodity.ListSpuInfoReq, callOptions ...callopt.Option) (r *commodity.ListSpuInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListSpuInfo(ctx, req)
}

func (p *kCommodityServiceClient) CreateCategory(ctx context.Context, req *commodity.CreateCategoryReq, callOptions ...callopt.Option) (r *commodity.CreateCategoryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCategory(ctx, req)
}

func (p *kCommodityServiceClient) DeleteCategory(ctx context.Context, req *commodity.DeleteCategoryReq, callOptions ...callopt.Option) (r *commodity.DeleteCategoryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCategory(ctx, req)
}

func (p *kCommodityServiceClient) ViewCategory(ctx context.Context, req *commodity.ViewCategoryReq, callOptions ...callopt.Option) (r *commodity.ViewCategoryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewCategory(ctx, req)
}

func (p *kCommodityServiceClient) UpdateCategory(ctx context.Context, req *commodity.UpdateCategoryReq, callOptions ...callopt.Option) (r *commodity.UpdateCategoryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCategory(ctx, req)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommodityServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommodityServiceStreamClient struct {
	*kClient
}

func (p *kCommodityServiceStreamClient) CreateSpu(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSpuClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.CreateSpu(ctx)
}

func (p *kCommodityServiceStreamClient) UpdateSpu(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSpuClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.UpdateSpu(ctx)
}

func (p *kCommodityServiceStreamClient) CreateSpuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSpuImageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.CreateSpuImage(ctx)
}

func (p *kCommodityServiceStreamClient) UpdateSpuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSpuImageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.UpdateSpuImage(ctx)
}

func (p *kCommodityServiceStreamClient) CreateSku(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSkuClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.CreateSku(ctx)
}

func (p *kCommodityServiceStreamClient) UpdateSku(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSkuClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.UpdateSku(ctx)
}

func (p *kCommodityServiceStreamClient) CreateSkuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_CreateSkuImageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.CreateSkuImage(ctx)
}

func (p *kCommodityServiceStreamClient) UpdateSkuImage(ctx context.Context, callOptions ...streamcall.Option) (stream CommodityService_UpdateSkuImageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.UpdateSkuImage(ctx)
}
