package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}

	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動
	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to start server: %v", err)
			return err
		}
		return nil
	})

	// シグナルを待つ
	<-ctx.Done()
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown server: %v", err)
	}

	// Goメソッドで起動した別ゴルーチンのエラーを待つ
	return eg.Wait()
}
