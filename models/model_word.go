/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents relatively positioned rectangle which contains single word
type Word struct {
	// The X coordinate of the highest left point on the page layout where the rectangle that contains element begins.
	X float64 `json:"X"`
	// The Y coordinate of the highest left point on the page layout where the rectangle that contains element begins.
	Y float64 `json:"Y"`
	// The width of the rectangle which contains the element (in pixels).
	Width float64 `json:"Width"`
	// The height of the rectangle which contains the element (in pixels).
	Height float64 `json:"Height"`
	// The element value
	Value string `json:"Value,omitempty"`
	// The characters contained by the word
	Characters []Character `json:"Characters,omitempty"`
}
