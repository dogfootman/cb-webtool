package mcis

import (
	tbmcir "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcir"
)

type TbVmPriority struct {
	Priority string `json:"priority"`

	VmSpec tbmcir.TbSpecInfo `json:"vmSpec"`
}
