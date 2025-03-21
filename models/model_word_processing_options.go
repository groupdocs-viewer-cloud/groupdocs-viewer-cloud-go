/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Provides options for rendering word processing documents
type WordProcessingOptions struct {
	// Enables tracked changes (revisions) rendering
	RenderTrackedChanges bool `json:"RenderTrackedChanges,omitempty"`
	// Left page margin (for HTML rendering only)
	LeftMargin float64 `json:"LeftMargin,omitempty"`
	// Right page margin (for HTML rendering only)
	RightMargin float64 `json:"RightMargin,omitempty"`
	// Top page margin (for HTML rendering only)
	TopMargin float64 `json:"TopMargin,omitempty"`
	// Bottom page margin (for HTML rendering only)
	BottomMargin float64 `json:"BottomMargin,omitempty"`
	// The size of the page.
	PageSize PageSizeEnum `json:"PageSize,omitempty"`
	// This option enables kerning and other OpenType Features when rendering Arabic, Hebrew, Indian Latin-based, or Cyrillic-based scripts.
	EnableOpenTypeFeatures bool `json:"EnableOpenTypeFeatures,omitempty"`
	// When rendering to HTML or PDF, you can set this option to `true` to disable navigation from the table of contents. For HTML rendering, `a` tags with relative links will be replaced with `span` tags, removing functionality but preserving visual appearance. For PDF rendering, the table of contents will be rendered as plain text without links to document sections.
	UnlinkTableOfContents bool `json:"UnlinkTableOfContents,omitempty"`
	// Determines if fields of certain types should be updated before saving the input WordProcessing document to the HTML, PDF, PNG, or JPEG output formats. Default value for this property is true — fields will be updated before saving.
	UpdateFields bool `json:"UpdateFields,omitempty"`
}
