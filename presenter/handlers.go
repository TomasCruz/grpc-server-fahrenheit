package presenter

import (
	context "context"

	"github.com/TomasCruz/grpc-server-fahrenheit/api"
	"github.com/TomasCruz/grpc-server-fahrenheit/model"
)

// Server represents the gRPC server
type Server struct {
}

// Health returns health status of the service
func (s *Server) Health(ctx context.Context, req *api.NoParamsMsg) (*api.HealthMsg, error) {
	h, err := model.Health()
	return &api.HealthMsg{Health: h}, err
}

// F2C calculates temperature in C from F
func (s *Server) F2C(ctx context.Context, req *api.ConversionMsg) (*api.ConversionMsg, error) {
	f := req.GetNumber()
	c, err := model.F2C(f)
	return &api.ConversionMsg{Number: c}, err
}

// C2F calculates temperature in F from C
func (s *Server) C2F(ctx context.Context, req *api.ConversionMsg) (*api.ConversionMsg, error) {
	c := req.GetNumber()
	f, err := model.C2F(c)
	return &api.ConversionMsg{Number: f}, err
}
