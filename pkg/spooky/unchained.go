package spooky

import (
	"fmt"
	"github.com/alexandrevicenzi/unchained"
	"github.com/urfave/cli/v2"
)

func makeHash(pass string, algorithm string) (string, error) {
	hash, err := unchained.MakePassword(pass, unchained.GetRandomString(12), algorithm)
	return hash, err
}

func GenerateHashes(c *cli.Context) error {
	algorithm := c.String("algorithm")
	for _, pass := range c.Args().Slice() {
		hash, err := makeHash(pass, algorithm)
		if err != nil {
			return err
		}
		fmt.Printf("[%s] %s\n", pass, hash)
	}
	return nil
}

func CheckPassword(c *cli.Context) error {
	password := c.String("password")
	hash := c.String("hash")

	valid, err := unchained.CheckPassword(password, hash)

	if valid {
		fmt.Printf("[%s] [%s] [valid]\n", password, hash)
	} else {
		if err == nil {
			fmt.Printf("[%s] [%s] [invalid]\n", password, hash)
		} else {
			return err
		}
	}
	return nil
}
