# go-httpc
Go httpc client

This is a fork of the backup from [gmkit](https://github.com/graymeta/gmkit) with external dependencies removed.

example:
```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jasonhancock/go-backoff"
	"github.com/jasonhancock/go-logger"
	"github.com/ns-jsattler/go-httpc"
)

func main() {
	l := logger.New(os.Stdout, "httpctest", "info", logger.FormatJSON)
	argoUrl, ok := os.LookupEnv("ARGOCD_SERVER")
	if !ok {
		l.Fatal("ARGO_SERVER environment variable not set")
	}

	token, ok := os.LookupEnv("ARGOCD_AUTH_TOKEN")
	if !ok {
		l.Fatal("ARGOCD_AUTH_TOKEN environment variable not set")
	}

	boff := backoff.New(backoff.WithLogger(l))

	opts := []httpc.ClientOptFn{
		httpc.WithBaseURL(argoUrl),
		httpc.WithAuth(httpc.BearerTokenAuth(token)),
		httpc.WithRetryClientTimeouts(),
		httpc.WithRetryResponseErrors(),
		httpc.WithBackoff(boff),
		httpc.WithEncode(httpc.JSONEncode()),
	}

	c := &http.Client{Timeout: 30 * time.Second}

	client := httpc.New(c, opts...)
	var resp ArgoClusterResponse
	err := client.GET("/api/v1/clusters").
		ContentType("application/json").
		Success(httpc.StatusOK()).
		DecodeJSON(&resp).
		Do(context.TODO())
	if err != nil {
		l.Fatal(err)
	}

	for _, v := range resp.Items {
		fmt.Printf("Cluster: %s\n", v.Name)
	}
}

type ArgoClusterResponse struct {
	Items []struct {
		Annotations map[string]string `json:"annotations"`
		Labels      map[string]string `json:"labels"`
		Name        string            `json:"name"`
		Namespaces  []string          `json:"namespaces"`
		Project     string            `json:"project"`
		Server      string            `json:"server"`
	} `json:"items"`
	Metadata struct {
		Continue           string `json:"continue"`
		RemainingItemCount string `json:"remainingItemCount"`
		ResourceVersion    string `json:"resourceVersion"`
		SelfLink           string `json:"selfLink"`
	} `json:"metadata"`
}

```