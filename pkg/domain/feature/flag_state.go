package feature

type FlagState struct {
	Value         interface{}
	Timestamp     int64
	VariationType string
	TrackEvents   bool
	Failed        bool
	ErrorCode     string
	Reason        string
	Metadata      map[string]interface{}
}
