package main

import (
	"bytes"
	"github.com/ReSc/fmt"

	. "github.com/ReSc/godel/core/reflect"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	run()
}

func run() {
	model, err := LoadModelFile("./model.xml")
	PanicIf(err)

	for _, name := range model.Packages.SortedKeys() {
		if name == "global" {
			continue
		}
		p := model.Packages[name]
		file := render(p)
		reformat(file)
	}
}

func render(p *Package) string {
	dir := p.GetPackageDirectory()
	file := filepath.Join(dir, p.Name+".go")
	fmt.Printline("rendering package %v: %v", p.Name, file)
	PanicIf(os.MkdirAll(dir, os.FileMode(0777)))
	outf, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	PanicIf(err)
	defer outf.Close()

	NewModelRenderer(outf).RenderPackage(p)
	return file
}

func reformat(file string) {
	cmd := exec.Command("gofmt", "-w", file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printline(err.Error() + ": " + stderr.String())
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
