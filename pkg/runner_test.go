package pongo2runner_test

import (
	"github.com/stretchr/testify/assert"
	pongo2runner "github.com/swisscom/pongo2-runner/pkg"
	"io/ioutil"
	"os"
	"testing"
)

func TestPongo2Runner_Render(t *testing.T) {
	dirEntries, err := os.ReadDir("../examples/")
	if err != nil {
		t.Fatalf("unable to list dir contents: %v", err)
	}

	os.Clearenv()
	os.Setenv("CC", "clang")
	os.Setenv("SHELL", "/bin/zsh")

	for _, e := range dirEntries {
		if !e.IsDir() {
			testRender(t, e.Name())
		}
	}
}

func testRender(t *testing.T, fileName string) {
	file, err := os.Open("../examples/" + fileName)
	if err != nil {
		t.Fatalf("unable to open test file: %v", err)
	}

	expectedResult, err := os.Open("../examples/results/" + fileName)
	if err != nil {
		t.Fatalf("unable to open result file: %v", err)
	}

	pr := pongo2runner.New(file)
	res, err := pr.Render()

	if err != nil {
		t.Fatalf("unable to render template: %v", err)
	}

	resultBytes, err := ioutil.ReadAll(expectedResult)
	if err != nil {
		t.Fatalf("unable to read expected result file: %v", err)
	}

	assert.Equal(t, string(resultBytes), res)
}
