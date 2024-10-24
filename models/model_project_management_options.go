/*
 * GroupDocs.Viewer Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
*
*/

package models

import (
	"time"
)

// Rendering options for Project file formats. Project file formats include files with extensions: .mpt, .mpp
type ProjectManagementOptions struct {
	PageSize PageSizeEnum `json:"PageSize,omitempty"`
	TimeUnit TimeUnitEnum `json:"TimeUnit,omitempty"`
	// The start date of a Gantt Chart View to render.
	StartDate time.Time `json:"StartDate,omitempty"`
	// The end date of a Gantt Chart View to render.
	EndDate time.Time `json:"EndDate,omitempty"`
}
