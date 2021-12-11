package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/app"
)

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go app.Run(ctx, wg)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
