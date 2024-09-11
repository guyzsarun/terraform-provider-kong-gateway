# terraform-provider-kong-gateway

This repository contains a Terraform provider for [Kong Gateway](https://docs.konghq.com/gateway/?utm_source=github&utm_campaign=terraform-provider-kong-gateway).

## Capabilities

This provider can manage the following resources:

- Service
- Route
- Plugins
- Consumers
- Consumer Groups
  - Consumer Group Members
- Credentials
  - Basic Auth
  - Key Auth
  - JWT
  - HMAC Auth
  - ACL
- Upstream
- Target
- Vault
- Certificate
- CA Certificate
- SNI
- Key
- Key Set

## Usage

The provider can be installed from the Terraform registry. Before using the provider, you must configure a `server_url`.

```hcl
terraform {
  required_providers {
    kong-gateway = {
      source = "kong/kong-gateway"
    }
  }
}

provider "kong-gateway" {
  server_url = "http://localhost:8001"
}
```

## Examples

The examples directory contains sample usage for all supported resources. For a full list of supported parameters for each resource, see the [provider documentation](https://registry.terraform.io/providers/Kong/kong-gateway/latest/docs).
