package main

import (
	"github.com/justym/deadcode-sample/module1"
)

func main() {
	_ = module1.NewApples(module1.AppleColorRed, module1.AppleColorGreen, module1.AppleColorRed)
	_ = module1.NewOranges(module1.OrangeColorRed, module1.OrangeColorOrange, module1.OrangeColorRed)
}
