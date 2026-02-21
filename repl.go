package main

import ("strings"
		"os"
		"bufio"
		"fmt")

func cleanInput(text string)[]string{
	str_array := strings.Fields(text)
	for idx,val := range str_array{
		str_array[idx]=strings.ToLower(val)
	}
	return str_array
}

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if has_data := scanner.Scan(); has_data==true{
			clean_text := cleanInput(scanner.Text())
			fmt.Printf("Your command was: %s\n",clean_text[0])
		}
	}
}
