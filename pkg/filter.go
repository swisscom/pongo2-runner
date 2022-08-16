package pongo2runner

import (
	"encoding/base64"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"os"
)

func nullFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return nil, nil
}

func base64EncodeFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(base64.StdEncoding.EncodeToString([]byte(in.String()))), nil
}

func base64DecodeFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	decoded, err := base64.StdEncoding.DecodeString(in.String())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to decode b64 string")
		return nil, nil
	}

	return pongo2.AsValue(string(decoded)), nil
}

func defaultFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if in.String() == "" {
		return param, nil
	}

	return in, nil
}

func registerFilters() {
	if pongo2.FilterExists(NamespaceFilter) {
		return
	}

	_ = pongo2.RegisterFilter(NamespaceFilter, nullFilter)
	_ = pongo2.RegisterFilter("b64encode", base64EncodeFilter)
	_ = pongo2.RegisterFilter("b64decode", base64DecodeFilter)
	_ = pongo2.RegisterFilter("default", defaultFilter)
}
