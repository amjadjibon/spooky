package fakeit

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/urfave/cli/v2"
)

func FakeIt(input string) {
	faker := gofakeit.New(0)

	switch input {
	case "name":
		fmt.Println(faker.Name())
	case "person":
		jsonBytes, err := json.Marshal(faker.Person())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonBytes))
	case "email":
		fmt.Println(faker.Email())
	case "password":
		fmt.Printf("%v\n",faker.Password(true, true, true, true, false, 20))
	default:
		fmt.Println(faker.Word())
	}

}

func Fake(c *cli.Context) error {
	input := c.Args().First()

	FakeIt(input)

	return nil
}
