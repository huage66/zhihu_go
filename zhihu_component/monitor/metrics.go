package monitor

type Options func() interface{}

var (

)

func UsePrometheus(options ...Options) {

}

// 默认的监控指标
func DefaultMetrics() {

}