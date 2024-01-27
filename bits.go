package bits

func CheckBit(value byte, bit uint8) bool {
	return (value & (1 << bit)) != 0
}

func SetBit(value byte, bit uint8) byte {
	return value | (1 << bit)
}

func ClearBit(value byte, bit uint8) byte {
	return value &^ (1 << bit)
}

func GetBit(value byte, bit uint8) byte {
	return value & (1 << bit)
}
