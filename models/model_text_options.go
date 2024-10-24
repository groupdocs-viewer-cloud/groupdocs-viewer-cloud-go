/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering text documents
type TextOptions struct {
	// Max chars per row on page. Default value is 85.
	MaxCharsPerRow int32 `json:"MaxCharsPerRow,omitempty"`
	// Max rows per page. Default value is 55.
	MaxRowsPerPage int32 `json:"MaxRowsPerPage,omitempty"`
}
