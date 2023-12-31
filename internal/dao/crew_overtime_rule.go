// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_config/internal/dao/internal"
)

// internalCrewOvertimeRuleDao is internal type for wrapping internal DAO implements.
type internalCrewOvertimeRuleDao = *internal.CrewOvertimeRuleDao

// crewOvertimeRuleDao is the data access object for table cqgf_crew_overtime_rule.
// You can define custom methods on it to extend its functionality as you wish.
type crewOvertimeRuleDao struct {
	internalCrewOvertimeRuleDao
}

var (
	// CrewOvertimeRule is globally public accessible object for table cqgf_crew_overtime_rule operations.
	CrewOvertimeRule = crewOvertimeRuleDao{
		internal.NewCrewOvertimeRuleDao(),
	}
)

// Fill with you ideas below.
