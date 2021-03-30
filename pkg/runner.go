package pongo2runner

import (
	"fmt"
	"github.com/flosch/pongo2/v4"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const Pongo2RunnerNamespaceFilter = "_pongo2Runner_namespace"

type Pongo2Runner struct {
	source io.Reader
	directory string
}

func (p *Pongo2Runner) Directory() string {
	return p.directory
}

func (p *Pongo2Runner) SetDirectory(directory string) {
	p.directory = directory
}

func New(source io.Reader) Pongo2Runner {
	// runnerDir is either CWD or /tmp
	runnerDir, err := os.Getwd()
	if err != nil {
		runnerDir = os.TempDir()
	}
	registerFilters()
	return Pongo2Runner{
		source: source,
		directory: runnerDir,
	}
}

/*
	Render renders the template and returns it as a string
*/
func (p *Pongo2Runner) Render() (string, error) {
	bytesArr, err := ioutil.ReadAll(p.source)
	if err != nil {
		return "", fmt.Errorf("unable to read from source: %v", err)
	}

	template, err := pongo2.FromBytes(bytesArr)
	if err != nil {
		return "", fmt.Errorf("unable to parse template: %v", err)
	}

	ctx := pongo2.Context{}
	envMap := map[string]string{}

	for _, v := range os.Environ() {
		envVar := strings.SplitN(v, "=", 2)
		envMap[envVar[0]] = envVar[1]
	}

	ctx["env"] = envMap
	template.Options.AutoEscape = false
	result, err := template.Execute(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to render template: %v", err)
	}

	return result, nil
}