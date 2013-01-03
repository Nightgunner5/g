// +build profile

package main

import (
	"fmt"
	"runtime/pprof"
	"os"
)

func init() {
	// We drop the file handle here without closing it, which is normally
	// bad practice, but it's a one-time thing and it wouldn't be closed
	// until the end of the program's execution anyway. pls dnt dailyWTF
	f, err := os.Create(fmt.Sprintf("cpu.%d.prof", os.Getpid()))
	if err != nil {
		panic(err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	cleanup = append(cleanup, pprof.StopCPUProfile)

	cleanup = append(cleanup, memProf)
}

func memProf() {
	f, err := os.Create(fmt.Sprintf("heap.%d.prof", os.Getpid()))
	if err != nil {
		panic(err)
	}

	err = pprof.Lookup("heap").WriteTo(f, 0)
	f.Close()
	if err != nil {
		panic(err)
	}
}
