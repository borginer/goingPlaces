// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	defWidth, defHeight = 600, 320    // canvas size in pixels
	cells               = 100         // number of grid cells
	xyrange             = 30.0        // axis ranges (-xyrange..+xyrange)
	angle               = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type SVGData struct {
	width   int
	height  int
	xyscale float64
	zscale  float64
}

var svgData SVGData

func SVGSet(prop string, val int) {
	if prop == "width" {
		svgData.width = val
	} else if prop == "height" {
		svgData.height = val
	} else {
		return
	}
	SVGSetScale()
}
func SVGSetScale() {
	svgData.xyscale = float64(svgData.width) / 2 / xyrange
	svgData.zscale = float64(svgData.height) * 0.4
}

func main() {
	http.HandleFunc("/", handler)

	SVGSet("height", defHeight)
	SVGSet("width", defWidth)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func ParseInt(s string) int {
	if len(s) > 0 {
		var err error
		width := int64(0)
		width, err = strconv.ParseInt(s, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing int: %v", err)
			return 0
		}
		return int(width)
	}
	return 0
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	width_str := query.Get("width")
	height_str := query.Get("height")
	color := query.Get("color")

	width := ParseInt(width_str)
	if width > 0 {
		SVGSet("width", width)
	}
	height := ParseInt(height_str)
	if height > 0 {
		SVGSet("height", height)
	}
	if len(color) == 0 {
		color = "white"
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, svgData.width, svgData.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(svgData.width)/2 + (x-y)*cos30*svgData.xyscale
	sy := float64(svgData.height)/2 + (x+y)*sin30*svgData.xyscale - z*svgData.zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
