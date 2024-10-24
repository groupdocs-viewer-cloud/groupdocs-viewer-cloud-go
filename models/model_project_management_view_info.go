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

// Represents view information for MS Project document
type ProjectManagementViewInfo struct {
	// The date time from which the project started
	StartDate time.Time `json:"StartDate"`
	// The date time when the project is to be completed
	EndDate time.Time `json:"EndDate"`
}
