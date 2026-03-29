package main

import (
	"log"

	device2 "github.com/Algatux/nvidia-gpu-manager/internal/device"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	log.Print("Starting nvidia-gpu-manager")
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Fatalf("Unable to initialize NVML: %v", nvml.ErrorString(ret))
	}

	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			log.Fatalf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Fatalf("Unable to get device count: %v", nvml.ErrorString(ret))
	}

	for i := 0; i < count; i++ {

		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, nvml.ErrorString(ret))
		}

		info := device2.NewDeviceInfo(device, i)

		log.Print(info)
		log.Print(info.GetFullName())
	}

}
