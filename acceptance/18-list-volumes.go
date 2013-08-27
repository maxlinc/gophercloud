package main

import (
	"fmt"
	"flag"
	"github.com/rackspace/gophercloud"
)

var quiet = flag.Bool("quiet", false, "Quiet mode, for acceptance testing.  $? still indicates errors though.")

func main() {
	flag.Parse()
	withIdentity(false, func(acc gophercloud.AccessProvider) {
		withBlockStorageApi(acc, func(servers gophercloud.CloudServersProvider) {

			log("Listing Volumes")

			vols, err := servers.ListVolumes()

			if err != nil {
				panic(err)
			}

			if len(vols) < 1 {
				panic("no volumes listed")
			} else {
				for _,v:=range vols{
					log(fmt.Sprintf("volume name: %s", v.Display_name))
				}
			}
			log("Done Listing Volumes")
		})
	})
}

func log(s string) {
	if !*quiet {
		fmt.Println(s)
	}
}
