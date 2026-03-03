package help

import (
	"fmt"

	"example.com/mycli-01/utils"
)

func Help(commands map[string]utils.CommandInfo) {
	for key, value := range commands {
		fmt.Println("---------------------------------------------------------------------")
		fmt.Println("Command:", key)
		fmt.Println("Description:", value.CommandInfo)
		fmt.Println("Valid values for command:", value.ValidArgsforCommand)
		fmt.Println("Flags:", value.Flags)

	}
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("NOTE: Flags can be grouped together like this -xyz")
	fmt.Println("Thank You!")
}
