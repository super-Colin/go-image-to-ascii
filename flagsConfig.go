package main

import "fmt"

func flagConfig() {
	var flags map[string]*scoFlag

	colorCloseToReq := MakeIntFlag("cctr", "The distance requirement between colors for them to be blended", 80, 0, 255)
	var colorCloseToReqFlag scoFlag = &colorCloseToReq
	flags[colorCloseToReq.flagName] = &colorCloseToReqFlag

	colorDistanceReq := MakeIntFlag("cdr", "The distance requirement between colors for them to be distinct", 40, 0, 255)
	var colorDistanceReqFlag scoFlag = &colorDistanceReq
	flags[colorDistanceReq.flagName] = &colorDistanceReqFlag

	// Flags to export
	fmt.Println("flags: ", flags)

}
