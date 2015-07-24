package main

import (
	"fmt"
	"github.com/profitbricks/profitbricks-sdk-go"
	"time"
)

func main() {

	//Sets username and password
	profitbricks.SetAuth("username", "password")
	//Sets depth.
	profitbricks.SetDepth("5")

	dcrequest := profitbricks.CreateDatacenterRequest{
		DCProperties: profitbricks.DCProperties{
			Name:        "example.go3",
			Description: "description",
			Location:    "us/lasdev",
		},
	}

	datacenter := profitbricks.CreateDatacenter(dcrequest)

	serverrequest := profitbricks.CreateServerRequest{
		ServerProperties: profitbricks.ServerProperties{
			Name:  "go01",
			Ram:   1024,
			Cores: 2,
		},
	}
	server := profitbricks.CreateServer(datacenter.Id, serverrequest)

	/*images := profitbricks.ListImages()

	fmt.Println(images.Items)
	*/
	volumerequest := profitbricks.CreateVolumeRequest{
		VolumeProperties: profitbricks.VolumeProperties{
			Size:        1,
			Name:        "Volume Test",
			LicenceType: "LINUX",
		},
	}

	storage := profitbricks.CreateVolume(datacenter.Id, volumerequest)

	serverupdaterequest := profitbricks.ServerProperties{
		Name:  "go01renamed",
		Cores: 1,
		Ram:   256,
	}

	profitbricks.PatchServer(datacenter.Id, server.Id, serverupdaterequest)
	//It takes a moment for a volume to be provisioned so we wait.
	time.Sleep(60 * time.Second)
	
	profitbricks.AttachVolume(datacenter.Id, server.Id, storage.Id)

	volumes := profitbricks.ListVolumes(datacenter.Id)
	fmt.Println(volumes.Items)
	servers := profitbricks.ListServers(datacenter.Id)
	fmt.Println(servers.Items)
	datacenters := profitbricks.ListDatacenters()
	fmt.Println(datacenters.Items)

	profitbricks.DeleteServer(datacenter.Id, server.Id)
	profitbricks.DeleteDatacenter(datacenter.Id)
}