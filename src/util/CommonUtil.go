package util

// 공통으로 사용할 function 정의
import (
	// "encoding/base64"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "os"
	// "strconv"
	"strings"
	// "time"
	// "bytes"
	// "encoding/json"
	// "math"
	// "io/ioutil"
	// echosession "github.com/go-session/echo-session"
	// "github.com/labstack/echo"
	// "github.com/cloud-barista/cb-webtool/src/model"
)

// providerName 소문자로
func GetProviderName(provider string) string {
	return strings.ToLower(provider)
}

// MCIS 상태값의 앞부분만 사용. 소문자로
func GetMcisStatus(mcisStatus string) string {
	statusArr := strings.Split(mcisStatus, "-")
	status := strings.ToLower(statusArr[0])
	return status
}

// VM 상태값을 소문자로
func GetVmStatus(vmStatus string) string {
	status := strings.ToLower(vmStatus)
	return status
}

/////////// Map Control 참고

// var mcisIdArr []string
// var vmIdArr []string
// // mcisStatusTotalMap := make(map[string]map[string]int)
// // vmStatusTotalMap := make(map[string]map[string]int)

// mcisStatusRunning := 0
// mcisStatusStopped := 0
// mcisStatusTerminated := 0

// vmStatusRunning := 0
// // vmStatusResuming := 0
// vmStatusInclude := 0
// vmStatusSuspended := 0
// vmStatusTerminated := 0
// vmStatusUndefined := 0
// vmStatusPartial := 0
// vmStatusEtc := 0

// for mcisIndex, mcisInfo := range mcisList {
// 	// log.Println(" mcisInfo  ", index, mcisInfo)
// 	mcisIdArr = append(mcisIdArr, mcisInfo.ID)
// 	mcisStatus := util.GetMcisStatus(mcisInfo.Status)
// 	if mcisStatus == util.MCIS_STATUS_RUNNING {
// 		mcisStatusRunning++
// 	} else if mcisStatus == util.MCIS_STATUS_TERMINATED {
// 		mcisStatusTerminated++
// 	} else {
// 		mcisStatusStopped++
// 	}

// 	vmList := mcisInfo.VMs

// 	for vmIndex, vmInfo := range vmList {
// 		// log.Println(" vmInfo  ", vmIndex, vmInfo)
// 		vmStatus := util.GetVmStatus(vmInfo.Status)
// 		vmIdArr = append(vmIdArr, vmInfo.ID)
// 		if vmStatus == util.VM_STATUS_RUNNING {
// 			vmStatusRunning++
// 			// }else if vmStatus == util.VM_STATUS_RESUMING {
// 			// 	vmStatusResuming++
// 		} else if vmStatus == util.VM_STATUS_INCLUDE {
// 			vmStatusInclude++
// 			// } else if vmStatus == util.VM_STATUS_SUSPENDED {
// 			// 	vmStatusSuspended++
// 		} else if vmStatus == util.VM_STATUS_TERMINATED {
// 			vmStatusTerminated++
// 			// }else if vmStatus == util.VM_STATUS_UNDEFINED {
// 			// 	vmStatusUndefined++
// 			// }else if vmStatus == util.VM_STATUS_PARTIAL {
// 			// 	vmStatusPartial++
// 		} else {
// 			vmStatusEtc++
// 			log.Println("vmStatus  ", vmIndex, vmStatus)
// 		}
// 	}
// 	// vmStatusMap := make(map[string]int)
// 	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌.
// 	// vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
// 	// vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming
// 	// vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
// 	// vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
// 	// vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
// 	// vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
// 	// vmStatusMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
// 	// vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
// 	log.Println("mcisInfo.ID  ", mcisInfo.ID)
// 	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
// 	// vmStatusArr[mcisIndex] = vmStatusMap

// 	// UI에서는 3가지로 통합하여 봄
// 	// vmStatusMap["RUNNING"] = vmStatusRunning
// 	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
// 	// vmStatusMap["TERMINATED"] = vmStatusTerminated
// 	// vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
// 	// vmIdArr = append(vmIdArr, vmInfo.ID)
// 	// vmStatusArr = append(vmStatusArr, vmStatusMap)

// 	log.Println("mcisIndex  ", mcisIndex)
// }
// mcisStatusMap := make(map[string]int)
// mcisStatusMap["RUNNING"] = mcisStatusRunning
// mcisStatusMap["STOPPED"] = mcisStatusStopped
// mcisStatusMap["TERMINATED"] = mcisStatusTerminated
// // mcisStatusTotalMap[mcisInfo.ID] = mcisStatusMap

// vmStatusMap := make(map[string]int)
// vmStatusMap["RUNNING"] = vmStatusRunning
// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
// vmStatusMap["TERMINATED"] = vmStatusTerminated
// // vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
///////////
