package presenter

import (
	context "context"
	"fmt"
	"strings"

	"github.com/TomasCruz/grpc-server-fahrenheit/callstack"
	"google.golang.org/grpc"
)

// UnaryInterceptorChainer is shamelessly copied from go-grpc-middleware.ChainUnaryServer
func UnaryInterceptorChainer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		chainer := func(
			currentInter grpc.UnaryServerInterceptor,
			currentHandler grpc.UnaryHandler) grpc.UnaryHandler {

			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}

		chainedHandler := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req)
	}
}

// ErrRepackagingInterceptor logs  and then repackages errors to be returned to the user
func ErrRepackagingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		callstack.LogErrStack(err)
		errMsg := err.Error()

		if strings.Contains(errMsg, "connect: connection refused") {
			errMsg = "DB connection error"
		} else if colonIndex := strings.Index(errMsg, ":"); colonIndex != -1 {
			errMsg = errMsg[:colonIndex]
		}

		err = fmt.Errorf(errMsg)
	}

	return
}

// DummyInterceptor is a dummy, showcasing interceptor chaining
func DummyInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	return handler(ctx, req)
}
