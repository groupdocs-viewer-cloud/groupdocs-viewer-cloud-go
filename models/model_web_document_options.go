/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// This rendering options enables you to customize the appearance of the output HTML/PDF/PNG/JPEG when rendering Web documents.
type WebDocumentOptions struct {
	PageSize PageSizeEnum `json:"PageSize,omitempty"`
	// The distance (in points) between the left edge of the page and the left boundary  of the body text. The default value is 5 points.
	LeftMargin float64 `json:"LeftMargin,omitempty"`
	// The distance (in points) between the right edge of the page and the right boundary of the body text. The default value is 5 points.
	RightMargin float64 `json:"RightMargin,omitempty"`
	// The distance (in points) between the top edge of the page and the top boundary of the body text. The default value is 72 points.
	TopMargin float64 `json:"TopMargin,omitempty"`
	// The distance (in points) between the bottom edge of the page and the bottom boundary of the body text. The default value is 72 points.
	BottomMargin float64 `json:"BottomMargin,omitempty"`
}
