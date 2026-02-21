package main

import ("strings"
		"os"
		"bufio"
		"fmt")

type cliCommand struct {
	name        string
	description string
	callback    func() error

}

func getCliCommands()map[string]cliCommand{
	var cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit", 
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help", 
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return cliCommands
}



func cleanInput(text string)[]string{
	str_array := strings.Fields(text)
	for idx,val := range str_array{
		str_array[idx]=strings.ToLower(val)
	}
	return str_array
}

func startRepl(){
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if has_value := scanner.Scan();has_value{
			cleanCommand := cleanInput(scanner.Text())
			if val,bool:=cliCommands[cleanCommand[0]]; bool==true{
				val.callback()
			}else{
				fmt.Print("Unknown command")
			}
		}
		fmt.Print("\n")
	}
}

func commandExit() error{
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	cliCommands := getCliCommands()
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	for _,value := range cliCommands{
		fmt.Printf("\n%s: %s",value.name,value.description)
	}
	return nil
}
