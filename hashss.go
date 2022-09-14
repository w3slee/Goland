package main
import (
    "fmt"
    "hash/crc32"
)

/* 
    A common use for crc32 is to compare two files
    If the Sum32 value for both files is the same, it's highly likely (
*/ 
func main() {
    h := crc32.NewIEEE()
    h.Write([]byte("test"))
    v := h.Sum32()
    fmt.Println(v)
}
