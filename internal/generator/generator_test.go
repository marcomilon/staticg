package generator_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/marcomilon/gstatic/internal/generator"
)

var out string = os.TempDir() + "gstatictest/public"

func TestFileGenerator(t *testing.T) {

	setup(t)

	yaml := generator.Yaml{"testdata/yaml", out}
	err := generator.Generate(yaml)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := out + string(os.PathSeparator) + "index.html"
	if _, err := os.Stat(index); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	section := out + string(os.PathSeparator) + "/section/section.html"
	if _, err := os.Stat(section); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}
}

func TestYamlParser(t *testing.T) {

	setup(t)

	yaml := generator.Yaml{"testdata/yaml", out}
	err := generator.Generate(yaml)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := out + string(os.PathSeparator) + "index.html"
	indexTpl, err := ioutil.ReadFile(index)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	indexResult := string(indexTpl)
	indexExpected := "<p>Hello world</p>"
	if strings.ToLower(indexResult) != strings.ToLower(indexExpected) {
		t.Errorf("expected %v; got %v", indexExpected, indexResult)
	}

	section := out + string(os.PathSeparator) + "/section/section.html"
	sectionTpl, err := ioutil.ReadFile(section)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	sectionResult := string(sectionTpl)
	sectionExpected := "<p>Marco</p>"
	if strings.ToLower(sectionResult) != strings.ToLower(sectionExpected) {
		t.Errorf("expected %v; got %v", sectionExpected, sectionResult)
	}
}

func setup(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(out, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(out)
}
