/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// View options
type ViewOptions struct {
	// File info
	FileInfo   *FileInfo      `json:"FileInfo,omitempty"`
	ViewFormat ViewFormatEnum `json:"ViewFormat,omitempty"`
	// The output path Default value is 'viewer\\{input file path}_{file extension}\\'
	OutputPath string `json:"OutputPath,omitempty"`
	// The path to directory containing custom fonts in storage
	FontsPath string `json:"FontsPath,omitempty"`
	// Text watermark
	Watermark *Watermark `json:"Watermark,omitempty"`
	// Rendering options
	RenderOptions interface{} `json:"RenderOptions,omitempty"`
}
