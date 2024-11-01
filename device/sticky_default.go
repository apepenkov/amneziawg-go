//go:build !linux

package device

import (
	"github.com/apepenkov/amneziawg-go/conn"
	"github.com/apepenkov/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
