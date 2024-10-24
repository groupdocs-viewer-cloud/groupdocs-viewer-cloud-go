/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Reference to resource
type Resource struct {
	// Path of resource file in storage
	Path string `json:"Path,omitempty"`
	// ULR to retrieve resource.
	DownloadUrl string `json:"DownloadUrl,omitempty"`
}
