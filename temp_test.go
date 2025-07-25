package getting_started

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCode(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func parsePaths(input string) map[string]string {
	// Create a map to store the host -> container path mappings
	result := make(map[string]string)

	// Split the input string by ';' to get each host:container pair
	pairs := strings.Split(input, ";")

	for _, pair := range pairs {
		// Split each pair by ':' to separate host and container paths
		paths := strings.Split(pair, ":")
		if len(paths) == 2 {
			hostPath := paths[0]
			containerPath := paths[1]
			// Add the host and container paths to the map
			result[hostPath] = containerPath
		}
	}

	return result
}

func Test1(t *testing.T) {
	// Example input
	input := "host_path_1:container_path_1;host_path_2:container_path_2;host_path_3:container_path_3"

	// Parse the input into a map
	parsedPaths := parsePaths(input)

	// Print the result
	for host, container := range parsedPaths {
		fmt.Printf("Host: %s -> Container: %s\n", host, container)
	}
}

func TestTimerSchedule(t *testing.T) {
	// schedule a task to run every second
	for t := range time.Tick(1 * time.Second) {
		fmt.Println(t)
	}
}

func TestTimerFormat(t *testing.T) {

	defaultName := "gateway-" + time.Now().Format("20060102150405")
	fmt.Println(defaultName)

}

func TestTrim(t *testing.T) {
	original := "/curvefs/client/mnt/mnt/test3"
	prefix := "/curvefs/client/mnt"

	// Remove the prefix
	extracted := strings.TrimPrefix(original, prefix)

	fmt.Println(extracted) // Output: /mnt/test3
}

func TestAppend(t *testing.T) {
	// Append elements to a slice
	var numbers []string
	numbers = append(numbers, "a")
	numbers = append(numbers, "b", "c")
	fmt.Println(numbers) // Output: [1 2 3 4]

	resStr := strings.Join(numbers, "")
	fmt.Println(resStr)
}
