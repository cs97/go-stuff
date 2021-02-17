package main
import (
    "fmt"
    "net"
)
func main() {
    for i := 1; i <= 65555; i++ {
        address := fmt.Sprintf("127.0.0.1:%d", i)
        conn, err := net.Dial("tcp", address)
        if err != nil {
            // port is closed or filtered.
            continue
        }
        conn.Close()
        fmt.Printf("%d open\n", i)
    }
}

