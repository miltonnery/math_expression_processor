package ascii_art

import "github.com/common-nighthawk/go-figure"

func WelcomeTitle(phrase, fontName string) {
	myFigure := figure.NewFigure(phrase, fontName, true)
	myFigure.Print()
}

func WelcomeMessage(phrase, fontName string) {
	slogan := figure.NewFigure(phrase, fontName, true)
	slogan.Print()
}
