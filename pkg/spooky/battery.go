package spooky

import (
	"fmt"
	"github.com/distatus/battery"
	"github.com/urfave/cli/v2"
)

func GetBatteryInfo(c *cli.Context) error {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return err
	}
	for i, battery := range batteries {
		fmt.Printf("Bat%d: ", i+1)
		fmt.Printf("state: %s, ", battery.State.String())
		fmt.Printf("current capacity: %f mWh, ", battery.Current)
		fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		fmt.Printf("design capacity: %f mWh, ", battery.Design)
		fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
		fmt.Printf("voltage: %f V, ", battery.Voltage)
		fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)
	}
	return nil
}
