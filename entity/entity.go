package entity

import (
	"time"
)

// IPScanTask  ip_scan_tasks 表
type IPScanTask struct {
	ID         int       `files:"id"`          // 记录ID
	TaskName   string    `files:"task_name"`   // 任务名称
	StartTime  time.Time `files:"start_time"`  // 任务开始时间
	EndTime    time.Time `files:"end_time"`    // 任务接收时间
	ASN        string    `files:"asn"`         // ASN编号
	Name       string    `files:"name"`        // ASN名称
	Netblock   string    `files:"netblock"`    // ASN的网段
	NumIPs     int       `files:"num_ips"`     // IP数量
	Country    string    `files:"country"`     // ASN所属国家编码
	Domain     string    `files:"domain"`      // ASN所属者的域名
	CreateTime time.Time `files:"create_time"` // 任务的创建时间
	Runtime    int       `files:"run_time"`    // 任务的运行时间(单位:秒)
}

// IPPrefix 结构体
type IPPrefix struct {
	Netblock string `files:"netblock"`
	ID       string `files:"id"`
	Name     string `files:"name"`
	Country  string `files:"country"`
	Size     string `files:"size"`
	Status   string `files:"status"`
	Domain   string `files:"domain"`
}

// ASNInfo ASN结构体
type ASNInfo struct {
	ASN       string     `files:"asn"`
	Name      string     `files:"name"`
	Country   string     `files:"country"`
	Allocated string     `files:"allocated"`
	Registry  string     `files:"registry"`
	Domain    string     `files:"domain"`
	NumIPs    int        `files:"num_ips"`
	Type      string     `files:"type"`
	Prefixes  []IPPrefix `files:"prefixes"`
}

// ASNFilter ASN过滤器
type ASNFilter struct {
	// Asn ASN编号
	ASN int
	// Country ASN所在国家
	Country string
	//Netblock ASN网段
	Netblock string
}
type BaseIP struct {
	ASN  string `json:"asn"`
	Ct   string `json:"country"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}
