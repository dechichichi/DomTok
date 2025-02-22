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

// Code generated by hertz generator.

package payment

import (
	"context"
	"go.uber.org/zap"

	"github.com/west2-online/DomTok/pkg/logger"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	api "github.com/west2-online/DomTok/app/gateway/model/api/payment"
	"github.com/west2-online/DomTok/app/gateway/pack"
	"github.com/west2-online/DomTok/app/gateway/rpc"
	"github.com/west2-online/DomTok/kitex_gen/payment"
	"github.com/west2-online/DomTok/pkg/errno"
)

// ProcessPayment .
// @router /api/payment/process [POST]
func ProcessPayment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PaymentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// TODO 这鬼地方是api. 还是 payment.
	resp := new(api.PaymentResponse)
	c.JSON(consts.StatusOK, resp)
}

// RequestPaymentToken .
// @router /api/payment/token [GET]
func RequestPaymentToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PaymentTokenRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Error("RequestPaymentTokenRPC failed", zap.Error(err)) // 打印错误日志
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	// 调用 RPC 获取支付令牌
	resp, err := rpc.RequestPaymentTokenRPC(ctx, &payment.PaymentTokenRequest{
		OrderID: req.OrderID,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	// 返回成功的响应
	pack.RespData(c, resp)
}

// ProcessRefund .
// @router /api/payment/refund [POST]
func ProcessRefund(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RefundRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// TODO 这鬼地方是api. 还是 payment.？
	resp := new(payment.RefundResponse)

	c.JSON(consts.StatusOK, resp)
}

// RequestRefundToken .
// @router /api/payment/refund-token [GET]
func RequestRefundToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RefundTokenRequest

	// 解析并校验请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Error("RequestRefundTokenRPC failed", zap.Error(err)) // 记录错误日志
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	// 调用 RPC 获取退款令牌
	resp, err := rpc.RequestRefundRPC(ctx, &payment.RefundTokenRequest{
		OrderID: req.OrderID,
		UserID:  req.UserID,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	// 返回成功响应
	pack.RespData(c, resp)
}
