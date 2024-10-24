/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Page information
type PageInfo struct {
	// The page number
	Number int32 `json:"Number"`
	// The width of the page in pixels when viewing as JPG or PNG
	Width int32 `json:"Width"`
	// The height of the page in pixels when viewing as JPG or PNG
	Height int32 `json:"Height"`
	// The page visibility indicator
	Visible bool `json:"Visible"`
	// The lines contained by the page when viewing as JPG or PNG with enabled Text Extraction
	Lines []Line `json:"Lines,omitempty"`
}
