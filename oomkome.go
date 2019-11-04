package main

import (
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: oomkome <memory_size>")
		println("\tmemory_size can be a number of bytes, or be in form of xK, xM, xG")
		return
	}
	mul := 1.0
	mem := os.Args[1]

	switch mem[len(mem)-1] {
	case 'K':
		mul = 1024
		mem = mem[0 : len(mem)-1]
	case 'M':
		mul = 1024 * 1024
		mem = mem[0 : len(mem)-1]
	case 'G':
		mul = 1024 * 1024 * 1024
		mem = mem[0 : len(mem)-1]
	}

	fsize, err := strconv.ParseFloat(mem, 10)
	if err != nil {
		println("Failed to parse", os.Args[1], ":", err.Error())
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				println(err.Error())
			}
		}
	}()

	size := int(mul * fsize)
	println("Eating bytes:", size)
	start := time.Now()
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = uint8(i % 255)
	}

	elapsed := time.Now().Sub(start)
	println("Done in", elapsed.String())
	println("Filling speed: ", int64(float64(size)/(elapsed.Seconds()*1024*1024)), "MB/s")

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	println("System heap:", ms.HeapSys/(1024*1024), "MB")
	println("Sleeping 10s")
	time.Sleep(10 * time.Second)
}
