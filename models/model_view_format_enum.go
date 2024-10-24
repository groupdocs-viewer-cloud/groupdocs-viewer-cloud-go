/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// View format (HTML, PNG, JPG, or PDF) Default value is HTML.
type ViewFormatEnum string

const (
	ViewFormatHtml ViewFormatEnum = "HTML"
	ViewFormatPng  ViewFormatEnum = "PNG"
	ViewFormatJpg  ViewFormatEnum = "JPG"
	ViewFormatPdf  ViewFormatEnum = "PDF"
)
