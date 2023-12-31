package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sebastianrau/go-easyConfig/pkg/demo"

	easyconfig "github.com/sebastianrau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configTemplateName = "../go-easyConfigDemo/easyConfigDemo.template.yaml"
		configDataName     = "../go-easyConfigDemo/easyConfigDemo.data.encrypted.yaml"
	)

	var (
		demoCfg demo.DemoConfig
	)

	var (
		decryptionKeyPath = flag.String("k", "", "Decryption Key File")
	)
	flag.Parse()

	err := easyconfig.TemplateFromFile(
		configTemplateName,
		configDataName,
		&demoCfg)

	demo.CheckError(err)

	if *decryptionKeyPath == "" {
		fmt.Println("no encryption key found")
		os.Exit(-1)
	}

	err = easyconfig.DecryptFromFile(*decryptionKeyPath, &demoCfg)
	if err != nil {
		demo.CheckError(err)
	}

	fmt.Println(demoCfg.String())
}
