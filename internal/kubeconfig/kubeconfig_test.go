package kubeconfig

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConnectionDetails(t *testing.T) {
	defer func(f func(string, string) ([]byte, error)) { fetch = f }(fetch)

	var fetchErr error
	fetch = func(id, apiKey string) ([]byte, error) {
		if fetchErr != nil {
			return nil, fetchErr
		}
		return []byte(id + "/" + apiKey), nil
	}
	attr := map[string]any{"id": "c1"}

	check := func(step string, want map[string][]byte) {
		t.Helper()
		got, err := ConnectionDetails(attr)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", step, err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("%s: -want +got:\n%s", step, diff)
		}
	}

	check("unknown cluster", nil)

	fetchErr = errors.New("not active")
	SetAPIKey("c1", "k1")
	check("not active, nothing fetched yet", nil)

	fetchErr = nil
	check("active", map[string][]byte{Key: []byte("c1/k1")})

	fetchErr = errors.New("not active")
	check("not active, keeps last kubeconfig", map[string][]byte{Key: []byte("c1/k1")})

	fetchErr = nil
	SetAPIKey("c1", "k2")
	check("new API key", map[string][]byte{Key: []byte("c1/k2")})

	Forget("c1")
	check("forgotten", nil)
}
