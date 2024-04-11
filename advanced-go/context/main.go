package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	start := time.Now()

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	ctxCh := make(chan slowFunctionResponse)

	go slowFunction(ctxCh)

	for {
		select {
		case <-ctx.Done():
			log.Fatalf("[ERROR] task not completed under the allowed time")
			return

		case r := <-ctxCh:
			fmt.Println("[SUCCESS] task succeeded under the allowed time")

			fmt.Printf("\nthe task took %v to complete \n", time.Since(start))

			fmt.Printf("\n%+v\n", r)
			return
		}
	}
}

type slowFunctionResponse struct {
	status int
	price  float64
	err    error
}

func slowFunction(ctxCh chan slowFunctionResponse) {
	ctxCh <- thirdPartyRequest()
}

func thirdPartyRequest() slowFunctionResponse {
	time.Sleep(time.Second * 2)

	return slowFunctionResponse{
		status: 200,
		price:  98.29,
		err:    nil,
	}
}
