/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Rendering options for CAD file formats. CAD file formats include files with extensions: .dwg, .dxf, .dgn, .ifc, .stl
type CadOptions struct {
	// Scale factor allows to change the size of the output document. Values higher than 1 will enlarge output result and values between 0 and 1 will make output result smaller. This option is ignored when either Height or Width options are set.
	ScaleFactor float64 `json:"ScaleFactor,omitempty"`
	// Width of the output result in pixels
	Width int32 `json:"Width,omitempty"`
	// Height of the output result in pixels
	Height int32 `json:"Height,omitempty"`
	// The drawing regions to render This option supported only for DWG and DWT file types The RenderLayouts and LayoutName options are ignored when rendering by tiles
	Tiles []Tile `json:"Tiles,omitempty"`
	// Indicates whether layouts from CAD document should be rendered
	RenderLayouts bool `json:"RenderLayouts,omitempty"`
	// The name of the specific layout to render. Layout name is case-sensitive
	LayoutName string `json:"LayoutName,omitempty"`
	// The CAD drawing layers to render By default all layers are rendered; Layer names are case-sensitive
	Layers []string `json:"Layers,omitempty"`
}
