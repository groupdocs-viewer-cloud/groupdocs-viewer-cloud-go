/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// The Visio files processing documents view options.
type VisioRenderingOptions struct {
	// Render only Visio figures, not a diagram.
	RenderFiguresOnly bool `json:"RenderFiguresOnly,omitempty"`
	// Figure width, height will be calculated automatically. Default value is 100.
	FigureWidth int32 `json:"FigureWidth,omitempty"`
}
