package pongo2runner

import (
	"encoding/base64"
	"github.com/flosch/pongo2/v4"
)

func nullFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error){
	return nil, nil
}

func base64EncodeFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error){
	return pongo2.AsValue(base64.StdEncoding.EncodeToString([]byte(in.String()))), nil
}

func registerFilters() {
	if pongo2.FilterExists(Pongo2RunnerNamespaceFilter) {
		return
	}

	_ = pongo2.RegisterFilter(Pongo2RunnerNamespaceFilter, nullFilter)
	_ = pongo2.RegisterFilter("b64encode", base64EncodeFilter)
}