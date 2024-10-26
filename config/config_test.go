package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want port %d, got %d", wantPort, got.Port)
	}

	wantEnv := "local"
	if got.Env != wantEnv {
		t.Errorf("want env %s, got %s", wantEnv, got.Env)
	}
}
