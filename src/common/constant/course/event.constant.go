package course

var (
	// InsertDataTopicList topic name in convention as <application>.<service>.<method>.<status>
	InsertDataTopicList = []string{"cugetreg.scraper.scrape.pending", "cugetreg.backend.create.pending"}
)

const (
	ExchangeName = "course"
	ExchangeKind = "topic"
)
