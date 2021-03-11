package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
	pongo2runner "github.com/swisscom/pongo2-runner/pkg"
	"os"
)

var args struct {
	Source string `arg:"positional,required"`
	Debug  *bool  `arg:"-d"`
}

func main() {
	arg.MustParse(&args)
	logger := logrus.New()
	if args.Debug != nil {
		logger.SetLevel(logrus.DebugLevel)
	}

	file, err := os.Open(args.Source)
	if err != nil {
		logger.Fatalf("unable to open source file: %v", err)
	}

	runner := pongo2runner.New(file)
	result, err := runner.Render()

	if err != nil {
		logger.Fatalf("unable to render template: %v", err)
	}

	fmt.Printf(result)
}
