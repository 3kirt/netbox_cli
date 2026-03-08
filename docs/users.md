# users

Commands for the NetBox Users API — users, groups, permissions, and API tokens.

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
| `groups` | User groups | list, get, create, update, delete |
| `permissions` | Object permissions | list, get, create, update, delete |
| `tokens` | API tokens | list, get, create, update, delete |
| `users` | User accounts | list, get, create, update, delete |

## Examples

```bash
# List all users
netbox-cli users users list

# Get a user by ID
netbox-cli users users get --id 1

# Create a user
echo '{
  "username": "jsmith",
  "password": "changeme123"
}' | netbox-cli users users create

# Update a user's email
echo '{"email": "jsmith@example.com"}' \
  | netbox-cli users users update --id 3

# Delete a user
netbox-cli users users delete --id 3

# Create a group
echo '{"name": "Network Operators"}' \
  | netbox-cli users groups create

# Create an API token for a user
echo '{"user": 2, "description": "Automation token"}' \
  | netbox-cli users tokens create

# List all tokens and show user and key prefix
netbox-cli users tokens list | jq '.[] | {id, user: .user.display, key: .key[:8]}'

# Create an object permission
echo '{
  "name": "Read IPAM",
  "object_types": ["ipam.prefix", "ipam.ipaddress"],
  "actions": ["view"],
  "groups": [1]
}' | netbox-cli users permissions create
```
