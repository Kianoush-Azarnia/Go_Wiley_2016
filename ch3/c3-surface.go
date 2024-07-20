// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320              // canvas size in pixels
	cells         = 40                    // number of grid cells
	xyrange       = 5.0                   // axis ranges (-xyrange..+xyrange)
	xyscale       = width / (2 * xyrange) // pixels per x or y unit
	zscale        = height * 0.4          // pixels per x unit
	angle         = math.Pi / 6           // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	server()
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		color, _ := getParam("color", r)
		height, _ := getParam("height", r)
		width, _ := getParam("width", r)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(surfaceSVG(color, height, width)))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getParam(param string, r *http.Request) (val int, err error) {
	theParam := r.URL.Query().Get(param)
	// Convert the parameter to an integer
	val, err = strconv.Atoi(theParam)
	if err != nil {
		// Default to 600 param if the parameter is missing or invalid
		return 600, err
	}
	return val, err
}

func surfaceSVG(color, height, width int) (res string) {
	res += fmt.Sprintf(
		"<svg xmlns='http://www.w3.org/2000/svg'"+
			"style='stroke: grey; fill: %d; stroke-width: 0.7' "+
			"width='%d' height='%d'>",
		color, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			res += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	res += fmt.Sprintf("</svg>")
	return res
}

func corner(i, j int) (sx, sy float64) {
	// find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, xy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	return (x*x + 3*y*y) * math.Exp(-x*x-y*y)
}
