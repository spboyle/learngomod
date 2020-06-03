package baker

import (
	"fmt"
	"io/ioutil"
)

func MakeBread(i string, j string) {
	fmt.Printf("%s + bread + %s\n", i, j)
}

func BakeBread(ingredients ...string) {

	for i, ingredient := range ingredients {
		fmt.Printf("Step %d. Adding %s\n", i+1, ingredient)
	}

	fmt.Println("Mixing all together...")
	fmt.Println("Baking...")
	fmt.Println("Done!")

	bread, err := ioutil.ReadFile("bread.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bread))

	fmt.Println("**** ENJOY ****")
}
