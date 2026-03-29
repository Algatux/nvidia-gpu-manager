package exception

import (
	"log"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func CheckReturn(messageFun func() string, ret nvml.Return) {
	if ret != nvml.SUCCESS {
		log.Fatal(messageFun())
	}
}
