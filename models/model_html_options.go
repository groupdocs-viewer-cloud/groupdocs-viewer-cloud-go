/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Options for rendering document into HTML
type HtmlOptions struct {
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
	// Controls output HTML document resources (styles, images and fonts) linking. By default this option is disabled and all the resources are embedded into HTML document.
	ExternalResources bool `json:"ExternalResources,omitempty"`
	// Path for the HTML resources (styles, images and fonts). For example when resource path is http://example.com/api/pages/{page-number}/resources/{resource-name} the {page-number} and {resource-name} templates will be replaced with page number and resource name accordingly. This option is ignored when ExternalResources option is disabled.
	ResourcePath string `json:"ResourcePath,omitempty"`
	// Indicates whether rendering will provide responsive web pages, that look well on different device types. Default value is false.
	IsResponsive bool `json:"IsResponsive,omitempty"`
	// Enables HTML content and HTML resources minification
	Minify bool `json:"Minify,omitempty"`
	// When enabled prevents adding any fonts into HTML document
	ExcludeFonts bool `json:"ExcludeFonts,omitempty"`
	// This option is supported for presentations only. The list of font names, to exclude from HTML document
	FontsToExclude []string `json:"FontsToExclude,omitempty"`
	// Indicates whether to optimize output HTML for printing.
	ForPrinting bool `json:"ForPrinting,omitempty"`
	// The height of an output image in pixels. (When converting single image to HTML only)
	ImageHeight int32 `json:"ImageHeight,omitempty"`
	// The width of the output image in pixels. (When converting single image to HTML only)
	ImageWidth int32 `json:"ImageWidth,omitempty"`
	// Max height of an output image in pixels. (When converting single image to HTML only)
	ImageMaxHeight int32 `json:"ImageMaxHeight,omitempty"`
	// Max width of an output image in pixels. (When converting single image to HTML only)
	ImageMaxWidth int32 `json:"ImageMaxWidth,omitempty"`
	// Enables HTML content will be rendered to single page
	RenderToSinglePage bool `json:"RenderToSinglePage,omitempty"`
	// Allows to remove the JavaScript source code from the links in resultant HTML documents, when rendering input documents, which have the scripts. By default is enabled (true).
	RemoveJavaScript bool `json:"RemoveJavaScript,omitempty"`
}
