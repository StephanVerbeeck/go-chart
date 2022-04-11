package chart

import (
	"fmt"
	"math"
	"strings"
)

// TicksProvider is a type that provides ticks.
type TicksProvider interface {
	GetTicks(r Renderer, defaults Style, vf ValueFormatter) []Tick
}

// Tick represents a label on an axis.
type Tick struct {
	Value float64
	Label string
}

// Ticks is an array of ticks.
type Ticks []Tick

// Len returns the length of the ticks set.
func (t Ticks) Len() int {
	return len(t)
}

// Swap swaps two elements.
func (t Ticks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Less returns if i's value is less than j's value.
func (t Ticks) Less(i, j int) bool {
	return t[i].Value < t[j].Value
}

// String returns a string representation of the set of ticks.
func (t Ticks) String() string {
	var values []string
	for i, tick := range t {
		values = append(values, fmt.Sprintf("[%d: %s]", i, tick.Label))
	}
	return strings.Join(values, ", ")
}

// GenerateContinuousTicks generates a set of ticks.
func GenerateContinuousTicks(r Renderer, ra Range, isVertical bool, style Style, vf ValueFormatter) []Tick {
	if vf == nil {
		vf = FloatValueFormatter
	}

	var ticks []Tick
	min, max := ra.GetMin(), ra.GetMax()

	minLabel := vf(min)
	style.GetTextOptions().WriteToRenderer(r)
	labelBox := r.MeasureText(minLabel)

	var tickSize float64
	if isVertical {
		tickSize = float64(labelBox.Height() + DefaultMinimumTickVerticalSpacing)
	} else {
		tickSize = float64(labelBox.Width() + DefaultMinimumTickHorizontalSpacing)
	}

	valueInterval := ra.GetInterval()                          // suggested value step per drawn tick
	valueRange := math.Abs(max - min)                          // difference in value of the first and last tick
	visualRange := float64(ra.GetDomain())                     // visual size available for drawing
	tickCount := int(float64(visualRange) / float64(tickSize)) // max number of ticks that fit visually
	valueStep := valueRange / float64(tickCount)

	if valueInterval != 0 { // is a value interval per tick suggested for this graph?
		for _, intervalsPerTick := range []float64{1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 500000, 1000000} {
			tickCountWithInterval := int(valueRange / (valueInterval * intervalsPerTick))
			if tickCountWithInterval <= tickCount { // is there enough room to draw this amount of ticks
				tickCount = tickCountWithInterval            // set tickCount that complies with suggested value interval per tick
				valueStep = valueInterval * intervalsPerTick // adapt values change per tick for the new tick count
				break
			}
		}
	}

	if tickCount > DefaultTickCountSanityCheck {
		tickCount = DefaultTickCountSanityCheck
	} else if tickCount < 1 {
		tickCount = 1
		valueStep = valueRange
	}

	// add ticks
	for x := 0; x <= tickCount; x++ {
		var tickValue float64
		if x == 0 || x == tickCount-1 || valueInterval != 0 {
			if ra.IsDescending() {
				tickValue = max - valueStep*float64(x)
			} else {
				tickValue = min + valueStep*float64(x)
			}
		} else { // round tick value
			roundTo := GetRoundToForDelta(valueRange) / 10
			if ra.IsDescending() {
				tickValue = max - RoundUp(valueStep*float64(x), roundTo)
			} else {
				tickValue = min + RoundUp(valueStep*float64(x), roundTo)
			}
		}
		ticks = append(ticks, Tick{
			Value: tickValue,
			Label: vf(tickValue),
		})
	}

	return ticks
}
