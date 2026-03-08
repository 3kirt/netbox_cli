# dcim

Commands for the NetBox DCIM API — devices, racks, cables, interfaces, and all physical infrastructure.

## Actions

| Action | Flags | Description |
|---|---|---|
| `list` | | Fetch all records |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON body from stdin |
| `delete` | `--id <ID>` | Delete a record |

## Resources

All DCIM resources support list, get, create, update, and delete.

| Resource | Description |
|---|---|
| `cable-terminations` | Individual termination points of a cable |
| `cables` | Cables between termination points |
| `console-port-templates` | Console port templates on device types |
| `console-ports` | Console ports on devices |
| `console-server-port-templates` | Console server port templates on device types |
| `console-server-ports` | Console server ports on devices |
| `device-bay-templates` | Device bay templates on device types |
| `device-bays` | Device bays on devices |
| `device-roles` | Device roles |
| `device-types` | Device type definitions |
| `devices` | Physical devices |
| `front-port-templates` | Front port templates on device types |
| `front-ports` | Front ports on devices |
| `interface-templates` | Interface templates on device types |
| `interfaces` | Device interfaces |
| `inventory-item-roles` | Inventory item roles |
| `inventory-item-templates` | Inventory item templates on device types |
| `inventory-items` | Inventory items installed in devices |
| `locations` | Locations within a site |
| `mac-addresses` | MAC address records |
| `manufacturers` | Hardware manufacturers |
| `module-bay-templates` | Module bay templates on device types |
| `module-bays` | Module bays on devices |
| `module-type-profiles` | Module type profiles |
| `module-types` | Module type definitions |
| `modules` | Modules installed in devices |
| `platforms` | Software platforms |
| `power-feeds` | Power feeds from panels to devices |
| `power-outlet-templates` | Power outlet templates on device types |
| `power-outlets` | Power outlets on devices |
| `power-panels` | Power panels in racks |
| `power-port-templates` | Power port templates on device types |
| `power-ports` | Power ports on devices |
| `rack-reservations` | Rack unit reservations |
| `rack-roles` | Rack roles |
| `rack-types` | Rack type definitions |
| `racks` | Racks |
| `rear-port-templates` | Rear port templates on device types |
| `rear-ports` | Rear ports on devices |
| `regions` | Geographic regions |
| `site-groups` | Site groups |
| `sites` | Sites |
| `virtual-chassis` | Virtual chassis |
| `virtual-device-contexts` | Virtual device contexts |

## Examples

### Sites and regions

```bash
# List all sites
netbox-cli dcim sites list

# Create a region
echo '{"name": "North America", "slug": "north-america"}' \
  | netbox-cli dcim regions create

# Create a site
echo '{
  "name": "New York DC",
  "slug": "new-york-dc",
  "status": "active",
  "region": 1
}' | netbox-cli dcim sites create

# Update a site's time zone
echo '{"time_zone": "America/New_York"}' \
  | netbox-cli dcim sites update --id 3

# Delete a site
netbox-cli dcim sites delete --id 3
```

### Racks

```bash
# List all racks
netbox-cli dcim racks list

# Create a rack role
echo '{"name": "Network", "slug": "network", "color": "0000ff"}' \
  | netbox-cli dcim rack-roles create

# Create a rack
echo '{
  "name": "A01",
  "site": 1,
  "status": "active",
  "u_height": 42
}' | netbox-cli dcim racks create

# Update a rack's location
echo '{"location": 2}' \
  | netbox-cli dcim racks update --id 5
```

### Devices

```bash
# List all devices
netbox-cli dcim devices list

# Get a device by ID
netbox-cli dcim devices get --id 100

# Create a manufacturer
echo '{"name": "Cisco", "slug": "cisco"}' \
  | netbox-cli dcim manufacturers create

# Create a platform
echo '{"name": "IOS-XE", "slug": "ios-xe"}' \
  | netbox-cli dcim platforms create

# Create a device type
echo '{
  "manufacturer": 1,
  "model": "Catalyst 9300",
  "slug": "catalyst-9300",
  "u_height": 1
}' | netbox-cli dcim device-types create

# Create a device role
echo '{"name": "Access Switch", "slug": "access-switch", "color": "00ff00"}' \
  | netbox-cli dcim device-roles create

# Create a device
echo '{
  "name": "sw-access-01",
  "device_type": 1,
  "role": 1,
  "site": 1,
  "status": "active"
}' | netbox-cli dcim devices create

# Update a device's serial number
echo '{"serial": "FDO2142A0BC"}' \
  | netbox-cli dcim devices update --id 100

# Delete a device
netbox-cli dcim devices delete --id 100
```

### Interfaces and cables

```bash
# List all interfaces on a device (filter with jq)
netbox-cli dcim interfaces list | jq '.[] | select(.device.id == 100)'

# Create an interface
echo '{
  "device": 100,
  "name": "GigabitEthernet0/1",
  "type": "1000base-t",
  "enabled": true
}' | netbox-cli dcim interfaces create

# Create a cable between two interfaces
echo '{
  "a_terminations": [{"object_type": "dcim.interface", "object_id": 201}],
  "b_terminations": [{"object_type": "dcim.interface", "object_id": 202}],
  "status": "connected"
}' | netbox-cli dcim cables create
```

### jq tips

```bash
# List devices with their site and role
netbox-cli dcim devices list | jq '.[] | {id, name, site: .site.name, role: .role.name}'

# Find all active devices at a site
netbox-cli dcim devices list | jq '.[] | select(.site.id == 1 and .status.value == "active")'

# List racks and their fill percentage
netbox-cli dcim racks list | jq '.[] | {id, name, u_height}'
```
