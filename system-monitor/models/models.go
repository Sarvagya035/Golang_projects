package models

type System_meterics struct {
	Hostname             string  `json:"hostname"`
	TotalMemory          uint64  `json:"total_memory"`
	FreeMemory           uint64  `json:"free_memory"`
	MemoryUsedPercentage float64 `json:"memory_used_percentage"`
	Architecture         string  `json:"architecture"`
	OS                   string  `json:"os"`
	CPUCores             int     `json:"number_of_cpu_cores"`
	CPUUsage             float64 `json:"cpu_used_percentage"`
}
