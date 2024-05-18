package mytheme

import (
	_ "embed"

	"image/color"

	"fyne.io/fyne/v2"
	"github.com/hlinfocc/gouidemo/assets"

	"fyne.io/fyne/v2/theme"
)

type theme1 struct{}

var _ fyne.Theme = (*theme1)(nil)

func (*theme1) Font(s fyne.TextStyle) fyne.Resource {
	return assets.ResourceHarmonyOSSansSCLightTtf
}

func (*theme1) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	// log.Println("ThemeVariant:", v)
	return theme.DefaultTheme().Color(n, 1)
}

func (*theme1) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*theme1) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func NewTheme() fyne.Theme {
	return &theme1{}
}
