package utils

type CPUSample struct {
	User, Nice, System, Idle, Iowait, Irq, Softirq, Steal uint64
}

type SystemMetrics struct {
	CPUUsagePercent    float64
	MemoryUsagePercent float64
	AvailableMemoryMB  int
	CPUCoreCount       int
}
