package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInt(device string) int {
	_, err := os.Stat(device)
	if err != nil {
		fmt.Println("Error while getting stats of", device, " - ", device)
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(device)
	if err != nil {
		fmt.Println("Error while reading", device, " - ", device)
		os.Exit(1)
	}

	value, err := strconv.Atoi(strings.Trim(string(data), "\n"))
	if err != nil {
		fmt.Println("Error parsing value for", device, " - ", err)
		os.Exit(1)
	}

	return value
}

func writeInt(device string, n int) {
	data := []byte(strconv.Itoa(n) + "\n")
	err := ioutil.WriteFile(device, data, os.ModeCharDevice)
	if err != nil {
		fmt.Println("Error while writing", n, "to", device, " - ", err)
		os.Exit(1)
	}
}

func changeDeviceByN(device string, n int) {
	currentBrightness := readInt(device + "/brightness")
	maxBrightness := readInt(device + "/max_brightness")

	if maxBrightness <= 0 {
		fmt.Println("Found max_brightness with value=0. This doesn't seem correct.")
		os.Exit(1)
	}

	changeBy := (float64(n) / float64(100)) * float64(maxBrightness)
	newBrightness := float64(currentBrightness) + changeBy
	newBrightness = math.Max(newBrightness, 0)
	newBrightness = math.Min(newBrightness, float64(maxBrightness))

	writeInt(device+"/brightness", int(newBrightness))
}

func main() {
	devicePtr := flag.String("device", "", "The options are 'keyboard' or 'display' or '/sys/class/etcetcetc'.")
	amountPtr := flag.Int("amount", 5, "Increate or decreate by this percentage. Ex: 5 or -5.")
	flag.Parse()

	switch *devicePtr {
	case "":
		fmt.Println("No device specified:")
		flag.PrintDefaults()
	case "keyboard":
		changeDeviceByN("/sys/class/leds/smc::kbd_backlight/", *amountPtr)
	case "display":
		changeDeviceByN("/sys/class/backlight/gmux_backlight/", *amountPtr)
	default:
		devicePath := fmt.Sprint("/sys/class/", *devicePtr, "/")
		if _, err := os.Stat(devicePath); os.IsNotExist(err) {
			fmt.Println("Missing device for path:", devicePath)
			os.Exit(1)
		}

		changeDeviceByN(devicePath, *amountPtr)
	}
}
