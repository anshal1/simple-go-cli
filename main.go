package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	greet "example.com/mycli-01/commands/greet"
	"example.com/mycli-01/commands/help"
	"example.com/mycli-01/utils"
)

var validCommands = map[string]utils.CommandInfo{
	"greet": {
		CommandInfo: "Greets the user by name passed to this command",
		// nil means that the command can take any number of argument
		ValidArgsforCommand: nil,
		// nil means that this command does not take any flags
		Flags: []string{"-r", "-c"},
	},
	"help": {
		CommandInfo:         "Shows all the details regarding this program",
		ValidArgsforCommand: nil,
		Flags:               nil,
	},
}

func main() {
	// os.Args contains the program name and the argument passed to the program
	// os.Argsp[0] contains the program name/program file name
	// os.Args[1] and so on will contain the args passed to the program
	args := parseCommands(os.Args)
	for key, _ := range args {
		if key == "command" {
			commadExist := validCommands[args[strings.ToLower(key)]]
			if commadExist.CommandInfo == "" {
				fmt.Println("Unkown command")
				os.Exit(1)
			}
		}
	}
	command := args["command"]
	flags := args["flags"]
	if command == "" {
		fmt.Println("No command provided")
		os.Exit(1)
	}
	switch command {
	case "greet":

		err := validateFlags(flags, command)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if flags == "" {
			greet.Greet(args)
		} else {
			greet.GreetWithFlags(args)
		}

	case "help":
		err := validateFlags(flags, command)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		help.Help(validCommands)
	}
}

func parseCommands(args []string) map[string]string {
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("No args provided")
		os.Exit(1)
	}
	var argMap = make(map[string]string)
	for i, val := range args {
		switch i {
		case 0:
			continue
		case 1:
			// the name of the command
			if val == "" {
				fmt.Println("Please provide a valid command")
				os.Exit(1)
			}
			argMap["command"] = strings.ToLower(val)

		case 2:
			// if val is not empty and starts with a "-" then treat it as a flag
			if val != "" && val[0] == '-' {
				argMap["flags"] = strings.ToLower(val)
				// if val exits but does not starts with a "-" treat it as a value of a command
			} else if val != "" {
				argMap["value-1"] = val
			}
			// all the remaining args will be treated as values of the respective command
		default:
			argMap["value-"+strconv.Itoa(i-1)] = val

		}

	}
	return argMap
}

func validateFlags(flags string, commandName string) error {
	if flags == "" {
		return nil
	}
	commandData := validCommands[commandName]
	if commandData.CommandInfo == "" {
		return errors.New("Unsupported command was passed to the flag parser")
	}
	flagsRequired := commandData.Flags != nil
	if !flagsRequired && flags != "" {
		return errors.New("Flags are not required for this command")
	}
	for i := range flags {
		if flags[i] == '-' {
			continue
		}
		var flag strings.Builder
		flag.WriteRune('-')
		flag.WriteByte(flags[i])
		if !slices.Contains(commandData.Flags, flag.String()) {
			return errors.New("Invalid flag provided " + flag.String())
		}
	}
	return nil
}
