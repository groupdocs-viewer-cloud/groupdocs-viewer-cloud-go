/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Error
type ModelError struct {
	// Code
	Code string `json:"Code,omitempty"`
	// Message
	Message string `json:"Message,omitempty"`
	// Description
	Description string `json:"Description,omitempty"`
	// Inner Error
	InnerError *ErrorDetails `json:"InnerError,omitempty"`
}
