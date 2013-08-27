package gophercloud

import (
	"github.com/racker/perigee"
	"fmt"
)

func (gcp *genericServersProvider) ListVolumes() ([]Volume, error) {
	var vols []Volume

	err := gcp.context.WithReauth(gcp.access, func() error {
		url := gcp.endpoint + "/volumes"
		fmt.Println(url)
		return perigee.Get(url, perigee.Options{
			CustomClient: gcp.context.httpClient,
			Results:      &struct{ Volumes *[]Volume }{&vols},
			MoreHeaders: map[string]string{
				"X-Auth-Token": gcp.access.AuthToken(),
			},
		})
	})
	return vols, err
}


func (gcp *genericServersProvider) ShowVolumes() ([]Volume, error) {
	var vols []Volume

	err := gcp.context.WithReauth(gcp.access, func() error {
		url := gcp.endpoint + "/volumes"
		fmt.Println(url)
		return perigee.Get(url, perigee.Options{
			CustomClient: gcp.context.httpClient,
			Results:      &struct{ Volumes *[]Volume }{&vols},
			MoreHeaders: map[string]string{
				"X-Auth-Token": gcp.access.AuthToken(),
			},
		})
	})
	return vols, err
}

type Volume struct {
	Id	string	`json:"id"`
	Display_name	string	`json:"display_name"`
	Display_description	string	`json:"display_description"`
	Status	string	`json:"status"`
	Size	string	`json:"size"`
	Volume_type	string	`json:"volume_type"`
	CreatedAt	string	`json:"createdAt"`
}
