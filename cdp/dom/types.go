package dom

import "github.com/mailru/easyjson"

// Code generated by chromedp-gen. DO NOT EDIT.

// Quad an array of quad vertices, x immediately followed by y for each
// point, points clock-wise.
type Quad []float64

// BoxModel box model.
type BoxModel struct {
	Content      Quad              `json:"content"`                // Content box
	Padding      Quad              `json:"padding"`                // Padding box
	Border       Quad              `json:"border"`                 // Border box
	Margin       Quad              `json:"margin"`                 // Margin box
	Width        int64             `json:"width"`                  // Node width
	Height       int64             `json:"height"`                 // Node height
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// ShapeOutsideInfo cSS Shape Outside details.
type ShapeOutsideInfo struct {
	Bounds      Quad                  `json:"bounds"`      // Shape bounds
	Shape       []easyjson.RawMessage `json:"shape"`       // Shape coordinate details
	MarginShape []easyjson.RawMessage `json:"marginShape"` // Margin shape bounds
}

// Rect rectangle.
type Rect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}
