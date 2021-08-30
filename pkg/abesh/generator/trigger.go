package generator

import (
	"github.com/amjadjibon/spooky/pkg/abesh/abeshtemplate"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

func GenerateTrigger(c *cli.Context) error {
	fileName := c.String("file")

	if !strings.HasSuffix(fileName, ".go") {
		fileName += ".go"
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {

		}
	}(f)

	structName := c.String("struct")
	var structRune string
	if len(structName) > 0 {
		structRune = string(structName[0])
		structRune = strings.ToLower(structRune)
	}


	err = abeshtemplate.TriggerTemplate.ExecuteTemplate(f, "trigger", struct {
		Timestamp     time.Time
		Author        string
		PackageName   string
		InterfaceName string
		InterfaceRune string
		ServiceName   string
		ContractID    string
	}{
		Timestamp:     time.Now(),
		Author:        "Amjad Jibon",
		PackageName:   c.String("package"),
		InterfaceName: structName,
		InterfaceRune: structRune,
		ServiceName:   c.String("service"),
		ContractID:    c.String("contract"),
	})
	if err != nil {
		return err
	}
	return nil
}
