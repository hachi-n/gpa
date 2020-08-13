//go:generate pkger -o generated

package main

import (
	"fmt"
	_ "github.com/hachi-n/gpa/generated"
	"github.com/markbates/pkger"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Yaml struct {
	Hoge string `yaml:"hoge"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := pkger.Open("/templates/application.yml")
	if err != nil {
		return err
	}

	info, _ := f.Stat()
	b := make([]byte, info.Size())
	f.Read(b)

	y := new(Yaml)
	if err := yaml.Unmarshal(b, y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(y)

	return nil
}
