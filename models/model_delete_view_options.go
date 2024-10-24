/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Delete view options
type DeleteViewOptions struct {
	// File info
	FileInfo *FileInfo `json:"FileInfo,omitempty"`
	// The output path Default value is 'viewer\\{input file path}_{file extension}\\'
	OutputPath string `json:"OutputPath,omitempty"`
}
