/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// File info
type FileInfo struct {
	// File path in storage
	FilePath string `json:"FilePath,omitempty"`
	// Storage name
	StorageName string `json:"StorageName,omitempty"`
	// Version ID
	VersionId string `json:"VersionId,omitempty"`
	// Password to open file
	Password string `json:"Password,omitempty"`
}
