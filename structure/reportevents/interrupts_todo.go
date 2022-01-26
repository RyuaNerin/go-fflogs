package reportevents

import "github.com/RyuaNerin/go-fflogs/structure"

// TODO
type Interrupts struct {
	structure.BaseResponse

	Events            interface{} `json:"events"`
	Count             int64       `json:"count"`
	NextPageTimestamp int64       `json:"nextPageTimestamp"`
}
