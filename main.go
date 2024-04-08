package main

import (
	"context"
	"fmt"
	"io"
	"os"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
)

// stopInstance stops a started Google Compute Engine instance
func stopInstance(w io.Writer, projectID, zone, instanceName string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	req := &computepb.StopInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	op, err := instancesClient.Stop(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to stop instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance stopped\n")

	return nil
}

// startInstance starts a stopped Google Compute Engine instance (with unencrypted disks).
func startInstance(w io.Writer, projectID, zone, instanceName string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	req := &computepb.StartInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	op, err := instancesClient.Start(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to start instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance started\n")

	return nil
}

// startInstanceWithEncKey starts a stopped Google Compute Engine instance (with encrypted disks).
func startInstanceWithEncKey(w io.Writer, projectID, zone, instanceName, key string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"
	// key := "your_encryption_key"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	instanceReq := &computepb.GetInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	instance, err := instancesClient.Get(ctx, instanceReq)
	if err != nil {
		return fmt.Errorf("unable to get instance: %w", err)
	}

	req := &computepb.StartWithEncryptionKeyInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
		InstancesStartWithEncryptionKeyRequestResource: &computepb.InstancesStartWithEncryptionKeyRequest{
			Disks: []*computepb.CustomerEncryptionKeyProtectedDisk{
				{
					Source: proto.String(instance.GetDisks()[0].GetSource()),
					DiskEncryptionKey: &computepb.CustomerEncryptionKey{
						RawKey: proto.String(key),
					},
				},
			},
		},
	}

	op, err := instancesClient.StartWithEncryptionKey(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to start instance with encryption key: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance with encryption key started\n")

	return nil
}

func resetInstance(w io.Writer, projectID, zone, instanceName string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	req := &computepb.ResetInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	op, err := instancesClient.Reset(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to reset instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance reset\n")

	return nil
}

func main() {
	if err := startInstance(os.Stdout, "inbound-respect-407316", "us-west2-a", "upwork"); err != nil {
		fmt.Println(err)
	}
}
