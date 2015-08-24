package main

import (
	"fmt"
	"time"
	"log"
	"runtime/pprof"
	"flag"
	"os"
	"sync"
	"runtime"
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

func timeDecorator(fn func() (int, int), name string, wg *sync.WaitGroup) (func()) {
	return func() {
		fmt.Println(name, "== BEGIN =======================================================================")
		start := time.Now()
		sum, cnt := fn()
		fmt.Println(name, "Time:", time.Since(start)) // Сколько прошло времени. Это удобнее чем руками фиксировать конечное время
		fmt.Println(name, "Sum:", sum)
		fmt.Println(name, "Count:", cnt)
		fmt.Println(name, "== END =========================================================================")
		wg.Done()
	}
}

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
)

type I interface{}

func main() {

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

	// Предварительное описание пустого интерфейса позволяет подсократить запись цикла
	for i := range [5]I{} {
		fmt.Println(i)
	}

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(1)
	go timeDecorator(cycleInt,       "Int      :", &wg)()

	wg.Add(1)
	go timeDecorator(cycleClassic,   "Classic  :", &wg)()

	wg.Add(1)
	go timeDecorator(cycleInterface, "Interface:", &wg)()

	wg.Wait()

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

