package mainz

import "fmt"

func OnExit() {
	err := main.onExit()
	if err != nil {
		_, err := fmt.Printf("Problem running onExit: %e\n", err)
		if err != nil {
			fmt.Println("Big issues!")
			return
		}
		return
	}
}
