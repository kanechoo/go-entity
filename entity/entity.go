package entity

import (
	"time"
)

// IPTaskPO  ip_scan_tasks 表
type IPTaskPO struct {
	ID         int       `json:"id"`          // 记录ID
	TaskName   string    `json:"task_name"`   // 任务名称
	StartTime  time.Time `json:"start_time"`  // 任务开始时间
	EndTime    time.Time `json:"end_time"`    // 任务接收时间
	ASN        string    `json:"asn"`         // ASN编号
	Name       string    `json:"name"`        // ASN名称
	NetBlock   string    `json:"netblock"`    // ASN的网段
	NumIPs     int       `json:"num_ips"`     // IP数量
	Country    string    `json:"country"`     // ASN所属国家编码
	Domain     string    `json:"domain"`      // ASN所属者的域名
	CreateTime time.Time `json:"create_time"` // 任务的创建时间
	Runtime    int       `json:"run_time"`    // 任务的运行时间(单位:秒)
}

// IPNetBlockDTO 结构体
type IPNetBlockDTO struct {
	NetBlock string `json:"netblock"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Size     string `json:"size"`
	Status   string `json:"status"`
	Domain   string `json:"domain"`
}

// ASNInfoDTO ASN结构体
type ASNInfoDTO struct {
	ASN       string          `json:"asn"`
	Name      string          `json:"name"`
	Country   string          `json:"country"`
	Allocated string          `json:"allocated"`
	Registry  string          `json:"registry"`
	Domain    string          `json:"domain"`
	NumIPs    int             `json:"num_ips"`
	Type      string          `json:"type"`
	NetBlocks []IPNetBlockDTO `json:"prefixes"`
}

// ASNFilterDTO ASN过滤器
type ASNFilterDTO struct {
	// Asn ASN编号
	ASN int
	// Country ASN所在国家
	Country string
	//NetBlock ASN网段
	NetBlock string
}
type BaseIPDTO struct {
	ASN  string `json:"asn"`
	Ct   string `json:"country"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// CheckIPPO ip 表
type CheckIPPO struct {
	ID         int       `json:"id"`          // 记录ID
	IP         string    `json:"ip"`          // IP地址
	Port       uint16    `json:"port"`        // 端口
	ASN        string    `json:"asn"`         // IP地址所属ASN
	Country    string    `json:"country"`     // 国家
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	Check      int       `json:"check"`       //是否已检查可用性，1-已检查，0-未检查
}
type ValidIPPO struct {
	Query       string  `json:"query"`
	Port        int     `json:"port"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
}
