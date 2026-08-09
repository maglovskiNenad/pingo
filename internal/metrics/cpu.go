package metrics

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// cpuTimes stores the total CPU time and idle CPU time.
type cpuTimes struct {
	total uint64
	idle uint64
}

// readCPUTimes reads the first CPU line from /proc/stat
// and returns the total and idle CPU time.
func readCPUTimes() (cpuTimes, error) {
	// Open the Linux /proc/stat file.
	file, err := os.Open("/proc/stat")
	if err != nil {
		return cpuTimes{}, err
	}

	defer file.Close()
	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)
	
	// The first line contains the total CPU statistics.
	if scanner.Scan() {
		// Split the line into separate values.
		// Example:
		// cpu  123 45 67 890 ...
		fields := strings.Fields(scanner.Text())

		var values []uint64

		// Skip the first field ("cpu") and convert the numbers to uint64.
		for _, field := range fields[1:] {
			value, err := strconv.ParseUint(field, 10, 64)
			if err != nil {
				return cpuTimes{}, err
			}
			values = append(values,value)
		}
		// Add all CPU time values together.
		var total uint64
		

		for _, value := range values {
			total += value
		}

		// The fourth numeric value represents idle CPU time.
		idle := values[3]

		return cpuTimes {
			total: total,
			idle: idle,
		},nil
	}

	// Return a scanner error if reading the file failed.
	return cpuTimes{}, scanner.Err()
}

// CPUUsage calculates the current CPU usage as a percentage.
func CPUUsage() (float64, error) {
	// Take the first CPU measurement.
	first, err := readCPUTimes()
	if err != nil {
		return 0, err
	}
	// Wait a short time before taking the second measurement.
	time.Sleep(500* time.Millisecond)

	// Take the second CPU measurement.
	second, err := readCPUTimes()
	if err != nil {
		return 0, err 
	}

	// Calculate how much total CPU time passed
	// between the two measurements.
	totalDelta := second.total - first.total
	idleDelta := second.idle - first.idle

	// CPU usage is the part of the total time
	// where the CPU was not idle.
	usage := 100* (1 -float64(idleDelta)/float64(totalDelta))
	return usage, nil
}





















