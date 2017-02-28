//This is a example to get the cpudinfo data from pqos
// we need to add testcase for all rdtagent go code

package main

import (
    "fmt"
    "cgolib/cpuinfo"
)

//func main() {
func test() {
    _ = "breakpoint"
    // pq := new(cpuinfo.Pqos)
    pq, _ := cpuinfo.GetCpuInfo()
    fmt.Println("CPU info: ", pq)
}
