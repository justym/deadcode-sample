package module3

type Apple struct {
	Color AppleColor
}

type AppleColor string

const (
	AppleColorRed   AppleColor = "red"
	AppleColorGreen AppleColor = "green"
)

func (a *Apple) EqualColor(appleColor AppleColor) bool {
	return a.Color == appleColor
}

type Apples []*Apple

func (as Apples) FilterByAppleColor(appleColor AppleColor) Apples {
	filteredApples := make(Apples, 0, len(as))
	for _, apple := range as {
		if apple.EqualColor(appleColor) {
			filteredApples = append(filteredApples, apple)
		}
	}
	return filteredApples
}
