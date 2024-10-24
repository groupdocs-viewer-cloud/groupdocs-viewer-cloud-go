/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering Outlook data files
type OutlookOptions struct {
	// The name of the folder (e.g. Inbox, Sent Item or Deleted Items) to render
	Folder string `json:"Folder,omitempty"`
	// The keywords used to filter messages
	TextFilter string `json:"TextFilter,omitempty"`
	// The email-address used to filter messages by sender or recipient
	AddressFilter string `json:"AddressFilter,omitempty"`
	// The maximum number of messages or items, that can be rendered from one folder
	MaxItemsInFolder int32 `json:"MaxItemsInFolder,omitempty"`
}
