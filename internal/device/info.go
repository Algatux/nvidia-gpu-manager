package device

import (
	"fmt"
	"math"

	"github.com/Algatux/nvidia-gpu-manager/internal/exception"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

const GiB = 1024 * 1024 * 1024

type DeviceInfo struct {
	Index  int
	Name   string
	Memory nvml.Memory_v2
}

func NewDeviceInfo(device nvml.Device, index int) *DeviceInfo {

	name, ret := device.GetName()
	exception.CheckReturn(func() string { return "unable to retrive device name" }, ret)

	mem, ret := device.GetMemoryInfo_v2()
	exception.CheckReturn(func() string { return "unable to retrieve device memory info" }, ret)

	return &DeviceInfo{
		Index:  index,
		Name:   name,
		Memory: mem,
	}
}

func (di *DeviceInfo) GetFullName() string {
	return fmt.Sprintf("%s %.0fGB", di.Name, convertBytesToGiB(di.Memory.Total))
}

func convertBytesToGiB(bytes uint64) float64 {
	return math.Round(float64(bytes) / float64(GiB))
}
