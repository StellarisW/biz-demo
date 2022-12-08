// Code generated by Kitex v0.4.3. DO NOT EDIT.

package paymentsvc

import (
	payment "github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/payment"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler payment.PaymentSvc, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
