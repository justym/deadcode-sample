package main

import (
	"fmt"

	"github.com/justym/deadcode-sample/module2"
)

func main() {
	apples := module2.NewApples(module2.AppleColorRed, module2.AppleColorGreen, module2.AppleColorRed)
	oranges := module2.NewOranges(module2.OrangeColorRed, module2.OrangeColorOrange, module2.OrangeColorRed)

	fmt.Println(apples.FilterByAppleColor(module2.AppleColorGreen), oranges.FindByOrangeColor(module2.OrangeColorOrange))
}
