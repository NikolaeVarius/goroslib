// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
)

type InteractiveMarkerControl struct {
	msgs.Package                 `ros:"visualization_msgs"`
	Name                         msgs.String
	Orientation                  geometry_msgs.Quaternion
	OrientationMode              msgs.Uint8
	InteractionMode              msgs.Uint8
	AlwaysVisible                msgs.Bool
	Markers                      []Marker
	IndependentMarkerOrientation msgs.Bool
	Description                  msgs.String
}
