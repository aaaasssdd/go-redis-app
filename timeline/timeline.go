package timeline

type Timeline interface {
	Push(string, string) error
	FetchIndex(string, int64, int64) ([]string, error)
}

type TimeLineWrapper struct {
	Timeline
}

func (t *TimeLineWrapper) Add(key, message string) (err error) {
	return t.Timeline.Push(key, message)
}

func (t *TimeLineWrapper) FetchRecent(key string, length int64) ([]string, error) {
	return t.Timeline.FetchIndex(key, 0, length)
}

func (t *TimeLineWrapper) FetchFromIndex(key string, start, length int64) ([]string, error) {
	return t.Timeline.FetchIndex(key, start, length)
}
