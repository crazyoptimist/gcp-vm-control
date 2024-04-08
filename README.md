# Google Cloud VM Control

This application start/stop Google Cloud VMs as needed.

## Generate Application Default Credentials

To use this application, you need to install the gcloud CLI and initialize it.

https://cloud.google.com/sdk/docs/install

I 100% assume that your google account has necessary permissions to start and stop the VM

- To stop a VM: `compute.instances.stop` on the VM
- To start a VM: `compute.instances.start` on the VM
- To start a VM that uses encryption keys: `compute.instances.startWithEncryptionKey` on the VM
- To reset a VM: `compute.instances.reset` on the VM

```bash
# Initialize the cli
gcloud init

# Generate ADC (application default credentials)
gcloud auth application-default login
```

## Usage

```bash
gcp-vm-control --help
```
