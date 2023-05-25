package server

import (
	"errors"
	"fmt"
	config "gRPC-example/config"
	api "gRPC-example/definitions"
	"net"
	"os"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

func Run() {
	defer logger.Sync()
	sugar.Info("Calculator GRPC server")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		sugar.Fatal(err)
		os.Exit(1)
	}

	server := grpc.NewServer()

	api.RegisterCalculatorServer(server, &CalculatorServer{})

	err = server.Serve(listener)

	if err != nil {
		sugar.Fatal("Failed to server GRPC server")
		os.Exit(1)
	}

}

type CalculatorServer struct {
	api.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Calculate(ctx context.Context, request *api.Input) (*api.Output, error) {

	sugar.Info("Got Payload:", " Operand1: ", request.Operand1, " Operand2: ", request.Operand2, " Operation: ", request.Operation)
	var result int64
	switch request.Operation {
	case api.Operations_ADDITION:
		result = int64(request.Operand1) + int64(request.Operand2)
	case api.Operations_SUBTRACTION:
		result = int64(request.Operand1) - int64(request.Operand2)
	case api.Operations_MULTIPLICATION:
		result = int64(request.Operand1) * int64(request.Operand2)
	case api.Operations_DIVISION:
		if request.Operand2 == 0 {
			return &api.Output{Operand1: request.Operand1, Operand2: request.Operand2, Result: result, Operation: request.Operation}, errors.New("division by zero is not possible")
		}
		result = int64(request.Operand1) / int64(request.Operand2)
	}

	return &api.Output{Operand1: request.Operand1, Operand2: request.Operand2, Result: result, Operation: request.Operation}, nil
}
