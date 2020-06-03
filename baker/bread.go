package baker

import (
	"fmt"
	// "io/ioutil"
	"os"
	"path/filepath"
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

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	fmt.Printf("You are in %v\n", exPath)
	// bread, err := ioutil.ReadFile("baker/bread.txt")
	// if err != nil {
	//	panic(err)
	// }
	// fmt.Println(string(bread))

	fmt.Println("**** ENJOY ****")
}
