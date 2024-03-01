//go:build !linux

package device

import (
	"github.com/ebrahimtahernejad/wireguard-warp/conn"
	"github.com/ebrahimtahernejad/wireguard-warp/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
