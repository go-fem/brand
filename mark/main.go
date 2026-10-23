// Command mark writes the go-fem mark, svg/color/go-fem.svg:
//
//	cd mark && go run . > ../svg/color/go-fem.svg
//
// go-fem is finite elements: a domain cut into small triangles on which a PDE
// is solved. The mark is that mesh: a square of 3 x 3 nodes, every cell cut by
// the same diagonal into two triangles, with the nodes dotted.
//
// Nothing is placed by hand. The program lays the nodes on a regular grid,
// centred on the square, then draws the grid lines and one diagonal per cell.
// Change the grid size and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#B5D23F" // the gradient's light stop
	dark  = "#5C7A0A" // its dark stop

	stroke = 13.0 // the family's line width
	dotR   = 11.0 // radius of a node

	nodes = 3    // nodes per side
	pitch = 66.0 // distance between two neighbouring nodes
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	origin := 128 - float64(nodes-1)*pitch/2
	at := func(i float64) float64 { return origin + i*pitch }
	last := float64(nodes - 1)

	var lines []string
	// The outline and the interior grid lines.
	for i := 0.0; i <= last; i++ {
		lines = append(lines,
			fmt.Sprintf(`<path d="M%s %sH%s"/>`, num(at(0)), num(at(i)), num(at(last))),
			fmt.Sprintf(`<path d="M%s %sV%s"/>`, num(at(i)), num(at(0)), num(at(last))))
	}
	// One diagonal per cell, all the same way: each cell is two triangles.
	for i := 0.0; i < last; i++ {
		for j := 0.0; j < last; j++ {
			lines = append(lines, fmt.Sprintf(`<path d="M%s %sL%s %s"/>`,
				num(at(i)), num(at(j+1)), num(at(i+1)), num(at(j))))
		}
	}

	var dots []string
	for i := 0.0; i <= last; i++ {
		for j := 0.0; j <= last; j++ {
			dots = append(dots, fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(at(i)), num(at(j)), num(dotR)))
		}
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-fem">
  <defs><linearGradient id="femgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#femgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
  <g fill="#FFFFFF">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(lines, "\n    "), strings.Join(dots, "\n    "))
}
