package main

import (
	"context"
	"fmt"
	pb "github.com/omar-khawaja/grpc-calculator/calculator"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Change this slice of operands to which ever numbers you want to use
	operands := []int32{20, 5, 4}

	// The other methods you can use besides Multiply are
	// Add, Subtract, and Divide
	r, err := c.Multiply(ctx, &pb.Numbers{Operand: operands})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("The answer is %d\n", r.Result)
}
