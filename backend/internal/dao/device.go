// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalDeviceDao is an internal type for wrapping the internal DAO implementation.
type internalDeviceDao = *internal.DeviceDao

// deviceDao is the data access object for the table device.
// You can define custom methods on it to extend its functionality as needed.
type deviceDao struct {
	internalDeviceDao
}

var (
	// Device is a globally accessible object for table device operations.
	Device = deviceDao{
		internal.NewDeviceDao(),
	}
)

// Add your custom methods and functionality below.
