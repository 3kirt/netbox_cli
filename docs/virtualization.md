# virtualization

Commands for the NetBox Virtualization API — clusters, virtual machines, interfaces, and virtual disks.

## Actions

| Action | Flags | Description |
|---|---|---|
| `list` | | Fetch all records |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON body from stdin |
| `delete` | `--id <ID>` | Delete a record |

## Resource-specific flags

### `virtual-machines`

`list` accepts optional filter flags. Omitting all flags returns all records.

| Flag | Description |
|---|---|
| `--name <name>` | Filter by exact name |
| `--site <slug>` | Filter by site slug |
| `--role <slug>` | Filter by role slug |
| `--cluster <name>` | Filter by cluster name |
| `--status <status>` | Filter by status (`active`, `staged`, `offline`, `planned`, `decommissioning`) |
| `--tag <tag>` | Filter by tag (repeatable or comma-separated) |

`get` accepts `--name <name>` as an alternative to `--id`. Returns an error if multiple VMs share the name.

## Resources

| Resource | Description | Actions |
|---|---|---|
| `cluster-groups` | Cluster groups | list, get, create, update, delete |
| `cluster-types` | Cluster types | list, get, create, update, delete |
| `clusters` | Clusters | list, get, create, update, delete |
| `interfaces` | Virtual machine interfaces | list, get, create, update, delete |
| `virtual-disks` | Virtual disks | list, get, create, update, delete |
| `virtual-machines` | Virtual machines | list, get, create, update, delete |

## Examples

```bash
# List all virtual machines
netbox-cli virtualization virtual-machines list

# Get a virtual machine by ID
netbox-cli virtualization virtual-machines get --id 12

# Get a virtual machine by name
netbox-cli virtualization virtual-machines get --name web-01

# Create a cluster type
echo '{"name": "VMware vSphere", "slug": "vmware-vsphere"}' \
  | netbox-cli virtualization cluster-types create

# Create a cluster
echo '{"name": "prod-vsphere-01", "type": 1, "site": 2}' \
  | netbox-cli virtualization clusters create

# Create a virtual machine
echo '{
  "name": "web-01",
  "cluster": 1,
  "status": "active",
  "vcpus": 4,
  "memory": 8192,
  "disk": 100
}' | netbox-cli virtualization virtual-machines create

# Update a VM's status
echo '{"status": "staged"}' \
  | netbox-cli virtualization virtual-machines update --id 12

# Delete a virtual machine
netbox-cli virtualization virtual-machines delete --id 12

# Create a VM interface
echo '{"virtual_machine": 12, "name": "eth0", "enabled": true}' \
  | netbox-cli virtualization interfaces create

# Add a virtual disk
echo '{"virtual_machine": 12, "name": "sda", "size": 50}' \
  | netbox-cli virtualization virtual-disks create

# List VMs in a cluster
netbox-cli virtualization virtual-machines list --cluster prod-vsphere-01

# List active VMs at a site
netbox-cli virtualization virtual-machines list --site lon01 --status active

# List VMs by tag
netbox-cli virtualization virtual-machines list --tag k8s-node
```
