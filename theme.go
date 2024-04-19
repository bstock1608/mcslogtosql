//go:generate fyne bundle -o bundled.go  res

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//Theme set up: https://www.youtube.com/watch?v=tYHD7OU9xfM

type mcsTheme struct {
	fyne.Theme
}

func newMcsTheme(mAppColors AppColors) fyne.Theme {
	return &mcsTheme{Theme: theme.DefaultTheme()}
}

func (m mcsTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	var mcsThemeColors = init_AppColors()
	if name == theme.ColorNameDisabled {
		if variant == theme.VariantLight {
			return mcsThemeColors.GREY_XLIGHT
		}
		return mcsThemeColors.GREY_XLIGHT
	}

	return theme.DefaultTheme().Color(name, variant)
}

//==================================================================
//==================================================================
//==================================================================

type AppColors struct {
	GREEN       color.RGBA
	GREY_LIGHT  color.RGBA
	GREY_XLIGHT color.RGBA
	GREY        color.RGBA
	PURPLE      color.RGBA
}

func init_AppColors() AppColors {
	var mAppColors AppColors = AppColors{
		GREEN:       color.RGBA{7, 187, 6, 255},
		GREY_XLIGHT: color.RGBA{161, 161, 161, 255},
		GREY_LIGHT:  color.RGBA{130, 130, 130, 255},
		GREY:        color.RGBA{96, 96, 96, 255},
		PURPLE:      color.RGBA{156, 2, 249, 255}}
	return mAppColors
}
