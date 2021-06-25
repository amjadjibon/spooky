package spooky

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func GetUUID(c *cli.Context) error {
	id := uuid.New()
	fmt.Println(id.String())
	return nil
}
