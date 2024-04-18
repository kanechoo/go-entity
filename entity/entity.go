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

// IPNetBlock 结构体
type IPNetBlock struct {
	NetBlock string `json:"netblock"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Size     string `json:"size"`
	Status   string `json:"status"`
	Domain   string `json:"domain"`
}

// ASNInfo ASN结构体
type ASNInfo struct {
	ASN       string       `json:"asn"`
	Name      string       `json:"name"`
	Country   string       `json:"country"`
	Allocated string       `json:"allocated"`
	Registry  string       `json:"registry"`
	Domain    string       `json:"domain"`
	NumIPs    int          `json:"num_ips"`
	Type      string       `json:"type"`
	NetBlocks []IPNetBlock `json:"prefixes"`
}

// ASNFilter ASN过滤器
type ASNFilter struct {
	// Asn ASN编号
	ASN int
	// Country ASN所在国家
	Country string
	//NetBlock ASN网段
	NetBlock string
}
type IPBase struct {
	ASN  string `json:"asn"`
	Ct   string `json:"country"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// RawIPPO ip 表
type RawIPPO struct {
	ID         int       `json:"id"`          // 记录ID
	IP         string    `json:"ip"`          // IP地址
	Port       uint16    `json:"port"`        // 端口
	ASN        string    `json:"asn"`         // IP地址所属ASN
	Country    string    `json:"country"`     // 国家
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
type ValidIPPO struct {
	Query       string  `json:"query"`
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
