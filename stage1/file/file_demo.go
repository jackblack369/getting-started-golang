package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

const (
	defaultClientExampleConfPath = "/Users/dongwei/Documents/tmp/csi/client.conf"
	cacheDirPrefix               = "/Users/dongwei/Documents/tmp/csi/client/data/cache/"
)

type curvefsMounter struct {
	mounterParams map[string]string
}

func NewCurvefsMounter() *curvefsMounter {
	return &curvefsMounter{mounterParams: map[string]string{}}
}

func (cm *curvefsMounter) applyMountFlags_v3(
	origConfPath string,
	mountFlags []string,
	targetPath string,
	mountUUID string,
) (string, error) {
	confPath := defaultClientExampleConfPath + "." + mountUUID

	// Step 1: Copy the original configuration file to a new file
	data, err := os.ReadFile(origConfPath)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			"applyMountFlag: failed to read conf %v",
			origConfPath,
			err,
		)
	}
	err = os.WriteFile(confPath, data, 0644)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			"applyMountFlag: failed to write new conf %v",
			confPath,
			err,
		)
	}

	// Step 2: Read the new configuration file
	data, err = os.ReadFile(confPath)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			"applyMountFlag: failed to read new conf %v",
			confPath,
			err,
		)
	}

	// Step 3: Iterate over the mountFlags items
	lines := strings.Split(string(data), "\n")
	configMap := make(map[string]string)
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		configMap[parts[0]] = parts[1]
	}

	fmt.Println("configMap is ", configMap)

	cacheEnabled := false
	for _, flag := range mountFlags {
		parts := strings.SplitN(flag, "=", 2)
		if len(parts) == 2 {
			configMap[parts[0]] = parts[1]
			if parts[0] == "diskCache.diskCacheType" && (parts[1] == "2" || parts[1] == "1") {
				cacheEnabled = true
			}
		}
	}

	// Step 4: Write the updated configuration back to the new file
	var newData strings.Builder
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			newData.WriteString(line + "\n")
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if newValue, exists := configMap[parts[0]]; exists {
			if parts[0] == "disk_cache.cache_dir" {
				cacheDirs := strings.Split(newValue, ";")
				cacheDirsWithUUID := make([]string, len(cacheDirs))
				for i, cacheDir := range cacheDirs {
					cacheDirParts := strings.SplitN(cacheDir, ":", 2)
					if len(cacheDirParts) == 2 {
						cacheDirsWithUUID[i] = fmt.Sprintf("%s/%s:%s", cacheDirParts[0], mountUUID, cacheDirParts[1])
					} else {
						cacheDirsWithUUID[i] = fmt.Sprintf("%s/%s", cacheDirParts[0], mountUUID)
					}
				}
				newValue = strings.Join(cacheDirsWithUUID, ";")
			}
			newData.WriteString(fmt.Sprintf("%s=%s\n", parts[0], newValue))
			delete(configMap, parts[0])
		} else {
			newData.WriteString(line + "\n")
		}
	}

	// Add the new configuration items
	for key, value := range configMap {
		newData.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	}

	err = os.WriteFile(confPath, []byte(newData.String()), 0644)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			"applyMountFlag: failed to write updated conf %v",
			confPath,
			err,
		)
	}

	if cacheEnabled {
		cacheDir := cacheDirPrefix + mountUUID
		if err := os.MkdirAll(cacheDir, 0777); err != nil {
			return "", err
		}
	}

	return confPath, nil
}

func main() {
	curvefsMounter := NewCurvefsMounter()
	/**
		mountOptions:
		  - diskCache.diskCacheType=2
		  - block_cache.cache_store=disk  # disk or none
	  	  - disk_cache.cache_dir=/var/run/dingofs  # "/data1:100;/data2:200"
	  	  - disk_cache.cache_size_mb=102400   # MB
	*/
	mountOptions := []string{
		"diskCache.diskCacheType=2",
		"block_cache.cache_store=disk",
		"disk_cache.cache_dir=/curvefs/client/data/cache/4:512;/curvefs/client/data/cache/5:1024;/curvefs/client/data/cache/6:2048",
		"disk_cache.cache_size_mb=204800",
	}
	mountUUID := uuid.New().String()
	curvefsMounter.applyMountFlags_v3(defaultClientExampleConfPath, mountOptions, "/Users/dongwei/Documents/tmp/csi", mountUUID)
}
