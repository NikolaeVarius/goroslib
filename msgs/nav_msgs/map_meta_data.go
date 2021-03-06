// Autogenerated with msg-import, do not edit.
package nav_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
)

type MapMetaData struct {
	msgs.Package `ros:"nav_msgs"`
	MapLoadTime  msgs.Time
	Resolution   msgs.Float32
	Width        msgs.Uint32
	Height       msgs.Uint32
	Origin       geometry_msgs.Pose
}
