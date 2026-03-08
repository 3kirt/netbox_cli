# tenancy

Commands for the NetBox Tenancy API — tenants, tenant groups, contacts, and contact assignments.

## Actions

| Action | Flags | Description |
|---|---|---|
| `list` | | Fetch all records |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON body from stdin |
| `delete` | `--id <ID>` | Delete a record |

## Resources

| Resource | Description | Actions |
|---|---|---|
| `contact-assignments` | Assign a contact to an object | list, get, create, update, delete |
| `contact-groups` | Contact groups | list, get, create, update, delete |
| `contact-roles` | Contact roles | list, get, create, update, delete |
| `contacts` | Contacts | list, get, create, update, delete |
| `tenant-groups` | Tenant groups | list, get, create, update, delete |
| `tenants` | Tenants | list, get, create, update, delete |

## Examples

```bash
# List all tenants
netbox-cli tenancy tenants list

# Get a tenant by ID
netbox-cli tenancy tenants get --id 4

# Create a tenant group
echo '{"name": "Customers", "slug": "customers"}' \
  | netbox-cli tenancy tenant-groups create

# Create a tenant
echo '{"name": "Acme Corp", "slug": "acme-corp"}' \
  | netbox-cli tenancy tenants create

# Assign a tenant to a group
echo '{"name": "Acme Corp", "slug": "acme-corp", "group": 1}' \
  | netbox-cli tenancy tenants create

# Update a tenant's description
echo '{"description": "Primary enterprise customer"}' \
  | netbox-cli tenancy tenants update --id 4

# Delete a tenant
netbox-cli tenancy tenants delete --id 4

# Create a contact
echo '{
  "name": "Jane Smith",
  "email": "jane@example.com",
  "phone": "+1-555-0100"
}' | netbox-cli tenancy contacts create

# Create a contact role
echo '{"name": "NOC", "slug": "noc"}' \
  | netbox-cli tenancy contact-roles create

# Assign a contact to a device
echo '{
  "object_type": "dcim.device",
  "object_id": 10,
  "contact": 2,
  "role": 1
}' | netbox-cli tenancy contact-assignments create

# List all tenants and show name and slug
netbox-cli tenancy tenants list | jq '.[] | {id, name, slug}'
```
