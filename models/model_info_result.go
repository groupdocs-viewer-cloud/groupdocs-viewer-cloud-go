/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// View result information
type InfoResult struct {
	// File format extension
	FormatExtension string `json:"FormatExtension,omitempty"`
	// File format
	Format string `json:"Format,omitempty"`
	// View result pages
	Pages []PageInfo `json:"Pages,omitempty"`
	// Attachments
	Attachments []AttachmentInfo `json:"Attachments,omitempty"`
	// Represents view information for archive file
	ArchiveViewInfo *ArchiveViewInfo `json:"ArchiveViewInfo,omitempty"`
	// Represents view information for CAD drawing
	CadViewInfo *CadViewInfo `json:"CadViewInfo,omitempty"`
	// Represents view information for MS Project document
	ProjectManagementViewInfo *ProjectManagementViewInfo `json:"ProjectManagementViewInfo,omitempty"`
	// Represents view information for Outlook Data file
	OutlookViewInfo *OutlookViewInfo `json:"OutlookViewInfo,omitempty"`
	// Represents view information for PDF document
	PdfViewInfo *PdfViewInfo `json:"PdfViewInfo,omitempty"`
}
