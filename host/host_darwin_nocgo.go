// +build darwin
// +build !cgo

package host

import "github.com/M0Rf30/gopsutil/internal/common"

func SensorsTemperatures() ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
