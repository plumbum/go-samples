package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "os"
    "github.com/jpillora/overseer"
    "github.com/jpillora/overseer/fetcher"
    "strings"
)

var buildDate = "not set"

//create another main() to run the overseer process
//and then convert your old main() into a 'prog(state)'
func main() {

    log.Println("Main: start PID: ", os.Getpid(), os.Getppid())
    for _, env := range os.Environ() {
        if strings.HasPrefix(env, "GO") {
            log.Println("Main env: ", env)
        }
    }

    // Следует иметь в виду, что oerseer запускает два процесса.
    // Первый - процесс мониторинга занимается проверкой обновлений и принимает запросы
    // Второй - собственно процесс обработки запросов. Он перезапускается при обновлении
    // или если первый процесс получает сигнал SIGUSR2 (по умолчанию, может быть переопределён)
    // Overseer не является менеджером процессов. Если второй процесс завершится по каким-то причинам,
    // то будет завершен и первый.
    overseer.Run(overseer.Config{
        // Debug: true,
        Program: prog,
        PreUpgrade: checkBinary,
        Address: ":3000",
        Fetcher: &fetcher.HTTP{
            // Для тестирования в отдаче файла лучше всего воспользоваться `webfsd -F -p 4000 -l -`
            // Он отдаёт корректные заголовки при изменении файла
            URL:      "http://localhost:4000/overseer-test",
            Interval: 5 * time.Second,
        },
    })
}

// Фукнция вызывается после получения файла обновления по сети и позволяет проверить его корректность.
func checkBinary(tmpFilePath string) error {
    log.Println("Ready to upgrade from file: ", tmpFilePath)
    fi, err := os.Stat(tmpFilePath)
    if err != nil {
        return err
    }
    log.Println("File size: ", fi.Size())

    return nil
}

var counter = 0

//prog(state) runs in a child process
func prog(state overseer.State) {
    // pp.Println(state)
    log.Println("Prog: start PID: ", os.Getpid(), os.Getppid())
    log.Println("Building date: ", buildDate)
    log.Printf("app (%s) listening at %s...", state.ID, state.Listener.Addr())
    http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Get handle from ", r.RemoteAddr, r.RequestURI)
        log.Println("PID=", os.Getpid())
        fmt.Fprintln(w, "<h1>Test overseer server</h1>")
        fmt.Fprintln(w, `<a href="https://github.com/jpillora/overseer">Overseer home page</a>`)
        fmt.Fprintf(w, "<p>Build date: %s</p>", buildDate)
        counter++
        fmt.Fprintf(w, "<p>My app (%s) says hello %d times</p>\n", state.ID, counter)
        fmt.Fprintf(w, "<p>PID=%d, PPID=%d</p>\n", os.Getpid(), os.Getppid())
        fmt.Fprintf(w, "<p>Application: %s</p>\n", os.Args[0])

        fmt.Fprintf(w, "<hr/><p>Args: %v</p>", os.Args[1:])
        fmt.Fprintln(w, "<hr/><ul>\n")
        for _, env := range os.Environ() {
            fmt.Fprintf(w, "<li>Env: %v</li>\n", env)
        }
        fmt.Fprintln(w, "</ul>\n")

    }))
    http.Serve(state.Listener, nil)
    log.Println("Stop server ", os.Getpid())
}
