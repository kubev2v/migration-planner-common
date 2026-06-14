//go:build tools
// +build tools

package tools

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
	_ "golang.org/x/tools/cmd/goimports"
)
