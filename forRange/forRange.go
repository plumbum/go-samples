package main

// Use Go 1.5

import (
	"fmt"
	"time"
	"flag"
	"sync"
	"os"
	"log"
	"runtime/pprof"
	"runtime"
)

const COUNT = 2e9

// Если нужно в цикле изменять переменную от нуля до какого-то значения,
// то конструкция через range не только более наглядная, но и немного
// быстрее классического for
func cycleInterface() (uint64, uint64) {
	// Predefined variable types
	var sum uint64 = 0
	var cnt uint64 = 0
	var i int
	for i = range [COUNT]interface{}{} {
		sum += (uint64)(i)
		cnt++
	}
	return sum, cnt
}

func cycleStruct() (uint64, uint64) {
	// Predefined variable types
	var sum uint64 = 0
	var cnt uint64 = 0
	var i int
	for i = range [COUNT]struct{}{} {
		sum += (uint64)(i)
		cnt++
	}
	return sum, cnt
}

// А вот перебор с классической формой цикла почему-то медленнее
func cycleClassic() (uint64, uint64) {
	// Predefined variable types
	var sum uint64 = 0
	var cnt uint64 = 0
	var i int
	for i = 0; i < COUNT; i++ {
		sum += (uint64)(i)
		cnt++
	}
	return sum, cnt
}

func cycleInt() (uint64, uint64) {
	// Predefined variable types
	var sum uint64 = 0
	var cnt uint64 = 0
	var i int
	for i = range [COUNT]int{} {
		sum += (uint64)(i)
		cnt++
	}
	return sum, cnt
}

func timeDecorator(fn func() (uint64, uint64), name string, wg *sync.WaitGroup) (func()) {
	return func() {
		fmt.Println(name, "== BEGIN =======================================================================")
		runtime.GC()
		start := time.Now()
		sum, cnt := fn()
		fmt.Println(name, "Time:", time.Since(start)) // Сколько прошло времени. Это удобнее чем руками фиксировать конечное время
		fmt.Println(name, "Sum:", sum)
		fmt.Println(name, "Count:", cnt)
		fmt.Println(name, "== END =========================================================================")
		if wg != nil {
			wg.Done()
		}
	}
}

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	useRoutines = flag.Bool("routines", false, "Use Go Routines for work")
	testNum = flag.Int("test", -1, "0 - classic, 1 - int, 2 - interface, 3 - struct")
)

type I interface{}

func main() {

	// Предварительное описание пустого интерфейса позволяет подсократить запись цикла
	for i := range [5]I{} {
		fmt.Println(i)
	}


	// Set number of CPU usage
	runtime.GOMAXPROCS(runtime.NumCPU())

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

	start := time.Now()

	if *testNum >= 0 {
		switch *testNum {
		case 0:
			timeDecorator(cycleClassic,   "Classic  :", nil)()
		case 1:
			timeDecorator(cycleInt,       "Int      :", nil)()
		case 2:
			timeDecorator(cycleInterface, "Interface:", nil)()
		case 3:
			timeDecorator(cycleStruct,    "Struct   :", nil)()
		}
	} else {
		if ! *useRoutines {
			timeDecorator(cycleStruct,    "Struct   :", nil)()
			timeDecorator(cycleInterface, "Interface:", nil)()
			timeDecorator(cycleInt,       "Int      :", nil)()
			timeDecorator(cycleClassic,   "Classic  :", nil)()
			timeDecorator(cycleInterface, "Interface:", nil)()
			timeDecorator(cycleStruct,    "Struct   :", nil)()
			timeDecorator(cycleClassic,   "Classic  :", nil)()
			timeDecorator(cycleInt,       "Int      :", nil)()
		} else {
			var wg sync.WaitGroup

			wg.Add(1)
			go timeDecorator(cycleInt,       "Int      :", &wg)()

			wg.Add(1)
			go timeDecorator(cycleClassic,   "Classic  :", &wg)()

			wg.Add(1)
			go timeDecorator(cycleInterface, "Interface:", &wg)()

			wg.Add(1)
			go timeDecorator(cycleStruct,    "Struct   :", &wg)()

			wg.Wait()
		}
	}

	fmt.Println("Total time:", time.Since(start))

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

