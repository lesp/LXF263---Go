package main

import (
        "fmt"
        "time"
        "os"
        "github.com/stianeikeland/go-rpio"
        "net/http"
)
var (
        red = rpio.Pin(17)
        green = rpio.Pin(27)
)


func main() {
        if err := rpio.Open(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        defer rpio.Close()
        red.Output()
        green.Output()
        red.High()
        green.High()
        time.Sleep(2 * time.Second)
        red.Low()
        green.Low()
        for {
                resp, err := http.Get("http://bigl.es/")
                if resp.StatusCode == 200 {
                        green.High()
                        red.Low()
                        fmt.Println("Server Up\n")
                        time.Sleep(10 * time.Second)
                } else {
                        green.Low()
                        red.High()
                        fmt.Println("Server Down\n")
                        time.Sleep(10 * time.Second)
                }
                fmt.Println(resp.StatusCode,err)
        }
}

