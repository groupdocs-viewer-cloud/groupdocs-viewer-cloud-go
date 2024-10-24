/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering PDF documents
type PdfDocumentOptions struct {
	// Disables chars grouping to keep maximum precision during chars positioning when rendering the page
	DisableCharsGrouping bool `json:"DisableCharsGrouping,omitempty"`
	// Enables rendering of text and graphics according to z-order in original PDF document  when rendering into HTML
	EnableLayeredRendering bool `json:"EnableLayeredRendering,omitempty"`
	// Enables font hinting. The font hinting adjusts the display of an outline font. Supported only for TTF fonts when these fonts are used in source document.
	EnableFontHinting bool `json:"EnableFontHinting,omitempty"`
	// When this option enabled the output pages will have the same size in pixels as page size in a source PDF document. By default GroupDocs.Viewer calculates output image page size for better rendering quality. This option is supported when rendering into PNG or JPG formats.
	RenderOriginalPageSize bool             `json:"RenderOriginalPageSize,omitempty"`
	ImageQuality           ImageQualityEnum `json:"ImageQuality,omitempty"`
	// When this option is set to true, the text is rendered as an image in the output HTML. Enable this option to make text unselectable or to fix characters rendering and make HTML look like PDF. The default value is false. This option is supported when rendering into HTML.
	RenderTextAsImage bool `json:"RenderTextAsImage,omitempty"`
	// Enables rendering the PDF and EPUB documents to HTML with a fixed layout.
	FixedLayout bool `json:"FixedLayout,omitempty"`
	// Enables wrapping each image in the output HTML document in SVG tag to improve the output quality.
	WrapImagesInSvg bool `json:"WrapImagesInSvg,omitempty"`
	// Disables any license restrictions for all fonts in the current XPS/OXPS document.
	DisableFontLicenseVerifications bool `json:"DisableFontLicenseVerifications,omitempty"`
}
