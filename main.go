package main

import (
	"fmt"
	"log"
	"os"

	gcp "example.com/gcp-network/gpc"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

func main() {

	ctx := context.Background()

	// Creating client credentials.
	project := "gcpnetwork-349117" // TODO: Update placeholder value.
	jsonPath := "C:/Users/APostescu/Downloads/gcpnetwork-349117-17060bec007f.json"

	// Creating loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// Creating Client
	clientSrv, err := gcp.New(ctx, project, jsonPath, infoLog, errorLog)

	if err != nil {
		log.Fatalf("Creating the client failed %v", err)
	}

	// Create a VPC

	// Network resource
	rb := &compute.Network{
		IPv4Range:                             "",
		AutoCreateSubnetworks:                 false,
		CreationTimestamp:                     "",
		Description:                           "",
		EnableUlaInternalIpv6:                 false,
		GatewayIPv4:                           "",
		Id:                                    0,
		InternalIpv6Range:                     "",
		Kind:                                  "",
		Mtu:                                   1460,
		Name:                                  "pnetwork11231",
		NetworkFirewallPolicyEnforcementOrder: "",
		Peerings:                              []*compute.NetworkPeering{},
		RoutingConfig:                         &compute.NetworkRoutingConfig{RoutingMode: "REGIONAL"},
		SelfLink:                              "",
		SelfLinkWithId:                        "",
		Subnetworks:                           []string{},
		ServerResponse:                        googleapi.ServerResponse{},
		ForceSendFields:                       []string{"AutoCreateSubnetworks"},
		NullFields:                            []string{},
	}
	err = clientSrv.CreateVPC(ctx, rb)
	if err != nil {
		log.Fatalf("Error creating vpc err %v", err)
	}

	// Creating Instance
	zone := "europe-west4-c"
	instanceName := "instance111"
	machineType := "n1-standard-2"
	sourceImage := "projects/debian-cloud/global/images/family/debian-9"
	network := "pnetwork"
	err = clientSrv.CreateInstance(ctx, zone, instanceName, machineType, sourceImage, network, "pnetwork")

	if err != nil {
		log.Fatalf("Error creating instance err %v", err)
	}

	// Stopping an instance
	err = clientSrv.StopInstance(ctx, zone, instanceName)

	if err != nil {
		log.Fatalf("Stopping an instance failed. err %v", err)
	}

	// Getting an instance
	instance, err := clientSrv.GetInstance(ctx, zone, instanceName)
	if err != nil {
		log.Fatalf("Getting an instance failed. err: %v", err)
	}

	// Status of the instance
	fmt.Println(instance.Status)
}
