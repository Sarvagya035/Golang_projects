package models

type System_meterics struct {
	Hostname               string `json:"hostname"`
	Total_memory           string `json:"total_memory"`
	Free_memory            string `json:"free_memory"`
	Memory_used_percentage string `json:"memory_used_percentage"`
	Architecture           string `json:"architecture"`
	OS                     string `json:"operating_system"`
	CPU_cores              string `json:"cpu_cores"`
	CPU_usage              string `json:"cpu_usage"`
}
