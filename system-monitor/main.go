package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"system-meterics/models"
	"time"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func checkerr(err error) {

	if err != nil {
		fmt.Println("Error Happened", err)
		os.Exit(1)
	}
}

func main() {

	fmt.Println("Server is starting on port 4000....")

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	var system_meterics models.System_meterics

	name, os, arch := gethostInfo()
	cores, percent := getCpuInfo()
	totalmem, freemem, memused := getmemInfo()

	system_meterics = models.System_meterics{name, totalmem, freemem, memused, arch, os, cores, percent}
	byteinfo, err := json.Marshal(system_meterics)
	checkerr(err)
	w.Write(byteinfo)

}

func getmemInfo() (uint64, uint64, float64) {

	v, err := mem.VirtualMemory()

	checkerr(err)

	total_memory := v.Total
	free_memory := v.Free
	memory_used_percentage := v.UsedPercent

	return total_memory, free_memory, memory_used_percentage

}

func getCpuInfo() (int, float64) {

	cores, _ := cpu.Counts(true)
	percent, _ := cpu.Percent(time.Second, false)

	return cores, percent[0]
}

func gethostInfo() (string, string, string) {

	arch, _ := host.KernelArch()
	hostinfo, _ := host.Info()

	name := hostinfo.Hostname
	os := hostinfo.OS

	return name, os, arch
}
