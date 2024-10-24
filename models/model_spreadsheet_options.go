/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Rendering options for Spreadsheet file formats. Spreadsheet file formats include files with extensions: .xls, .xlsx, .xlsm, .xlsb, .csv, .ods, .ots, .xltx, .xltm, .tsv
type SpreadsheetOptions struct {
	// Allows to enable worksheets pagination. By default one worksheet is rendered into one page.
	PaginateSheets bool `json:"PaginateSheets,omitempty"`
	// The number of rows rendered into one page when PaginateSheets is enabled. Default value is 50.
	CountRowsPerPage int32 `json:"CountRowsPerPage,omitempty"`
	// The columns count to include into each page when splitting worksheet into pages.
	CountColumnsPerPage int32 `json:"CountColumnsPerPage,omitempty"`
	// Indicates whether to render grid lines
	RenderGridLines bool `json:"RenderGridLines,omitempty"`
	// By default empty rows are skipped. Enable this option in case you want to render empty rows.
	RenderEmptyRows bool `json:"RenderEmptyRows,omitempty"`
	// By default empty columns are skipped. Enable this option in case you want to render empty columns.
	RenderEmptyColumns bool `json:"RenderEmptyColumns,omitempty"`
	// Enables rendering of hidden rows.
	RenderHiddenRows bool `json:"RenderHiddenRows,omitempty"`
	// Enables rendering of hidden columns.
	RenderHiddenColumns bool `json:"RenderHiddenColumns,omitempty"`
	// Enables headings rendering.
	RenderHeadings bool `json:"RenderHeadings,omitempty"`
	// Enables rendering worksheet(s) sections which is defined as print area. Renders each print area in a worksheet as a separate page.
	RenderPrintAreaOnly bool                 `json:"RenderPrintAreaOnly,omitempty"`
	TextOverflowMode    TextOverflowModeEnum `json:"TextOverflowMode,omitempty"`
}
