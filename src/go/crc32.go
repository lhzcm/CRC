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

func CRC32(data []byte) (uint32, error) {
	var divisor, crc uint32 = 0x04C11DB7, 0x0
	var err error = nil
	if len(data) <= 0 {
		err = errors.New("数据不能为空")
		return 0x0, err
	}
	data = append(data, 0x00)
	data = append(data, 0x00)
	data = append(data, 0x00)
	data = append(data, 0x00)

	crc = uint32(data[0])
	crc = (crc << 8) | uint32(data[1])
	crc = (crc << 8) | uint32(data[2])
	crc = (crc << 8) | uint32(data[3])

	for _, item := range data[4:] {
		for i := 0; i < 8; i++ {
			if crc&0x80000000 > 0 {
				crc = crc << 1
				crc = (crc | uint32((item<<i)>>7)) ^ divisor
			} else {
				crc = crc << 1
				crc = (crc | uint32((item<<i)>>7))
			}
		}
	}
	return crc, err
}
