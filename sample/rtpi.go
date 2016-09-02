package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zoltanszekely/rtpi-go/rtpi"
	"github.com/zoltanszekely/rtpi-go/rtpi/types"
)

func str(s *string) string {
	if s == nil {
		return "n/a"
	}
	return *s
}

func handleResultError(result *types.Result) {
	if result.ErrorCode != nil && *result.ErrorCode != 0 {
		fmt.Println("error:")
		fmt.Printf("\tcode: %#v\n", *result.ErrorCode)
		fmt.Printf("\tmessage: %#v\n", str(result.ErrorMessage))
	}
}

func main() {
	stopID := flag.String("stop", "", "stop id")
	listStops := flag.Bool("listStops", false, "list all stops")
	flag.Parse()
	if (*listStops == true && *stopID != "") || (*listStops == false && *stopID == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := rtpi.Client{}

	if *listStops {
		result, err := client.GetBusStopInformation()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		handleResultError(&result.Result)
		if result.Results != nil {
			fmt.Println("results:")
			for _, result := range *result.Results {
				fmt.Printf("\tstop: %#v, short name: %#v, full name: %#v\n", str(result.StopID), str(result.ShortName), str(result.FullName))
			}
		}

	} else {
		result, err := client.GetRealTimeBusInformation(*stopID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		handleResultError(&result.Result)
		if result.Results != nil {
			fmt.Printf("stop: %#v\n", str(result.StopID))
			fmt.Println("results:")
			for _, result := range *result.Results {
				fmt.Printf("\tdue: %#v, route: %#v, origin: %#v, destination: %#v\n", str(result.DueTime), str(result.Route), str(result.Origin), str(result.Destination))
			}
		}
	}
}
