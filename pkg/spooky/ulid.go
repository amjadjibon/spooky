package spooky

import (
	"fmt"
	"github.com/oklog/ulid"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

func GetULID(c *cli.Context) error {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Println(id.String())
	return nil
}