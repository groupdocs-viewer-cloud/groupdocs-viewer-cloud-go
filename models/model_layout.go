/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents layout contained by the CAD drawing
type Layout struct {
	// The name of the layout
	Name string `json:"Name,omitempty"`
	// The width of the layout
	Width float64 `json:"Width,omitempty"`
	// The height of the layout
	Height float64 `json:"Height,omitempty"`
}
