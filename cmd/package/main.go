package main

import (
	"context"
	"time"

	schedulerx "github.com/agrawaltejas01/schedulerx/pkg"
)

func main() {
	ctx := context.Background()
	schedulerFreq := 3 * time.Second
	executorFreq := 5 * time.Second

	schedulerx.NewSchedulerX(ctx, schedulerFreq, executorFreq)

}
