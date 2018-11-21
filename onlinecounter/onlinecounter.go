package onlinecounter

type OnlineCounter interface {
	Online(string, int64) error
	Count(string) (int64, error)
}

type OnlineCounterWrapper struct {
	OnlineCounter
}

func (o *OnlineCounterWrapper) Online(timeKey string, id int64) error {
	return o.OnlineCounter.Online(timeKey, id)
}
func (o *OnlineCounterWrapper) Count(timeKey string) (int64, error) {
	return o.OnlineCounter.Count(timeKey)
}
