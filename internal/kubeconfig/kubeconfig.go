// Package kubeconfig adds the VKE kubeconfig to the connection details of
// Kubernetes resources. The Terraform provider stopped returning it on the
// vultr_kubernetes resource in v2.33.0 and moved it to a data source, which
// upjet cannot use.
package kubeconfig

import (
	"context"
	"encoding/base64"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vultr/govultr/v3"
	"golang.org/x/oauth2"
)

// Key is the connection secret key holding the kubeconfig.
const Key = "kubeconfig"

type cluster struct {
	apiKey     string
	kubeconfig []byte
}

var (
	mu       sync.Mutex
	clusters = map[string]cluster{}
)

// SetAPIKey records the API key to use when fetching the cluster's kubeconfig.
func SetAPIKey(clusterID, apiKey string) {
	mu.Lock()
	defer mu.Unlock()
	c := clusters[clusterID]
	c.apiKey = apiKey
	clusters[clusterID] = c
}

// Forget removes the cluster so its kubeconfig is no longer fetched.
func Forget(clusterID string) {
	mu.Lock()
	defer mu.Unlock()
	delete(clusters, clusterID)
}

// ConnectionDetails fetches the kubeconfig of the cluster in attr. The API
// refuses while the cluster is not active (e.g. upgrading), so the last
// kubeconfig fetched is returned instead of failing the reconcile.
func ConnectionDetails(attr map[string]any) (map[string][]byte, error) {
	id, _ := attr["id"].(string)
	mu.Lock()
	c, ok := clusters[id]
	mu.Unlock()
	if id == "" || !ok {
		return nil, nil
	}

	kc, err := fetch(id, c.apiKey)
	if err != nil {
		if c.kubeconfig == nil {
			return nil, nil
		}
		return map[string][]byte{Key: c.kubeconfig}, nil
	}

	mu.Lock()
	if cur, ok := clusters[id]; ok {
		cur.kubeconfig = kc
		clusters[id] = cur
	}
	mu.Unlock()
	return map[string][]byte{Key: kc}, nil
}

var fetch = func(id, apiKey string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})
	client := govultr.NewClient(oauth2.NewClient(ctx, ts))
	kc, _, err := client.Kubernetes.GetKubeConfig(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get kubeconfig")
	}
	raw, err := base64.StdEncoding.DecodeString(kc.KubeConfig)
	return raw, errors.Wrap(err, "cannot decode kubeconfig")
}
