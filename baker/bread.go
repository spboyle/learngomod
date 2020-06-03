package baker

import (
	"fmt"
)

func MakeBread(i string, j string) {
	fmt.Printf("%s + bread + %s\n", i, j)
}

func BakeBread(ingredients ...string) {

	for i, ingredient := range ingredients {
		fmt.Printf("Step %d. Adding %s\n", i, ingredient)
	}

	fmt.Println("Mixing all together...")
	fmt.Println("Baking...")
	fmt.Println("Done!")
}
