/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering Mail storage (Lotus Notes, MBox) data files.
type MailStorageOptions struct {
	// The keywords used to filter messages.
	TextFilter string `json:"TextFilter,omitempty"`
	// The email-address used to filter messages by sender or recipient.
	AddressFilter string `json:"AddressFilter,omitempty"`
	// The maximum number of messages or items for render. Default value is 0 - all messages will be rendered
	MaxItems int32 `json:"MaxItems,omitempty"`
}
