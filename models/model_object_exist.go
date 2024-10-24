/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Object exists
type ObjectExist struct {
	// Indicates that the file or folder exists.
	Exists bool `json:"Exists"`
	// True if it is a folder, false if it is a file.
	IsFolder bool `json:"IsFolder"`
}
