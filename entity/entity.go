package entity

import (
	"time"
)

// IPTaskPO  ip_scan_tasks 表
type IPTaskPO struct {
	// 记录ID
	ID         int       `json:"id" db:"id"`
	TaskName   string    `json:"task_name" db:"task_name"`     // 任务名称
	StartTime  time.Time `json:"start_time" db:"start_time"`   // 任务开始时间
	EndTime    time.Time `json:"end_time" db:"end_time"`       // 任务结束时间
	ASN        string    `json:"asn" db:"asn"`                 // ASN编号
	Name       string    `json:"name" db:"name"`               // ASN名称
	NetBlock   string    `json:"netblock" db:"netblock"`       // ASN网段
	NumIPs     int       `json:"num_ips" db:"num_ips"`         // ASN包含的IP数量
	Country    string    `json:"country" db:"country"`         // ASN所属国家
	Domain     string    `json:"domain" db:"domain"`           // ASN所属者域名
	CreateTime time.Time `json:"create_time" db:"create_time"` // 创建时间
	Runtime    int       `json:"run_time" db:"run_time"`       // 运行时间
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
	// Asn ASN编号(用,分隔)
	ASN string
	// NetBlockCountry ASN的NetBlock所在国家(用,分隔)
	NetBlockCountry string
	//NetBlock ASN网段(用,分隔)
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
	ID         int       `json:"id" db:"id"`                   // 记录ID
	IP         string    `json:"ip" db:"ip"`                   // IP地址
	Port       uint16    `json:"port" db:"port"`               // 端口
	ASN        string    `json:"asn" db:"asn"`                 // ASN编号
	Country    string    `json:"country" db:"country"`         // 国家
	CreateTime time.Time `json:"create_time" db:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" db:"update_time"` // 更新时间
	Check      int       `json:"check" db:"check"`             // 是否检查过
}
type ValidIPPO struct {
	Query       string  `json:"query" db:"query"`
	Port        int     `json:"port" db:"port"`
	Status      string  `json:"status" db:"status"`
	Country     string  `json:"country" db:"country"`
	CountryCode string  `json:"countryCode" db:"country_code"`
	Region      string  `json:"region" db:"region"`
	RegionName  string  `json:"regionName" db:"region_name"`
	City        string  `json:"city" db:"city"`
	Zip         string  `json:"zip" db:"zip"`
	Latitude    float64 `json:"lat" db:"latitude"`
	Longitude   float64 `json:"lon" db:"longitude"`
	Timezone    string  `json:"timezone" db:"timezone"`
	ISP         string  `json:"isp" db:"isp"`
	Org         string  `json:"org" db:"org"`
	AS          string  `json:"as" db:"as"`
}
