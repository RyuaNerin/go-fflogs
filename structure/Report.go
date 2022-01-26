package structure

type Report struct {
	BaseResponse
	ID        string   `json:"id"`        // (optional) The report code. Can be viewed on the site with the URL string of /reports/<code>.,
	Title     string   `json:"title"`     // (optional) The title for the report.,
	Owner     string   `json:"owner"`     // (optional) The name of the user that uploaded the report.,
	Zone      *int     `json:"zone"`      // (optional) The zone that the report is for. A value of -1 indicates that the zone for the report is not known.,
	StartTime *float32 `json:"startTime"` // (optional) A UNIX timestamp with millisecond precision, indicating the start time of the report as far as logged events.,
	EndTime   *float32 `json:"endTime"`   // (optional) A UNIX timestamp with millisecond precision, indicating the end time of the report as far as logged events.
}
