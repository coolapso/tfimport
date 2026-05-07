# AI Agents Guide for `tfimport`

## Application Overview
`tfimport` is a Go-based CLI tool designed to automate the process of importing existing infrastructure into a Terraform/OpenTofu state file. 

The core workflow is:
1. The user creates Terraform or Terragrunt manifests that reflect their existing infrastructure.
2. `tfimport` executes a `plan` (or reads an existing plan file).
3. The tool parses the JSON representation of the plan (`terraform show -json`).
4. It extracts resources marked for `create` and determines what string ID is required to import them into the state.

## Project Structure
- `cmd/`: Contains the CLI commands built with Cobra. `cmd/root.go` contains the main initialization, execution logic, and plan parsing logic.
- `internal/providers/`: Contains the logic for determining the specific import ID format required by different Terraform providers (e.g., AWS, Kubernetes).

## Provider Extensibility Guide (How to Add Support for New Resources/Providers)
Because Terraform providers define their own rules for `terraform import` (e.g., some need an ARN, some need a composite `namespace/name`, some need a simple string name), `tfimport` delegates the ID deduction logic to the `internal/providers` package.

### 🚫 STRICT WARNING: No Generic/Heuristic Resolvers
**Under no circumstances should you attempt to write a "generic" or "heuristic" resolver as a shortcut** (e.g., a function that generically looks for `name` or `metadata.name` across all a provider's resources). You MUST generate explicit switch statements mapping every resource individually.

**Why?**
1. **Inconsistency:** Terraform providers are notoriously inconsistent. While 90% of resources might use a simple `name`, the remaining 10% will have entirely unique requirements.
2. **Missing Edge Cases:** Generic resolvers always fail on complex resources. For example, `kubernetes_manifest` requires `apiVersion=X,kind=Y,namespace=Z,name=W`, which a generic `namespace/name` fallback would completely ruin.
3. **Silent Failures:** A heuristic fallback might return an incorrectly formatted ID, resulting in frustrating and confusing user errors during the actual `terraform import` execution. Explicitly generated mappings ensure unsupported or computed IDs return `""` immediately.

### Mass Generation (The Doc Cruncher Trick)
When bootstrapping a completely new provider or updating an existing one with hundreds of new resources, doing it manually is impossible. Instead, AI Agents MUST use the "Doc Cruncher" workflow. 

The `doc_cruncher.go` script uses a **Strategy Pattern** to accommodate the fact that different Terraform providers document their import formats entirely differently. 

1. **Clone the Provider Documentation:**
   Clone the official provider repository to a temporary directory (e.g., `/tmp`). Most HashiCorp providers store their resource markdown docs in `website/docs/r/` or `docs/resources/`.
   ```bash
   git clone --depth 1 https://github.com/hashicorp/terraform-provider-aws.git /tmp/terraform-provider-aws
   ```

2. **Use the Generator Script:**
   Navigate to the `scripts/` directory and execute the `doc_cruncher` tool using `go run .`. This script executes the specific strategy assigned to the provider to extract keys and generate Go code.
   ```bash
   cd scripts
   go run . -provider=aws -docs=/tmp/terraform-provider-aws/website/docs/r
   ```

3. **Adapt the Generator (Adding New Providers):**
   If you are tasked with adding a new provider (e.g., `google`), you **MUST NOT** hardcode logic directly into `doc_cruncher.go` in a way that overwrites or breaks the parsing logic for existing providers (like `aws` or `kubernetes`).
   
   Instead, you must create a new file `scripts/strategy_<provider>.go` (e.g., `scripts/strategy_google.go`) and add a new `ProviderStrategy` to the `strategies` map inside `doc_cruncher.go`:
   * **`MatchRegex`**: The Regex used to identify the "Import" command string within that provider's documentation.
   * **`ExtractFunc`**: The logic to extract the actual terraform keys (e.g., `name`, `id`, `namespace`) from the matched string.
   * **`GenerateFunc`**: The string template function responsible for writing the Go file structure. Note: If your provider uses standard top-level attributes, you can often just copy the `GenerateFunc` used by the AWS strategy.

**CRITICAL RULE:** The `doc_cruncher.go` must remain universally backwards compatible and idempotent. The intention is that we can run `doc_cruncher.go -provider=aws` six months from now, point it at the newest AWS provider documentation, and update all resources seamlessly. Do not introduce breaking changes to existing parsing strategies.

### Manual Addition & API Lookups (Custom Resolvers)
If a resource is not correctly mapping its Import ID after running the generator, or if the ID is purely opaque (like a VPC ID `vpc-xxxxxx`) and requires an API lookup using the provider SDK:

1. **Leverage the Context:** `tfimport` passes a `ProviderContext` down to the `internal/providers` package. This context lazy-loads cloud SDK clients (like the AWS SDK via `ctx.GetAWSClient()`) only when the specific provider requires it.
2. **Use the Custom Resolver Hook:** The `doc_cruncher.go` script automatically generates a hook for custom resolvers (e.g., `resolveCustomextractAWSImportID`). 
3. **Implement the Lookup:** Create or edit `internal/providers/<provider>_resolvers.go` and implement the custom logic. Use the context to execute API calls (like `DescribeVpcs` or `ListPolicies`) to find the resource based on attributes (like tags or prefixes) known in the plan. Return `""` to fall back to the standard generated mappings.

## Mandatory Code Quality Checks
As an AI Agent working on this project, **you are absolutely required to execute the following commands** before concluding your task to ensure the codebase remains clean and adheres to Go standards:

1. **Format the code:**
   ```bash
   go fmt ./...
   ```
2. **Run the linter:**
   ```bash
   golangci-lint run
   ```
If the linter produces any warnings or errors, you must fix them before finishing your implementation.

1. **Locate the Provider File:** Check if `internal/providers/<provider>.go` exists (e.g., `aws.go`, `gcp.go`). If not, create it.
2. **Update the Router:** If creating a new provider file, update `internal/providers/provider.go` to route the resource prefix (e.g., `google_`) to your new provider's extraction function.
3. **Read Provider Documentation:** To figure out what ID format a resource expects, search the official Terraform/OpenTofu provider documentation for the specific resource (e.g., `aws_s3_bucket`) and scroll down to the "Import" section at the very bottom.
4. **Implement the Logic:** The resource's planned configuration is passed as `config map[string]any`. You must extract the relevant fields from this map (like `name`, `bucket`, `metadata[0].name`) to build the string ID.

### Known Limitations & Assumptions
- **Auto-Generated IDs:** Some resources (like AWS VPCs or EC2 instances) use provider-generated physical IDs (e.g., `vpc-xxxxx`, `i-xxxxxxx`) that are not known during the plan phase before the resource actually exists. Currently, if an ID cannot be determined purely from the configuration or via a reliable SDK lookup, the provider function should return an empty string `""`.
- **Handling Computed Attributes (The Expression Tracer Pattern):** If a resource relies on an attribute from another resource being created in the *same* plan (e.g., an `aws_iam_role_policy_attachment` relying on a newly created `aws_iam_policy`'s ARN), that attribute will be "computed" (unknown). Terraform will omit it entirely from the JSON plan's `Change.After` block.
  To solve this without giving up and returning `""`, `tfimport` uses the **Expression Tracer Pattern** (see `resolveAttribute` in `internal/providers/aws_resolvers.go`). Since `ProviderContext` holds the full `tfjson.Plan` and `CurrentResource`, we can:
  1. Inspect the raw `Config.RootModule` to find the exact `ConfigResource` for the current address.
  2. Extract the `Expressions` for the missing attribute and trace its `References` (e.g., `aws_iam_policy.my_policy.arn`).
  3. Look up the referenced resource (`aws_iam_policy.my_policy`) in the plan's `ResourceChanges`.
  4. If the referenced attribute is a computable identifier (like `arn`, `id`, or `name`), we recursively call `GetImportID()` on the referenced resource to generate its ID ahead of time.
- **Kubernetes Specifics:** Kubernetes configuration usually nests identifiers inside a `metadata` array block.

## Example: Adding a New Resource
If you were to add support for `aws_dynamodb_table`:
1. Check the AWS provider docs for `aws_dynamodb_table` imports. It imports by `name`.
2. Edit `internal/providers/aws.go`.
3. Add a case for `"aws_dynamodb_table"`.
4. Extract the name: 
   ```go
   if name, ok := config["name"].(string); ok && name != "" {
       return name
   }
   ```
