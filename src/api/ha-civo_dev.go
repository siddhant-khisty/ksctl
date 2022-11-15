//
// ######################################
//		EARLY DEVELOPMENT (TEMPLATE)
//   NOTE: THIS FILE WILL BE REMOVED IN NEXT VERSION (remane it to main.go_template
// ######################################
//

package main

import "github.com/kubesimplify/ksctl/src/api/ha_civo"

func main() {
	ha_civo.CreateVM("control-plane-1", "LON1")
  //ha_civo.CreateVM("control-plane-2", "LON1")
	//ha_civo.CreateVM("control-plane-3", "LON1")
}
