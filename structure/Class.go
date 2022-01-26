package structure

type Class struct {
	BaseResponse
	ID    int         `json:"id"`    // A unique identifier representing this specific class.,
	Name  string      `json:"name"`  // The English name of the class.,
	Specs []ClassSpec `json:"specs"` // The possible specs for this class.
}
type ClassSpec struct {
	ID   int    `json:"id"`   // A unique identifier representing this specific spec.,
	Name string `json:"name"` // The English name of the spec.
}
