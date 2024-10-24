/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents relatively positioned rectangle which contains single line
type Line struct {
	// The X coordinate of the highest left point on the page layout where the rectangle that contains element begins.
	X float64 `json:"X,omitempty"`
	// The Y coordinate of the highest left point on the page layout where the rectangle that contains element begins.
	Y float64 `json:"Y,omitempty"`
	// The width of the rectangle which contains the element (in pixels).
	Width float64 `json:"Width,omitempty"`
	// The height of the rectangle which contains the element (in pixels).
	Height float64 `json:"Height,omitempty"`
	// The element value
	Value string `json:"Value,omitempty"`
	// The words contained by the line
	Words []Word `json:"Words,omitempty"`
}