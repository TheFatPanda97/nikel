package metrics

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/response"
	"runtime"
)

// GetMetrics returns runtime metrics for app health monitoring
func GetMetrics(c *gin.Context) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	response.SendSuccess(c, gin.H{
		"memory":           memStats.Alloc,
		"memory_humanized": humanize.Bytes(memStats.Alloc),
		"sys":              memStats.Sys,
		"sys_humanized":    humanize.Bytes(memStats.Sys),
		"routines":         runtime.NumGoroutine(),
	})
}
