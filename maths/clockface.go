package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

// Point represents a two-dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// SecondHand returns the endpoint of the second hand for time t.
func SecondHand(t time.Time) Point {
	return makeHand(secondHandPoint(t), secondHandLength)
}

// MinuteHand returns the endpoint of the minute hand for time t.
func MinuteHand(t time.Time) Point {
	return makeHand(minuteHandPoint(t), minuteHandLength)
}

// HourHand returns the endpoint of the hour hand for time t.
func HourHand(t time.Time) Point {
	return makeHand(hourHandPoint(t), hourHandLength)
}

func secondsInRadians(t time.Time) float64 {
	second := float64(t.Second())
	if second == 0 {
		return 0
	}
	return math.Pi / (secondsInHalfClock / second)
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

// SVGWriter writes an SVG representation of an analogue clock for time t.
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := SecondHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := MinuteHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := HourHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
