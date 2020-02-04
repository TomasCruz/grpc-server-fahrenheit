package presenter

import (
	context "context"

	"github.com/TomasCruz/grpc-server-fahrenheit/model"
)

// Server represents the gRPC server
type Server struct {
}

// Health returns health status of the service
func (s *Server) Health(ctx context.Context, req *NoParamsMsg) (*HealthMsg, error) {
	h, err := model.Health()
	return &HealthMsg{Health: h}, err
}

// F2C calculates temperature in C from F
func (s *Server) F2C(ctx context.Context, req *ConversionMsg) (*ConversionMsg, error) {
	f := req.GetInput()
	c, err := model.F2C(f)
	return &ConversionMsg{Input: c}, err
}

// C2F calculates temperature in F from C
func (s *Server) C2F(ctx context.Context, req *ConversionMsg) (*ConversionMsg, error) {
	c := req.GetInput()
	f, err := model.C2F(c)
	return &ConversionMsg{Input: f}, err
}
