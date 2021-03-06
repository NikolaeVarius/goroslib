// Autogenerated with msg-import, do not edit.
package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Inertia struct {
	msgs.Package `ros:"geometry_msgs"`
	M            msgs.Float64
	Com          Vector3
	Ixx          msgs.Float64
	Ixy          msgs.Float64
	Ixz          msgs.Float64
	Iyy          msgs.Float64
	Iyz          msgs.Float64
	Izz          msgs.Float64
}
