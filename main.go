package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("server listening at %s", url)
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
