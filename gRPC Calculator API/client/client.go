package client

import (
	"context"
	"fmt"
	config "gRPC-example/config"
	api "gRPC-example/definitions"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

func Run() {
	sugar.Info("GRPC Client")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sugar.Fatal(err)
		os.Exit(1)
	}
	sugar.Info("Connection established with server")

	defer conn.Close()
	c := api.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := get_input()

	response, err := c.Calculate(ctx, &request)

	if err != nil {
		sugar.Fatal(err)
		os.Exit(1)
	}

	sugar.Info("Got result: ", response.Result)

}

func get_input() api.Input {
	var operand1, operand2, temp_operator int32
	operator := api.Operations_ADDITION
	fmt.Println("Enter Operand1")
	fmt.Scanf("%d", &operand1)
	fmt.Println("Enter Operand2")
	fmt.Scanf("%d", &operand2)
	fmt.Println("Enter Operation Number: \n[1] Addition\n[2] Subtraction\n[3] Multiplication\n[4] Division")
	fmt.Scanf("%d", &temp_operator)
	switch temp_operator {
	case 2:
		operator = api.Operations_SUBTRACTION
	case 3:
		operator = api.Operations_MULTIPLICATION
	case 4:
		operator = api.Operations_DIVISION
	}

	return api.Input{Operand1: operand1, Operand2: operand2, Operation: operator}
}
