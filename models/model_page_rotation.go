/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Clockwise page rotation
type PageRotation struct {
	// Page number to rotate
	PageNumber    int32             `json:"PageNumber,omitempty"`
	RotationAngle RotationAngleEnum `json:"RotationAngle,omitempty"`
}
