package reportevents

import "fflogs/structure"

// TODO
type Summons struct {
	structure.BaseResponse
	Events            interface{} `json:"events"`
	Count             int64       `json:"count"`
	NextPageTimestamp int64       `json:"nextPageTimestamp"`
}
