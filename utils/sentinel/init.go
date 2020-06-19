package middleware

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/system"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	FlowPrefix    = "flow"
	BreakerPrefix = "breaker"
)

type Config struct {
	Sentinel Sentinel `json:"sentinel"`
}

type Sentinel struct {
	Flows    []FlowRule    `json:"flows"`
	Breakers []BreakerRule `json:"breakers"`
	System   []SystemRule  `json:"system"`
}

type FlowRule struct {
	Resource          string  `yaml:"resource"`          // 资源名
	MetricType        int     `yaml:"metricType"`        // 流量控制类型，0 并发， 1 QPS
	Count             float64 `yaml:"count"`             // 每秒只能通过多少个请求
	ControlBehavior   int     `yaml:"controlBehavior"`   // 流量控制效果，0 直接拒绝， 2 排队等待
	MaxQueueingTimeMs uint32  `yaml:"maxQueueingTimeMs"` // 排队等待最长等待时间
}

type BreakerRule struct {
	Resource            string  `yaml:"resource"`            // 资源名
	Strategy            int     `yaml:"strategy"`            // 熔断控制类型， 0慢请求, 1错误百分比, 2错误请求统计
	RetryTimeoutMs      uint32  `yaml:"retryTimeoutMs"`      // 熔断后等待多长时间后重试
	MinRequestAmount    uint64  `yaml:"minRequestAmount"`    // 触发熔断的最小请求数目，若当前统计窗口内的请求数小于此值，即使达到熔断条件规则也不会触发。
	StatIntervalMs      uint32  `yaml:"statIntervalMs"`      // 统计的时间窗口长度(单位ms)
	MaxAllowedRt        uint64  `yaml:"maxAllowedRt"`        // 响应时间超过该值，将被标记为慢请求
	MaxSlowRequestRatio float64 `yaml:"maxSlowRequestRatio"` // 最大响应时间，只在Strategy为0时有效
	MaxErrorRatio       float64 `yaml:"maxErrorRatio"`       // 最大错误占比，只在Strategy为1时有效
	MaxErrorCount       uint64  `yaml:"maxErrorCount"`       // 最大错误数，只在Strategy为2时有效
}

// 系统自适应流控
type SystemRule struct {
	MetricType   int     `yaml:"metricType"`
	TriggerCount float64 `yaml:"count"`
	Strategy     int     `yaml:"strategy"`
}

func init() {
	err := api.Init("gin-test/config/config.yaml")
	if err != nil {
		panic(err)
	}

	// 加载流量控制规则
	f, _ := ioutil.ReadFile("gin-test/config/config.yaml")
	var config Config
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}
	var flowRules []*flow.FlowRule
	for _, r := range config.Sentinel.Flows {
		flowRules = append(flowRules, &flow.FlowRule{
			Resource:          FlowPrefix + r.Resource,
			MetricType:        flow.MetricType(r.MetricType),
			ControlBehavior:   flow.ControlBehavior(r.ControlBehavior),
			Count:             r.Count,
			MaxQueueingTimeMs: r.MaxQueueingTimeMs,
		})
	}
	_, err = flow.LoadRules(flowRules)
	if err != nil {
		panic(err)
	}

	// 加载熔断降级控制规则
	var breakerRules []circuitbreaker.Rule
	for _, r := range config.Sentinel.Breakers {
		switch circuitbreaker.Strategy(r.Strategy) {
		case circuitbreaker.SlowRequestRatio:
			breakerRules = append(breakerRules, circuitbreaker.NewSlowRtRule(BreakerPrefix+r.Resource, r.StatIntervalMs, r.RetryTimeoutMs, r.MaxAllowedRt, r.MinRequestAmount, r.MaxSlowRequestRatio))
		case circuitbreaker.ErrorRatio:
			breakerRules = append(breakerRules, circuitbreaker.NewErrorRatioRule(BreakerPrefix+r.Resource, r.StatIntervalMs, r.RetryTimeoutMs, r.MinRequestAmount, r.MaxErrorRatio))
		case circuitbreaker.ErrorCount:
			breakerRules = append(breakerRules, circuitbreaker.NewErrorCountRule(BreakerPrefix+r.Resource, r.StatIntervalMs, r.RetryTimeoutMs, r.MinRequestAmount, r.MaxErrorCount))
		}
	}
	_, err = circuitbreaker.LoadRules(breakerRules)
	if err != nil {
		panic(err)
	}

	// 加载系统自适应流控规则
	var systemRules []*system.SystemRule
	for _, r := range config.Sentinel.System {
		systemRules = append(systemRules, &system.SystemRule{
			MetricType:   system.MetricType(r.MetricType),
			Strategy:     system.AdaptiveStrategy(r.Strategy),
			TriggerCount: r.TriggerCount,
		})
	}
	_, err = system.LoadRules(systemRules)
	if err != nil {
		panic(err)
	}
}
