/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering word processing documents
type WordProcessingOptions struct {
	// Enables tracked changes (revisions) rendering
	RenderTrackedChanges bool `json:"RenderTrackedChanges,omitempty"`
	// Left page margin (for HTML rendering only)
	LeftMargin float64 `json:"LeftMargin,omitempty"`
	// Right page margin (for HTML rendering only)
	RightMargin float64 `json:"RightMargin,omitempty"`
	// Top page margin (for HTML rendering only)
	TopMargin float64 `json:"TopMargin,omitempty"`
	// Bottom page margin (for HTML rendering only)
	BottomMargin float64 `json:"BottomMargin,omitempty"`
}
