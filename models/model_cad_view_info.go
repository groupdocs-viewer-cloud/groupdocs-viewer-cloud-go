/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents view information for CAD drawing
type CadViewInfo struct {
	// The list of layers contained by the CAD drawing
	Layers []Layer `json:"Layers,omitempty"`
	// The list of layouts contained by the CAD drawing
	Layouts []Layout `json:"Layouts,omitempty"`
}
