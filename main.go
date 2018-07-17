package main

import (
	"os"
	"strconv"

	"github.com/ajstarks/svgo"
)

// func(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
//   }

func main() {

	DrawCycleNormalTBLRSeat(1, 1, 1, 1, 0, 0)
	DrawCycleNormalTBLRSeat(1, 1, 0, 0, 0, 0)

	DrawCycleNormalTBLRSeat(2, 2, 2, 2, 0, 0)
	DrawCycleNormalTBLRSeat(2, 1, 2, 2, 0, 0)
	DrawCycleNormalTBLRSeat(2, 1, 0, 0, 0, 0)

	DrawCycleNormalTBLRSeat(4, 4, 4, 4, 4, 4)
	DrawCycleNormalTBLRSeat(3, 1, 3, 3, 0, 0)

}

//DrawCycleNormalTBLRSeat -
func DrawCycleNormalTBLRSeat(sizeXSeat int,
	sizeYSeat int,
	displaySeatTop int,
	displaySeatBottom int,
	displaySeatLeft int,
	displaySeatRight int,
) {

	width := (sizeXSeat * 45)
	height := (sizeYSeat * 45)

	if sizeXSeat == 1 {
		width = (sizeXSeat * 45) + 45
	}

	if sizeYSeat == 1 {
		height = (sizeYSeat * 45) + 45
	}

	if sizeXSeat == 2 {
		width = (sizeXSeat * 45) + 35
	}

	if sizeYSeat == 2 {
		height = (sizeYSeat * 45) + 35
	}

	if sizeXSeat == 3 {
		width = (sizeXSeat * 45) + 25
	}

	if sizeYSeat == 3 {
		height = (sizeYSeat * 45) + 25
	}

	image := "DrawTableSquarSeat" + strconv.Itoa(sizeXSeat) + "_" + strconv.Itoa(sizeYSeat) + ".svg"
	var _, err = os.Stat(image)
	if os.IsExist(err) {
		os.Remove(image)
	}
	var file, _ = os.Create(image)
	canvas := svg.New(file)
	canvas.Start(width, height)

	DrawCycleNormalSeatSide(canvas, "top", displaySeatTop, width, height)
	DrawCycleNormalSeatSide(canvas, "bottom", displaySeatBottom, width, height)
	DrawCycleNormalSeatSide(canvas, "left", displaySeatLeft, width, height)
	DrawCycleNormalSeatSide(canvas, "right", displaySeatRight, width, height)

	canvas.Rect(20, 20, width-40, height-40, "fill:#ffffff;stroke:grey;stroke-dasharray:null;stroke-linejoin:null;stroke-linecap:null;rx:5;ry:5;stroke-width:5")
	canvas.End()
	defer file.Close()

}

//DrawCycleNormalSeatSide -
func DrawCycleNormalSeatSide(canvas *svg.SVG, side string, seat, width, height int) {
	beginX := 0
	beginY := 0
	beginL := 0
	beginR := 0

	switch side {
	case "top":
		beginX = width / 2
		beginY = 20
		beginL = (width / 2) - 15
		beginR = (width / 2) + 15
	case "bottom":
		beginX = width / 2
		beginY = height - 20
		beginL = (width / 2) - 15
		beginR = (width / 2) + 15
	case "left":
		beginX = 20
		beginY = height / 2
		beginL = (height / 2) - 15
		beginR = (height / 2) + 15
	case "right":
		beginX = width - 20
		beginY = height / 2
		beginL = (height / 2) - 15
		beginR = (height / 2) + 15
	}

	for seat > 0 {
		if seat%2 == 1 {
			DrawCycleNormalSeat(canvas, beginX, beginY, 10)
			seat--
			beginL -= 15
			beginR += 15
			continue
		}
		if side == "left" || side == "right" {
			DrawCycleNormalSeat(canvas, beginX, beginL, 10)
			DrawCycleNormalSeat(canvas, beginX, beginR, 10)
		} else {
			DrawCycleNormalSeat(canvas, beginL, beginY, 10)
			DrawCycleNormalSeat(canvas, beginR, beginY, 10)
		}
		seat -= 2
		beginL -= 30
		beginR += 30
	}
}

//DrawCycleNormalSeat -
func DrawCycleNormalSeat(canvas *svg.SVG, x, y, r int) {
	canvas.Circle(x, y, r, "stroke-linecap:null;stroke-linejoin:null;stroke-dasharray:null;stroke-width:null;stroke:#000000;fill:#191919")
}
