package main

import (
	"fmt"

	"github.com/justym/deadcode-sample/module3"
)

func main() {
	apples := module3.Apples{&module3.Apple{module3.AppleColorRed}, &module3.Apple{module3.AppleColorGreen}, &module3.Apple{module3.AppleColorRed}}
	oranges := module3.Oranges{&module3.Orange{module3.OrangeColorRed}, &module3.Orange{module3.OrangeColorOrange}, &module3.Orange{module3.OrangeColorRed}}

	fmt.Println(apples, oranges)
}
