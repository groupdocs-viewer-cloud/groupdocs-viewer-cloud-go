/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Specifies output image quality for image resources when rendering into HTML. The default value is Low
type ImageQualityEnum string

const (
	ImageQualityLow    ImageQualityEnum = "Low"
	ImageQualityMedium ImageQualityEnum = "Medium"
	ImageQualityHigh   ImageQualityEnum = "High"
)
