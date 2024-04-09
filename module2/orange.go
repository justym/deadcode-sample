package module2

type Orange struct {
	Color OrangeColor
}

type OrangeColor string

const (
	OrangeColorRed    OrangeColor = "red"
	OrangeColorOrange OrangeColor = "orange"
)

func NewOrange(color OrangeColor) *Orange {
	return &Orange{
		Color: color,
	}
}

func (a *Orange) EqualColor(OrangeColor OrangeColor) bool {
	return a.Color == OrangeColor
}

type Oranges []*Orange

func NewOranges(colors ...OrangeColor) Oranges {
	oranges := make(Oranges, 0, len(colors))
	for _, color := range colors {
		oranges = append(oranges, NewOrange(color))
	}
	return oranges
}

func (os Oranges) FindByOrangeColor(orangeColor OrangeColor) *Orange {
	for _, orange := range os {
		if orange.EqualColor(orangeColor) {
			return orange
		}
	}
	return nil
}
