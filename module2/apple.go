package module2

type Apple struct {
	Color AppleColor
}

type AppleColor string

const (
	AppleColorRed   AppleColor = "red"
	AppleColorGreen AppleColor = "green"
)

func NewApple(color AppleColor) *Apple {
	return &Apple{
		Color: color,
	}
}

func (a *Apple) EqualColor(appleColor AppleColor) bool {
	return a.Color == appleColor
}

type Apples []*Apple

func NewApples(colors ...AppleColor) Apples {
	apples := make(Apples, 0, len(colors))
	for _, color := range colors {
		apples = append(apples, NewApple(color))
	}
	return apples
}

func (as Apples) FilterByAppleColor(appleColor AppleColor) Apples {
	filteredApples := make(Apples, 0, len(as))
	for _, apple := range as {
		if apple.EqualColor(appleColor) {
			filteredApples = append(filteredApples, apple)
		}
	}
	return filteredApples
}
