package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/disk"
)

func main() {

	d := flag.String("disk", "/", "Disk name.")
	warning := flag.Int("warning", 80, "Warning thresold for disk.")
	critical := flag.Int("critical", 90, "Critical thresold for disk.")
	flag.Parse()

	u, err := disk.Usage(*d)
	if err != nil {
		fmt.Printf("Can't get information - error %v.\n", err)
		os.Exit(3)
	}

	total := u.Total / 1024 / 1024

	used := u.Used / 1024 / 1024

	if int(u.UsedPercent) >= *critical {
		fmt.Printf("CRITICAL - Disk %v percent usage %v%% - Total Disk %v Mo - Used Disk %v Mo | disk_percent = %v,%v,%v,0,100\n", *d, int(u.UsedPercent), total, used, int(u.UsedPercent), *warning, *critical)
		os.Exit(2)

	} else if int(u.UsedPercent) >= *warning {
		fmt.Printf("WARNING - Disk %v percent usage %v%% - Total Disk %v Mo - Used Disk %v Mo | disk_percent = %v,%v,%v,0,100\n", *d, int(u.UsedPercent), total, used, int(u.UsedPercent), *warning, *critical)
		os.Exit(1)

	} else {
		fmt.Printf("OK - Disk %v percent usage %v%% - Total Disk %v Mo - Used Disk %v Mo | disk_percent = %v,%v,%v,0,100\n", *d, int(u.UsedPercent), total, used, int(u.UsedPercent), *warning, *critical)
		os.Exit(0)
	}

}
