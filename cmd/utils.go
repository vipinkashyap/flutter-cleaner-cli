package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func folderSize(path string) int64 {
    var size int64
    filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err == nil && !info.IsDir() {
            size += info.Size()
        }
        return nil
    })
    return size
}

func formatSize(bytes int64) string {
    mb := float64(bytes) / 1024 / 1024
    if mb < 1024 {
        return fmt.Sprintf("%.2f MB", mb)
    }
    gb := mb / 1024
    return fmt.Sprintf("%.2f GB", gb)
}