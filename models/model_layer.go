/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents layer contained by the CAD drawing
type Layer struct {
	// The name of the layer
	Name string `json:"Name,omitempty"`
	// The layer visibility indicator
	Visible bool `json:"Visible,omitempty"`
}
