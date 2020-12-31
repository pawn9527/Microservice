package discovery

import "context"

// 服务实例结构体
type InstanceInfo struct {
	ID                string            `json:"ID"`                // 服务器实例ID
	Service           string            `json:"Service,omitempty"` //服务发现时返回的服务名
	Name              string            `json:"Name"`              // 服务名
	Tags              []string          `json:"Tags,omitempty"`    // 标签，可用于进行服务过滤
	Address           string            `json:"Address"`           // 服务器实例HOsT
	Port              int               `json:"Port"`              // 服务实例端口
	Meta              map[string]string `json:"Meta,omitempty"`    // 元数据
	EnableTagOverride bool              `json:"EnableTagOverride"` // 是否允许标签覆盖
	Check             `json:"Check,omitempty"`                     // 健康检查相关配置
	Weights           `json:"Weights,omitempty"`                   // 权重
}

type Check struct {
	DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter"` // 多久之后注销服务
	Args                           []string `json:"Args,omitempty"`                 // 请求参数
	HTTP                           string   `json:"HTTP"`                           // 健康检查地址
	Interval                       string   `json:"Interval,omitempty"`             // Consul 主动检查间隔
	TTL                            string   `json:"TTL,omitempty"`                  // 服务实例主动维持
}

type Weights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
}

type DiscoveryClient struct {
	host string //Consul的Host
	port int    // Consul 的端口
}

func NewDiscoveryClient(host string, port int) *DiscoveryClient {
	return &DiscoveryClient{
		host: host,
		port: port,
	}
}

func (consulClient *DiscoveryClient) Register(ctx context.Context, serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, weights *Weights) error {


}
