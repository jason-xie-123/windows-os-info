//go:build windows
// +build windows

package main

import (
	"errors"
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func isWin10AndAbove() bool {
	versionInfo := windows.RtlGetVersion()

	if versionInfo.MajorVersion > 10 {
		return true
	} else if versionInfo.MajorVersion == 10 && versionInfo.BuildNumber > 10240 {
		return true
	}

	return false
}

func getCPUNum() (int, error) {
	return runtime.NumCPU(), nil
}

func GetOSVersion() (string, error) {
	versionInfo := windows.RtlGetVersion()
	return fmt.Sprintf("%d.%d.%d", versionInfo.MajorVersion, versionInfo.MinorVersion, versionInfo.BuildNumber), nil
}

func getOSArch() (string, error) {
	// IsWow64Process2（从 Windows 10 开始支持）
	if isWin10AndAbove() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		isWow64Process2 := kernel32.NewProc("IsWow64Process2")

		var hProcess syscall.Handle
		hProcess, err := syscall.GetCurrentProcess()
		if err != nil {
			return "", err
		}

		var processMachine uint16
		var nativeMachine uint16

		ret, _, callErr := isWow64Process2.Call(
			uintptr(hProcess),
			uintptr(unsafe.Pointer(&processMachine)),
			uintptr(unsafe.Pointer(&nativeMachine)),
		)

		if ret == 0 {
			return "", errors.New(callErr.Error())
		}

		// https://learn.microsoft.com/en-us/windows/win32/sysinfo/image-file-machine-constants
		const (
			imageFileMachineUnknown = 0
			imageFileMachineI386    = 0x014c // x86
			imageFileMachineAMD64   = 0x8664 // x64
			imageFileMachineARM64   = 0xAA64 // ARM64
		)

		switch nativeMachine {
		case imageFileMachineI386:
			return "x86", nil
		case imageFileMachineAMD64:
			return "x64", nil
		case imageFileMachineARM64:
			return "arm64", nil
		default:
			return "", fmt.Errorf("unknown (0x%x)", nativeMachine)
		}
	} else {
		// GetNativeSystemInfo 返回的是系统的“逻辑架构”（仿真平台的宿主架构），而不是实际机器的物理 CPU, 在 ARM64 设备上，如果是通过 x86 程序调用，
		// 它会认为自己是运行在 WOW64 模式下的 AMD64 系统上。如果希望在任何架构下都能判断出是否是 真正的 ARM64 物理 CPU，可以使用更底层的 API：
		// IsWow64Process2（从 Windows 10 开始支持）
		type PROCESSOR_ARCH struct {
			ProcessorArchitecture uint16
			Reserved              uint16
		}

		type SYSTEM_INFO struct {
			Arch                        PROCESSOR_ARCH
			DwPageSize                  uint32
			LpMinimumApplicationAddress uintptr
			LpMaximumApplicationAddress uintptr
			DwActiveProcessorMask       uint
			DwNumberOfProcessors        uint32
			DwProcessorType             uint32
			DwAllocationGranularity     uint32
			WProcessorLevel             uint16
			WProcessorRevision          uint16
		}

		kernel32 := windows.NewLazySystemDLL("kernel32.dll")
		procGetNativeSystemInfo := kernel32.NewProc("GetNativeSystemInfo")

		var info SYSTEM_INFO
		_, _, _ = procGetNativeSystemInfo.Call(uintptr(unsafe.Pointer(&info)))

		// https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info#members
		const (
			PROCESSOR_ARCHITECTURE_AMD64 = 9
			PROCESSOR_ARCHITECTURE_ARM   = 5
			PROCESSOR_ARCHITECTURE_ARM64 = 12
			PROCESSOR_ARCHITECTURE_IA64  = 6
			PROCESSOR_ARCHITECTURE_INTEL = 0
		)

		switch info.Arch.ProcessorArchitecture {
		case PROCESSOR_ARCHITECTURE_AMD64:
			return "x64", nil
		case PROCESSOR_ARCHITECTURE_INTEL:
			return "x86", nil
		case PROCESSOR_ARCHITECTURE_ARM64:
			return "arm64", nil
		case PROCESSOR_ARCHITECTURE_ARM:
			return "arm", nil
		case PROCESSOR_ARCHITECTURE_IA64:
			return "ia64", nil
		default:
			return "", fmt.Errorf("unknown (%d)", info.Arch.ProcessorArchitecture)
		}
	}
}
