/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Page information
type PageView struct {
	// Path of resource file in storage
	Path string `json:"Path,omitempty"`
	// ULR to retrieve resource.
	DownloadUrl string `json:"DownloadUrl,omitempty"`
	// Page number
	Number int32 `json:"Number"`
	// HTML resources.
	Resources []HtmlResource `json:"Resources,omitempty"`
}
