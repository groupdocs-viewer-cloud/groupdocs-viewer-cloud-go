/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

// Represents drawing region
type Tile struct {
	// The X coordinate of the lowest left point on the drawing where the tile begins
	StartPointX int32 `json:"StartPointX"`
	// The Y coordinate of the lowest left point on the drawing where the tile begins
	StartPointY int32 `json:"StartPointY"`
	// The width of the tile in pixels
	Width int32 `json:"Width"`
	// The height of the tile in pixels
	Height int32 `json:"Height"`
}
