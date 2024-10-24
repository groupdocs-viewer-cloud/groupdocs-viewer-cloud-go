/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Options for rendering document into image
type ImageOptions struct {
	// Page number from which rendering should be started
	StartPageNumber int32 `json:"StartPageNumber,omitempty"`
	// Count pages which should be rendered
	CountPagesToRender int32 `json:"CountPagesToRender,omitempty"`
	// Pages list to render. Ignored, if StartPageNumber and CountPagesToRender are provided
	PagesToRender []int32 `json:"PagesToRender,omitempty"`
	// Page rotations
	PageRotations []PageRotation `json:"PageRotations,omitempty"`
	// Default font name may be specified in following cases: - You want to generally specify the default font to fall back on, if particular font   in the document cannot be found during rendering. - Your document uses fonts, that contain non-English characters and you want to make sure   any missing font is replaced with one that has the same character set available.
	DefaultFontName string `json:"DefaultFontName,omitempty"`
	// Default encoding for the plain text files such as .csv, .txt and .eml files when encoding is not specified in header
	DefaultEncoding string `json:"DefaultEncoding,omitempty"`
	// This option enables TXT, TSV, and CSV files encoding detection. In case the encoding can't be detected the DefaultEncoding is used.
	DetectEncoding bool `json:"DetectEncoding,omitempty"`
	// When enabled comments will be rendered to the output
	RenderComments bool `json:"RenderComments,omitempty"`
	// When enabled notes will be rendered to the output
	RenderNotes bool `json:"RenderNotes,omitempty"`
	// When enabled hidden pages, sheets or slides will be rendered to the output
	RenderHiddenPages bool `json:"RenderHiddenPages,omitempty"`
	// Rendering options for Spreadsheet source file formats Spreadsheet file formats include files with extensions: .xls, .xlsx, .xlsm, .xlsb, .csv, .ods, .ots, .xltx, .xltm, .tsv
	SpreadsheetOptions *SpreadsheetOptions `json:"SpreadsheetOptions,omitempty"`
	// Rendering options for CAD source file formats CAD file formats include files with extensions: .dwg, .dxf, .dgn, .ifc, .stl
	CadOptions *CadOptions `json:"CadOptions,omitempty"`
	// Rendering options for Email source file formats Email file formats include files with extensions: .msg, .eml, .emlx, .ifc, .stl
	EmailOptions *EmailOptions `json:"EmailOptions,omitempty"`
	// Rendering options for MS Project source file formats Project file formats include files with extensions: .mpt, .mpp
	ProjectManagementOptions *ProjectManagementOptions `json:"ProjectManagementOptions,omitempty"`
	// Rendering options for PDF source file formats
	PdfDocumentOptions *PdfDocumentOptions `json:"PdfDocumentOptions,omitempty"`
	// Rendering options for WordProcessing source file formats
	WordProcessingOptions *WordProcessingOptions `json:"WordProcessingOptions,omitempty"`
	// Rendering options for Outlook source file formats
	OutlookOptions *OutlookOptions `json:"OutlookOptions,omitempty"`
	// Rendering options for Archive source file formats
	ArchiveOptions *ArchiveOptions `json:"ArchiveOptions,omitempty"`
	// Rendering options for Text source file formats
	TextOptions *TextOptions `json:"TextOptions,omitempty"`
	// Rendering options for Mail storage (Lotus Notes, MBox) data files.
	MailStorageOptions *MailStorageOptions `json:"MailStorageOptions,omitempty"`
	// Rendering options for Visio source file formats
	VisioRenderingOptions *VisioRenderingOptions `json:"VisioRenderingOptions,omitempty"`
	// This rendering options enables you to customize the appearance of the output HTML/PDF/PNG/JPEG when rendering Web documents.
	WebDocumentOptions *WebDocumentOptions `json:"WebDocumentOptions,omitempty"`
	// Allows to specify output image width.  Specify image width in case when you want to change output image dimensions. When Width has value and Height value is 0 then Height value will be calculated  to save image proportions.
	Width int32 `json:"Width,omitempty"`
	// Allows to specify output image height.  Specify image height in case when you want to change output image dimensions. When Height has value and Width value is 0 then Width value will be calculated  to save image proportions.
	Height int32 `json:"Height,omitempty"`
	// When enabled Viewer will extract text when it's possible (e.g. raster formats don't have text layer) and return it in the viewing result. This option might be useful when you want to add selectable text layer over the image.
	ExtractText bool `json:"ExtractText,omitempty"`
	// Allows to specify quality when rendering as JPG. Valid values are between 1 and 100.  Default value is 90.
	JpegQuality int32 `json:"JpegQuality,omitempty"`
	// Max width of an output image in pixels
	MaxWidth int32 `json:"MaxWidth,omitempty"`
	// Max height of an output image in pixels
	MaxHeight int32 `json:"MaxHeight,omitempty"`
}
