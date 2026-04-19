# Blueprint: IaC (Terraform/Tofu)

**Target**: Cloud Infrastructure
**Stack**: Terraform 1.9+ / OpenTofu, AWS/GCP, S3 Backend

## Directory Structure
```text
/
├── modules/            # Reusable components
├── environments/
│   ├── dev/            # Dev specific config
│   └── prod/           # Prod specific config
├── main.tf             # Entry point
├── variables.tf        # Input definitions
├── outputs.tf          # Exported values
├── providers.tf        # Cloud providers config
└── backend.tf          # Remote state config (S3/GCS)
```

## Mandatory Checks
- [ ] No secrets in `.tf` files (use `variables.tf` and environment variables).
- [ ] Remote state locking enabled.
- [ ] Tagging policy implemented (Owner, Environment, Project).
- [ ] Resource naming convention followed.

---
**Template Type**: Project Blueprint
