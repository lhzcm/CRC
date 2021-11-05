package main

import (
	"errors"
	"fmt"
)

func main() {
	var data []byte = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36}
	crc, _ := CRC32(data)
	fmt.Printf("%x\n", crc)
}

func CRC32(data []byte) (uint16, error) {
	var divisor, crc uint16 = 0x8005, 0x0
	var err error = nil
	if len(data) <= 0 {
		err = errors.New("数据不能为空")
		return 0x0, err
	}
	data = append(data, 0x00)
	data = append(data, 0x00)

	crc = uint16(data[0])
	crc = (crc << 8) | uint16(data[1])

	for _, item := range data[2:] {
		for i := 0; i < 8; i++ {
			if crc&0x8000 > 0 {
				crc = crc << 1
				crc = (crc | uint16((item<<i)>>7)) ^ divisor
			} else {
				crc = crc << 1
				crc = (crc | uint16((item<<i)>>7))
			}
		}
	}
	return crc, err
}
