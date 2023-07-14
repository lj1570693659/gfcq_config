// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_config/internal/dao/internal"
)

// internalProductLevelAssessDao is internal type for wrapping internal DAO implements.
type internalProductLevelAssessDao = *internal.ProductLevelAssessDao

// productLevelAssessDao is the data access object for table cqgf_product_level_assess.
// You can define custom methods on it to extend its functionality as you wish.
type productLevelAssessDao struct {
	internalProductLevelAssessDao
}

var (
	// ProductLevelAssess is globally public accessible object for table cqgf_product_level_assess operations.
	ProductLevelAssess = productLevelAssessDao{
		internal.NewProductLevelAssessDao(),
	}
)

// Fill with you ideas below.
