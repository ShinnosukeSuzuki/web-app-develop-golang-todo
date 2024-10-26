package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	t.Skip("リファクタリング中")

	//  portに0を指定することで、システムが自動的に利用可能なポートを選択
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	// テスト用のコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	t.Logf("try request to %q", url)
	rep, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %v", err)
	}
	defer rep.Body.Close()
	got, err := io.ReadAll(rep.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// HTTPサーバーの戻り値を検証
	want := fmt.Sprintf("%s\n", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数に終了通知を送る
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
