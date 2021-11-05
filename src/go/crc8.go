package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	//var data []byte = []byte{0x31, 0x32, 0x33, 0x34, 0x35}

	var data []byte

	file, err := os.OpenFile("testfile.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败！")
		return
	}

	data, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败！")
		return
	}

	fmt.Println("----------方法1------------")
	startTime := time.Now()
	var result byte = 0x00
	for i := 0; i <= 10; i++ {
		result, _ = GetCRC8(data)
	}
	elapsedTime := time.Since(startTime) /// time.Millisecond
	fmt.Printf("CRC8=%#v\n", result)
	fmt.Println(elapsedTime)

	fmt.Println("----------方法2------------")
	startTime = time.Now()
	for i := 0; i <= 10; i++ {
		result, _ = GetCRC8(data)
	}
	elapsedTime = time.Since(startTime) /// time.Millisecond
	fmt.Printf("CRC8=%#v\n", result)
	fmt.Println(elapsedTime)
}

func GetCRC8(data []byte) (byte, error) {
	var divisor, crc byte = 0x07, 0x00
	var err error = nil
	if len(data) <= 0 {
		err = errors.New("数据不能为空")
		return 0x0, err
	}
	data = append(data, 0x00)
	crc = data[0]
	for _, item := range data[1:] {
		for i := 0; i < 8; i++ {
			if crc&0x80 > 0 {
				crc = crc << 1
				crc = (crc | ((item << i) >> 7)) ^ divisor
			} else {
				crc = crc << 1
				crc = (crc | ((item << i) >> 7))
			}
		}
	}
	return crc, err
}

func GetCRC8_Method2(data []byte) (byte, error) {
	var divisor, crc byte = 0x07, 0x00
	var err error = nil
	if len(data) <= 0 {
		err = errors.New("数据不能为空")
		return 0x0, err
	}
	data = append(data, 0x00)
	crc = data[0]
	for _, item := range data[1:] {
		for i := 0; i < 8; i++ {
			if crc&0x80 > 0 {
				crc = crc << 1
				crc = (crc | (item >> (7 - i))) ^ divisor
				item = (item << (i + 1)) >> (i + 1)
			} else {
				crc = crc << 1
			}
		}
		crc = crc | item
	}
	return crc, err
}
