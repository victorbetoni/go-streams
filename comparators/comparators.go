package comparators

func compareRealNum(i1, i2 float64) int {
	if i1 < i2 {
		return -1
	}
	if i1 == i2 {
		return 0
	}
	return 1
}

func CompareInt(i1, i2 int) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareInt8(i1, i2 int8) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareInt32(i1, i2 int32) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareInt64(i1, i2 int64) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareUint(i1, i2 uint) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareUint8(i1, i2 uint8) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareUint32(i1, i2 uint32) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareUint64(i1, i2 uint64) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareUintptr(i1, i2 uintptr) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareFloat64(i1, i2 float64) int {
	return compareRealNum(float64(i1), float64(i2))
}

func CompareFloat32(i1, i2 float32) int {
	return compareRealNum(float64(i1), float64(i2))
}
