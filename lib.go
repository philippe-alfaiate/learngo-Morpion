package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	//"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const ssize = 900

const s900 = ssize
const s300 = s900 / 3
const s600 = s300 * 2
const s150 = s300 / 2
const s100 = s900 / 9
const s250 = s150 + s100
const s450 = s300 + s150
const s50 = s100 / 2
const s10 = s100 / 10

func verification(plateau [3][3]string, joueur string) bool {

	for i := 0; i < 3; i++ {
		if plateau[i][0] == joueur && plateau[i][1] == joueur && plateau[i][2] == joueur {

			return true
		}
	}
	for i := 0; i < 3; i++ {
		if plateau[0][i] == joueur && plateau[1][i] == joueur && plateau[2][i] == joueur {

			return true
		}
	}
	if plateau[0][0] == joueur && plateau[1][1] == joueur && plateau[2][2] == joueur {

		return true
	}
	if plateau[0][2] == joueur && plateau[1][1] == joueur && plateau[2][0] == joueur {

		return true
	}

	return false
}

func fin(win *pixelgl.Window, gagant string) bool {
	var plateau [3][3]string

	if win.Closed() {
		nonQuitté = false
	}

	if gagant == "pas encore" {

		return false
	}
	win.Clear(color.RGBA{255, 255, 255, 255})
	drawPlateau(win, plateau)

	for colonne, colonnes := range plateau {
		for ligne := range colonnes {

			x := float64(colonne) * s300
			y := float64(ligne) * s300

			if gagant == "O" {
				drawRond(win, pixel.V(x, y))
			} else if gagant == "X" {
				drawCroix(win, pixel.V(x, y))
			} else {
				drawRond(win, pixel.V(x, y))
				drawCroix(win, pixel.V(x, y))
			}

		}
		win.Update()
		time.Sleep(time.Millisecond * 500)
	}

	win.Update()
	time.Sleep(time.Millisecond * 500)
	return true
}

func menu2(win *pixelgl.Window) string {

	notChoose := true
	iaValeur := "O"
	for notChoose {
		win.Clear(color.RGBA{255, 255, 255, 255})
		drawCroix(win, pixel.V(s150, s300))
		drawRond(win, pixel.V(s450, s300))
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if win.MousePosition().X > s450 {
				iaValeur = "X"
			}
			notChoose = false
		}
		win.Update()
		if win.Closed() {
			notChoose = false
		}
	}

	return iaValeur
}

func menu1(win *pixelgl.Window) bool {

	notChoose := true
	val := true

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txtIA := text.New(pixel.V(s250, s300), basicAtlas)
	txtJoueur := text.New(pixel.V(s250, s600), basicAtlas)

	txtIA.Color = colornames.Red
	txtJoueur.Color = colornames.Red

	fmt.Fprintln(txtIA, "joueur vs machine")
	fmt.Fprintln(txtJoueur, "joueur vs joueur")

	for notChoose {
		win.Clear(color.RGBA{255, 255, 255, 255})

		txtJoueur.Draw(win, pixel.IM.Scaled(txtJoueur.Orig, 4))
		txtIA.Draw(win, pixel.IM.Scaled(txtIA.Orig, 4))
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if win.MousePosition().Y > s450 {
				val = false
			} else {
				val = true
			}
			notChoose = false
		}
		win.Update()
		if win.Closed() {
			notChoose = false
		}
	}

	return val
}

func menu(win *pixelgl.Window) string {

	iaValeur := ""

	if menu1(win) {
		iaValeur = menu2(win)
	}
	return iaValeur
}

func testcroix(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = color.RGBA{100, 100, 100, 255}

	imd.Push(pixel.V(s300, 0), pixel.V(s300, s900))
	imd.Line(s10)

	imd.Push(pixel.V(s600, 0), pixel.V(s600, s900))
	imd.Line(s10)

	imd.Push(pixel.V(0, s300), pixel.V(s900, s300))
	imd.Line(s10)

	imd.Push(pixel.V(0, s600), pixel.V(s900, s600))
	imd.Line(s10)

	imd.Draw(win)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			drawCroix(win, pixel.V(float64(i)*s300, float64(j)*s300))
		}
	}

}

func testrond(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = color.RGBA{100, 100, 100, 255}

	imd.Push(pixel.V(s300, 0), pixel.V(s300, s900))
	imd.Line(s10)

	imd.Push(pixel.V(s600, 0), pixel.V(s600, s900))
	imd.Line(s10)

	imd.Push(pixel.V(0, s300), pixel.V(s900, s300))
	imd.Line(s10)

	imd.Push(pixel.V(0, s600), pixel.V(s900, s600))
	imd.Line(s10)

	imd.Draw(win)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			drawRond(win, pixel.V(float64(i)*s300, float64(j)*s300))
		}
	}

}

func testCroixRond(win *pixelgl.Window) {
	win.Clear(color.RGBA{250, 250, 250, 255})
	testcroix(win)
	testrond(win)
	win.Update()
}

func testplateau(win *pixelgl.Window) {

	var plateau [3][3]string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if r.Intn(2) == 1 {
				plateau[i][j] = "X"
			} else {
				plateau[i][j] = "O"
			}
		}
	}
	drawPlateau(win, plateau)
	time.Sleep(time.Millisecond * 300)
}

func clic(plateau [3][3]string, position pixel.Vec) (bool, int, int) {
	var x, y int

	if position.X < s300 {
		x = 0
	} else if position.X < s600 {
		x = 1
	} else {
		x = 2
	}

	if position.Y < s300 {
		y = 0
	} else if position.Y < s600 {
		y = 1
	} else {
		y = 2
	}
	if plateau[x][y] != "" {
		return false, x, y
	}
	return true, x, y
}

func joueurIA(win *pixelgl.Window, plateau [3][3]string) (bool, int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	notFind := true
	var x, y int
	for notFind {
		x = r.Intn(3)
		y = r.Intn(3)
		if plateau[x][y] == "" {
			notFind = false
			return true, x, y
		}
	}
	return true, 0, 0
}

func joueur(win *pixelgl.Window, plateau [3][3]string) (bool, int, int) {

	action := win.JustPressed(pixelgl.MouseButtonLeft)
	clicValide := false
	x, y := 0, 0
	if action {
		clicValide, x, y = clic(plateau, win.MousePosition())
	}
	return clicValide, x, y
}

func changeJoueur(joueur string) string {
	if joueur == "O" {
		joueur = "X"
	} else {
		joueur = "O"
	}
	return joueur
}

func alors(win *pixelgl.Window, plateau [3][3]string, joueurValeur string, tour int) string {
	cFini := "pas encore"
	if verification(plateau, joueurValeur) {
		cFini = joueurValeur
		win.Clear(color.RGBA{250, 250, 250, 255})
		drawPlateau(win, plateau)
		win.Update()
		time.Sleep(time.Millisecond * 500)
	} else if tour == 9 {
		cFini = "égalité"
	}

	return cFini
}

func auProchain(win *pixelgl.Window, plateau [3][3]string, joueurValeur string, iaValeur string) (bool, int, int) {
	var valide bool
	var x, y int
	ia := true
	if iaValeur == "" {
		ia = false
	}
	if ia && iaValeur == joueurValeur {
		valide, x, y = joueurIA(win, plateau)
	} else {
		valide, x, y = joueur(win, plateau)
	}
	return valide, x, y
}
