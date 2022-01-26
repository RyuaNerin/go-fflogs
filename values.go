package fflogs

type Region string

func (v Region) String() string {
	return string(v)
}

var (
	RegionUS Region = "US"
	RegionEU Region = "EU"
	RegionKR Region = "KR"
	RegionTW Region = "TW"
	RegionCN Region = "CN"
)

type Metric string

func (v Metric) String() string {
	return string(v)
}

var (
	MetricDps         Metric = "dps"
	MetricHps         Metric = "hps"
	MetricBossDps     Metric = "bossdps"
	MetricTankHps     Metric = "tankhps"
	MetricPlayerSpeed Metric = "playerspeed"
)

type Timeframe string

func (v Timeframe) String() string {
	return string(v)
}

var (
	TimeframeToday      Timeframe = "today"
	TimeframeHistorical Timeframe = "historical"
)
