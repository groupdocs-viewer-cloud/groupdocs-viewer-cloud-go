/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering archive files
type ArchiveOptions struct {
	// The folder inside the archive to be rendered
	Folder string `json:"Folder,omitempty"`
	// The filename to display in the header. By default the name of the source file is displayed.
	FileName string `json:"FileName,omitempty"`
	// Number of records per page (for rendering to HTML only)
	ItemsPerPage int32 `json:"ItemsPerPage,omitempty"`
}
