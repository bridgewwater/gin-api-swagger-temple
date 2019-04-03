package ssc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// @Summary ssc/health
// @Description HealthCheck shows OK as the ping-pong result.
// @Tags sys
// @Success 200 "OK" "OK"
// @Router /ssc/health [get]
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// @Summary ssc/disk
// @Description HealthCheck DiskCheck checks the disk usage.
// @Tags sys
// @Success 200 "OK" "OK | >= 95 CRITICAL | >=90 WARNING"
// @failure 429 "DISK WARNING" "need check"
// @failure 500 "DISK CRITICAL" "must check"
// @Router /ssc/disk [get]
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}

// @Summary ssc/cpu
// @Description HealthCheck CPUCheck checks the cpu usage.
// @Tags sys
// @Success 200 "CPU OK" "server run ok"
// @failure 429 "CPU WARNING" "need check"
// @failure 500 "CPU CRITICAL" "must check"
// @Router /ssc/cpu [get]
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	c.String(status, "\n"+message)
}

// @Summary ssc/ram
// @Description HealthCheck RAMCheck checks the disk usage.
// @Tags sys
// @Success 200 "RAM info" "OK | >= 95 CRITICAL | >=90 WARNING"
// @failure 429 "RAM WARNING" "need check"
// @failure 500 "RAM CRITICAL" "must check"
// @Router /ssc/ram [get]
func RAMCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}
