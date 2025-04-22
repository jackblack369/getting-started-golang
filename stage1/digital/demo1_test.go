package digital

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	port1 := generateNodePort("pod1")
	port2 := generateNodePort("pod2")
	port3 := generateNodePort("pvc-cc5cb2f5-e6fb-4a31-9476-6d16dedbc607")
	port4 := generateNodePort("ceph17-test-pv-ceph-2")
	fmt.Println("port1 is ", port1)
	fmt.Println("port2 is ", port2)
	fmt.Println("pvc-cc5cb2f5-e6fb-4a31-9476-6d16dedbc607 is ", port3)
	fmt.Println("ceph17-test-pv-ceph-2 is ", port4)
}

// Helper function to generate a unique NodePort based on pod name
func generateNodePort(podName string) int {
	// Create a hash from the pod name to ensure uniqueness
	hash := sha256.Sum256([]byte(podName))
	// Use the hash to generate a number in the NodePort range [30000, 32767]
	nodePort := int(hash[0])%32767 + 30000
	if nodePort > 32767 {
		nodePort = 32767
	}
	return nodePort
}
