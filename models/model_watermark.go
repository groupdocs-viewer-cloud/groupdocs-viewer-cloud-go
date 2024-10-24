/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Text watermark
type Watermark struct {
	// Watermark text.
	Text string `json:"Text,omitempty"`
	// Watermark color. Supported formats {Magenta|(112,222,11)|(50,112,222,11)}. Default value is \"Red\".
	Color    string                `json:"Color,omitempty"`
	Position WatermarkPositionEnum `json:"Position,omitempty"`
	// Watermark size in percents. Default value is 100.
	Size int32 `json:"Size,omitempty"`
}
