// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gea/app/dao/internal"
)

// sysLogininforDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysLogininforDao struct {
	*internal.SysLogininforDao
}

var (
	// SysLogininfor is globally public accessible object for table sys_logininfor operations.
	SysLogininfor = &sysLogininforDao{
		internal.SysLogininfor,
	}
)

// Fill with you ideas below.