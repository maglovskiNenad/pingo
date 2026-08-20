package metrics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	TotalKB     uint64
	AvailableKB uint64
	UsedKB      uint64
	Usage       float64
}

func ReadMemory() (MemoryInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemoryInfo{}, err
	}
	defer file.Close()

	var total uint64
	var available uint64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)

			value, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return MemoryInfo{}, err
			}

			total = value
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)

			value, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return MemoryInfo{}, err
			}

			available = value
		}
	}

	if err := scanner.Err(); err != nil {
		return MemoryInfo{}, err
	}

	if total == 0 {
		return MemoryInfo{}, fmt.Errorf("MemTotal not found")
	}

	used := total - available
	usage := float64(used) / float64(total) * 100

	return MemoryInfo{
		TotalKB:     total,
		AvailableKB: available,
		UsedKB:      used,
		Usage:       usage,
	}, nil
}