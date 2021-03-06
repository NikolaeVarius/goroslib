// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Range struct {
	msgs.Package  `ros:"sensor_msgs"`
	Header        std_msgs.Header
	RadiationType msgs.Uint8
	FieldOfView   msgs.Float32
	MinRange      msgs.Float32
	MaxRange      msgs.Float32
	Range         msgs.Float32
}
