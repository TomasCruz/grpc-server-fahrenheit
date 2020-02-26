package presenter

import (
	context "context"

	"github.com/TomasCruz/grpc-server-fahrenheit/api"
	"github.com/TomasCruz/grpc-server-fahrenheit/model"
)

// Server represents the gRPC server
type Server struct {
}

// F2C calculates temperature in C from F
func (s *Server) F2C(ctx context.Context, req *api.ConversionMsg) (cm *api.ConversionMsg, err error) {
	f := req.GetNumber()

	c, err := model.F2C(f)
	cm = &api.ConversionMsg{Number: c}

	return
}

// C2F calculates temperature in F from C
func (s *Server) C2F(ctx context.Context, req *api.ConversionMsg) (cm *api.ConversionMsg, err error) {
	c := req.GetNumber()

	f, err := model.C2F(c)
	cm = &api.ConversionMsg{Number: f}

	return
}
