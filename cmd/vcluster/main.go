package main

import (
	"fmt"

	"github.com/loft-sh/vcluster/cmd/vcluster/cmd"

	// Make sure dep tools picks up these dependencies
	_ "github.com/go-openapi/loads"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	fmt.Println("=========================================================")
	cmd.RunRoot()
}
