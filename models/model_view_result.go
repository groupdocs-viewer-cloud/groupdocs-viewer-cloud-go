/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// View result information
type ViewResult struct {
	// View result pages
	Pages []PageView `json:"Pages,omitempty"`
	// Attachments
	Attachments []AttachmentView `json:"Attachments,omitempty"`
	// ULR to retrieve file.
	File *Resource `json:"File,omitempty"`
}
