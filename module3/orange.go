package module3

type Orange struct {
	Color OrangeColor
}

type OrangeColor string

const (
	OrangeColorRed    OrangeColor = "red"
	OrangeColorOrange OrangeColor = "orange"
)

func (a *Orange) EqualColor(OrangeColor OrangeColor) bool {
	return a.Color == OrangeColor
}

type Oranges []*Orange

func (os Oranges) FindByOrangeColor(orangeColor OrangeColor) *Orange {
	for _, orange := range os {
		if orange.EqualColor(orangeColor) {
			return orange
		}
	}
	return nil
}
