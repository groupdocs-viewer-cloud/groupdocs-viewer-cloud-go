/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// The text overflow mode for rendering spreadsheet documents into HTML
type TextOverflowModeEnum string

const (
	TextOverflowModeOverlay              TextOverflowModeEnum = "Overlay"
	TextOverflowModeOverlayIfNextIsEmpty TextOverflowModeEnum = "OverlayIfNextIsEmpty"
	TextOverflowModeAutoFitColumn        TextOverflowModeEnum = "AutoFitColumn"
	TextOverflowModeHideText             TextOverflowModeEnum = "HideText"
)
