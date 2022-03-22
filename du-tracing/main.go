package main

import (
	"fmt"
	"io"
	"os"
	"runtime/trace"

	"github.com/Wei-N-Ning/gotypes/pkg/iterator"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator/fs"
)

func addTwo(lhs int64, rhs int64) int64 {
	return lhs + rhs
}

func getSizeFast(item fs.Item) int64 {
	if info, err := item.DirEntry.Info(); err != nil {
		return 0
	} else {
		return info.Size()
	}
}

func getSizeSlow(item fs.Item) int64 {
	if item.DirEntry.IsDir() {
		return 0
	} else {
		r, err := os.Open(item.Path)
		if err != nil {
			return 0
		}
		bs := make([]byte, 128)
		totalRead := 0
		for {
			numRead, err := r.Read(bs)
			if err == io.EOF {
				break
			}
			totalRead += numRead
		}
		return int64(totalRead)
	}
}

const (
	Options = `The valid choices are:
	--- fast-IO ---
	f1: serial for-each
	f2: serial map-reduce
	f3: parallel map (ordered) then reduce
	f4: parallel map (unordered) then reduce
	f5: parallel map-reduce
	
	--- slow-IO ---
	s1: serial for-each
	s2: serial map-reduce
	s3: parallel map (ordered) then reduce
	s4: parallel map (unordered) then reduce
	s5: parallel map-reduce
`
)

func main() {
	args := os.Args[1:]
	dirPath := "."
	if len(args) > 0 {
		dirPath = args[0]
	}
	opt := ""
	if len(args) > 1 {
		opt = args[1]
	}

	var x int64

	func() {
		err := trace.Start(os.Stderr)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()

		switch opt {
		case "f1":
			fs.DirIter(dirPath).ForEach(func(item fs.Item) { x += getSizeFast(item) })
		case "f2":
			x = iterator.MapReduce(fs.DirIter(dirPath), 0, getSizeFast, addTwo)
		case "f3":
			x = iterator.ParMap(fs.DirIter(dirPath), func(item fs.Item) int64 { return getSizeFast(item) }).Reduce(0, addTwo)
		case "f4":
			x = iterator.ParMapUnord(fs.DirIter(dirPath), func(item fs.Item) int64 { return getSizeFast(item) }).Reduce(0, addTwo)
		case "f5":
			x = iterator.ParMapReduce(fs.DirIter(dirPath), 0, getSizeFast, addTwo)
		case "s1":
			fs.DirIter(dirPath).ForEach(func(item fs.Item) { x += getSizeSlow(item) })
		case "s2":
			x = iterator.MapReduce(fs.DirIter(dirPath), 0, getSizeSlow, addTwo)
		case "s3":
			x = iterator.ParMap(fs.DirIter(dirPath), func(item fs.Item) int64 { return getSizeSlow(item) }).Reduce(0, addTwo)
		case "s4":
			x = iterator.ParMapUnord(fs.DirIter(dirPath), func(item fs.Item) int64 { return getSizeSlow(item) }).Reduce(0, addTwo)
		case "s5":
			x = iterator.ParMapReduce(fs.DirIter(dirPath), 0, getSizeSlow, addTwo)
		default:
			fmt.Println("Unsupported option! " + Options)
			os.Exit(1)
		}
	}()
	fmt.Println(x/(1024*1024), "M")
}
