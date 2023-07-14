// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_config/internal/dao/internal"
)

// internalCrewHoursIndexDao is internal type for wrapping internal DAO implements.
type internalCrewHoursIndexDao = *internal.CrewHoursIndexDao

// crewHoursIndexDao is the data access object for table cqgf_crew_hours_index.
// You can define custom methods on it to extend its functionality as you wish.
type crewHoursIndexDao struct {
	internalCrewHoursIndexDao
}

var (
	// CrewHoursIndex is globally public accessible object for table cqgf_crew_hours_index operations.
	CrewHoursIndex = crewHoursIndexDao{
		internal.NewCrewHoursIndexDao(),
	}
)

// Fill with you ideas below.
