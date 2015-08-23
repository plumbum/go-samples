package main

import (
	"fmt"
	"time"
	"log"
	"runtime/pprof"
	"flag"
	"os"
)

const COUNT = 1e10

// Если нужно в цикле изменять переменную от нуля до какого-то значения,
// то конструкция через range не только более наглядная, но и немного
// быстрее классического for
func cycleInterface() (int, int) {
	sum := 0
	cnt := 0
	for i := range [COUNT]interface{}{} {
		sum += i
		cnt++
	}
	return sum, cnt
}

func cycleClassic() (int, int) {
	sum := 0
	cnt := 0
	for i := 0; i < COUNT; i++ {
		sum += i
		cnt++
	}
	return sum, cnt
}

func cycleInt() (int, int) {
	sum := 0
	cnt := 0
	for i := range [COUNT]int{} {
		sum += i
		cnt++
	}
	return sum, cnt
}

func timeDecorator(fn func() (int, int)) (func()) {
	return func() {
		log.Println("== BEGIN =======================================================================")
		timeBegin := time.Now().UnixNano()
		sum, cnt := fn()
		timeEnd := time.Now().UnixNano()
		fmt.Println("Sum:", sum)
		fmt.Println("Count:", cnt)
		fmt.Println("Time:", (timeEnd - timeBegin)/1000, "us")
		log.Println("== END =========================================================================")
	}
}

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
)

type I interface{}

func main() {

	// Profiling http://blog.golang.org/profiling-go-programs

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Предварительное описание пустого интерфейса позволяет подсократить запись цикла
	for i := range [5]I{} {
		fmt.Println(i)
	}

	fmt.Println("Dummy cycle. Wait a few seconds...")

	fmt.Println("Int range for")
	timeDecorator(cycleInt)()
	fmt.Println("Classic for")
	timeDecorator(cycleClassic)()
	fmt.Println("Interface range for")
	timeDecorator(cycleInterface)()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

