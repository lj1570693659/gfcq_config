// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_config/internal/dao/internal"
)

// internalProductTypeDao is internal type for wrapping internal DAO implements.
type internalProductTypeDao = *internal.ProductTypeDao

// productTypeDao is the data access object for table cqgf_product_type.
// You can define custom methods on it to extend its functionality as you wish.
type productTypeDao struct {
	internalProductTypeDao
}

var (
	// ProductType is globally public accessible object for table cqgf_product_type operations.
	ProductType = productTypeDao{
		internal.NewProductTypeDao(),
	}
)

// Fill with you ideas below.