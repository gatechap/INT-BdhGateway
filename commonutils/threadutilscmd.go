package commonutils

import (
	"errors"
	"math"

	"intbackend/bdhgateway/commonmodel"
)

func CalculateThreadInfo(maxthread int, arraysize int) (*commonmodel.ThreadInfo, *error) {
	if maxthread == 0 {
		err := errors.New("maxthread cannot be zero")
		return nil, &err
	}
	if arraysize == 0 {
		err := errors.New("arraysize cannot be zero")
		return nil, &err
	}

	floor64 := math.Floor(float64(arraysize) / float64(maxthread))
	if floor64 >= math.MaxInt64 || floor64 <= math.MinInt64 {
		err := errors.New("floor function arraysize divide by maxthread is out of int64 range")
		return nil, &err
	}

	iFloor := int(floor64)
	iMod := int(math.Mod(float64(arraysize), float64(maxthread)))
	iThreadLoop := iFloor
	if iMod > 0 {
		iThreadLoop = iThreadLoop + 1
	}
	iFullElementLoop := iFloor
	iElementNoneFullLoop := iMod

	res := commonmodel.ThreadInfo{
		MaxThread:                 maxthread,
		ArraySize:                 arraysize,
		NumbersOfThreadLoop:       iThreadLoop,
		NumbersOfFullElementLoop:  iFullElementLoop,
		LastElementOfNoneFullLoop: iElementNoneFullLoop,
	}

	return &res, nil
}
