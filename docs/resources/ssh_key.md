# vdsina_ssh_key (Resource)

Manages a VDSina SSH key for server authentication.

## Example Usage

### Basic SSH Key

```hcl
resource "vdsina_ssh_key" "deploy" {
  name = "deploy-key"
  data = file("~/.ssh/id_rsa.pub")
}
```

### SSH Key with Server

```hcl
resource "vdsina_ssh_key" "main" {
  name = "main-key"
  data = var.ssh_public_key
}

resource "vdsina_server" "web" {
  datacenter  = 4
  server_plan = 1
  template    = 23
  ssh_key     = vdsina_ssh_key.main.id
  name        = "web-server"
}
```

## Schema

### Required

* `name` - (String) Name of the SSH key.
* `data` - (String, Sensitive) Public SSH key content (e.g., contents of `~/.ssh/id_rsa.pub` or `~/.ssh/id_ed25519.pub`).

### Read-Only

* `id` - (String) The ID of the SSH key.

## Import

SSH keys can be imported using the key ID:

```bash
terraform import vdsina_ssh_key.example 15579
```
