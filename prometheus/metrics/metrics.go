package metrics

type Metric interface {
	Collect()
}
