// Code generated by "stringer -type=coil"; DO NOT EDIT.

package plc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[heartbeat-0]
	_ = x[matchReset-1]
	_ = x[stackLightGreen-2]
	_ = x[stackLightOrange-3]
	_ = x[stackLightRed-4]
	_ = x[stackLightBlue-5]
	_ = x[stackLightBuzzer-6]
	_ = x[fieldResetLight-7]
	_ = x[cargoShipMagnetRed-8]
	_ = x[cargoShipMagnetBlue-9]
	_ = x[cargoShipLightRed-10]
	_ = x[cargoShipLightBlue-11]
	_ = x[sandstormUpRed-12]
	_ = x[sandstormUpBlue-13]
	_ = x[rocketLightRedNear-14]
	_ = x[rocketLightRedFar-15]
	_ = x[rocketLightBlueNear-16]
	_ = x[rocketLightBlueFar-17]
	_ = x[coilCount-18]
}

const _coil_name = "heartbeatmatchResetstackLightGreenstackLightOrangestackLightRedstackLightBluestackLightBuzzerfieldResetLightcargoShipMagnetRedcargoShipMagnetBluecargoShipLightRedcargoShipLightBluesandstormUpRedsandstormUpBluerocketLightRedNearrocketLightRedFarrocketLightBlueNearrocketLightBlueFarcoilCount"

var _coil_index = [...]uint16{0, 9, 19, 34, 50, 63, 77, 93, 108, 126, 145, 162, 180, 194, 209, 227, 244, 263, 281, 290}

func (i coil) String() string {
	if i < 0 || i >= coil(len(_coil_index)-1) {
		return "coil(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _coil_name[_coil_index[i]:_coil_index[i+1]]
}
