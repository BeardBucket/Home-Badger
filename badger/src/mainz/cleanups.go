package mainz

import "fmt"

func OnExit() {
	err := main.OnExit()
	if err != nil {
		_, err := fmt.Printf("Problem running OnExit: %e\n", err)
		if err != nil {
			fmt.Println("Big issues!")
			return
		}
		return
	}
}
