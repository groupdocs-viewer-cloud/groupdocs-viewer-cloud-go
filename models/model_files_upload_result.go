/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// File upload result
type FilesUploadResult struct {
	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`
	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}
