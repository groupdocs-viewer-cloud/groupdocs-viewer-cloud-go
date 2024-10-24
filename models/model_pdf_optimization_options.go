/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Contains the PDF optimization options to apply to the output PDF file.
type PdfOptimizationOptions struct {
	// Enables optimization the output PDF file for viewing online with a web browser. This optimization allows a browser to display the first pages of a PDF file when     you open the document, instead of waiting for the entire file to download.
	Lineriaze bool `json:"Lineriaze,omitempty"`
	// Enables removing annotation from the output PDF file.
	RemoveAnnotations bool `json:"RemoveAnnotations,omitempty"`
	// Enables removing form fields from a PDF file.
	RemoveFormFields bool `json:"RemoveFormFields,omitempty"`
	// Enables converting the output PDF file to a grayscale.
	ConvertToGrayScale bool `json:"ConvertToGrayScale,omitempty"`
	// Subsets fonts in the output PDF file. If the file uses embedded fonts, it contains all font data. GroupDocs.Viewer can subset embedded fonts to reduce the file size.
	SubsetFonts bool `json:"SubsetFonts,omitempty"`
	// Enables compressing images in the output PDF file. Use this option to allow other compressing options: PdfOptimizationOptions.ImageQuality and PdfOptimizationOptions.MaxResolution.
	CompressImages bool `json:"CompressImages,omitempty"`
	// Sets the image quality in the output PDF file (in percent). To change the image quality, first set the PdfOptimizationOptions.CompressImages property to true.
	ImageQuality int32 `json:"ImageQuality,omitempty"`
	// Enables setting the maximum resolution in the output PDF file. To allow this option, set the GroupDocs.Viewer.Options.PdfOptimizationOptions.CompressImages property to true. This option allows setting the GroupDocs.Viewer.Options.PdfOptimizationOptions.MaxResolution property.
	ResizeImages bool `json:"ResizeImages,omitempty"`
	// Sets the maximum resolution in the output PDF file. To allow this option, set the GroupDocs.Viewer.Options.PdfOptimizationOptions.CompressImages and GroupDocs.Viewer.Options.PdfOptimizationOptions.MaxResolution properties to true. The default value is 300.
	MaxResolution int32 `json:"MaxResolution,omitempty"`
	// Enables optimization of spreadsheets in the PDF files. This optimization allows to reduce the output file size by setting up border lines. Besides that, it removes the Arial and Times New Roman characters of 32-127 codes.
	OptimizeSpreadsheets bool `json:"OptimizeSpreadsheets,omitempty"`
}
