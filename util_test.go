//go:build windows
// +build windows

package main

import (
	"fmt"
	"testing"
)

// go test -run TestGetCPUNum
func TestGetCPUNum(t *testing.T) {
	cpuNum, err := getCPUNum()
	if err != nil {
		t.Error(fmt.Errorf("getCPUNum: err=%v", err))
		return
	}
	fmt.Printf("getCPUNum: cpuNum=%d\n", cpuNum)
}

// go test -run TestGetOSVersion
func TestGetOSVersion(t *testing.T) {
	osVersion, err := GetOSVersion()
	if err != nil {
		t.Error(fmt.Errorf("GetOSVersion: err=%v", err))
		return
	}
	fmt.Printf("GetOSVersion: osVersion=%s\n", osVersion)
}

// go test -run TestGetOSArch
func TestGetOSArch(t *testing.T) {
	osArch, err := getOSArch()
	if err != nil {
		t.Error(fmt.Errorf("getOSArch: err=%v", err))
		return
	}
	fmt.Printf("getOSArch: osArch=%s\n", osArch)
}
