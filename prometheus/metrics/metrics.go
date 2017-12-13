package metrics

type Metric interface {
	Tick()
}
