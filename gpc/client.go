package gcp

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type Client struct {
	// Struct from Google Library
	Compute *compute.Service

	// Required user input
	ProjectID string
	JSONPath  string
	// Loggers
	infoLog  *log.Logger
	errorLog *log.Logger
}

func New(ctx context.Context, project, jsonPath string, infoLog, errorLog *log.Logger) (*Client, error) {
	srv, err := compute.NewService(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Printf("err:%v\n", err)
		return nil, err
	}
	return &Client{
		srv,
		project,
		jsonPath,
		infoLog,
		errorLog,
	}, nil
}

func (c *Client) CreateVPC(ctx context.Context, vpc *compute.Network) error {
	c.infoLog.Println("Creating VPC...")
	oper, err := c.Compute.Networks.Insert(c.ProjectID, vpc).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Create VPC err : %v\n", err)
		return err
	}
	_, err = c.Compute.GlobalOperations.Wait(c.ProjectID, oper.Name).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Waiting for VPC creation err")
	}
	return err
}

func (c *Client) GetVPC(ctx context.Context, vpcName string) (*compute.Network, error) {
	resp, err := c.Compute.Networks.Get(c.ProjectID, vpcName).Context(ctx).Do()
	if err != nil {
		fmt.Printf("Get VPC err: %v\n", err)
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateInstance(ctx context.Context, zone, instanceName, machineType, sourceImage, network, subnetwork string) error {
	c.infoLog.Println("Creating instance ... ")
	prefix := "https://www.googleapis.com/compute/v1/projects/" + c.ProjectID
	region := "europe-west4"

	computeInstance := &compute.Instance{
		Name:        instanceName,
		Description: "An instance created with CreateInstance",
		MachineType: prefix + "/zones/" + zone + "/machineTypes/" + machineType,
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTANT",
				InitializeParams: &compute.AttachedDiskInitializeParams{

					SourceImage: sourceImage,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				Network:    prefix + "/global/networks/" + network,
				Subnetwork: prefix + "/regions/" + region + "/subnetworks/" + subnetwork,
			},
		},
	}

	oper, err := c.Compute.Instances.Insert(c.ProjectID, zone, computeInstance).Do()
	if err != nil {
		c.errorLog.Printf("Insert instance error:%v", err)
		return err
	}
	_, err = c.Compute.ZoneOperations.Wait(c.ProjectID, zone, oper.Name).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Waiting for insert instance oper error:%v", err)
		return err
	}
	return nil
}

func (c *Client) GetInstance(ctx context.Context, zone, instance string) (*compute.Instance, error) {
	resp, err := c.Compute.Instances.Get(c.ProjectID, zone, instance).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Get instance error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *Client) DeleteInstance(ctx context.Context, zone, instance string) error {
	oper, err := c.Compute.Instances.Delete(c.ProjectID, zone, instance).Do()
	if err != nil {
		c.errorLog.Printf("Delete instance error: %v", err)
		return err
	}
	_, err = c.Compute.ZoneOperations.Wait(c.ProjectID, zone, oper.Name).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Waiting for delete instance oper error: %v", err)
		return err
	}
	return nil
}

func (c *Client) StartInstance(ctx context.Context, zone, instance string) error {
	c.infoLog.Println("Starting an instance...")
	oper, err := c.Compute.Instances.Start(c.ProjectID, zone, instance).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Start instance error: %v\n", err)
		return err
	}
	_, err = c.Compute.ZoneOperations.Wait(c.ProjectID, zone, oper.Name).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Waiting for start instance oper error: %v\n", err)
		return err
	}
	return nil
}

func (c *Client) StopInstance(ctx context.Context, zone, instance string) error {
	c.infoLog.Println("Stopping an instance")
	oper, err := c.Compute.Instances.Stop(c.ProjectID, zone, instance).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Stop instance error: %v\n", err)
		return err
	}
	_, err = c.Compute.ZoneOperations.Wait(c.ProjectID, zone, oper.Name).Context(ctx).Do()
	if err != nil {
		c.errorLog.Printf("Waiting for stop instance oper error: %v\n", err)
		return err
	}
	return nil
}
