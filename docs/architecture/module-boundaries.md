# Module boundaries

## Current modules

| Module        | Owns                                  | Must not own               |
| ------------- | ------------------------------------- | -------------------------- |
| config        | environment decoding and validation   | business defaults          |
| database      | pool lifecycle, migrations, health    | tenant policy              |
| server        | HTTP and Connect transport            | domain decisions           |
| resource      | desired and observed state vocabulary | provider SDK calls         |
| plan          | proposed typed steps                  | direct execution           |
| policy        | risk and authorization decisions      | transport identity parsing |
| operation     | durable execution state semantics     | provider-specific commands |
| provider      | infrastructure capability contract    | user authentication        |
| workflow      | orchestration contract                | business validation        |
| observability | structured logs and later telemetry   | secret payloads            |

Generated Protobuf types are transport types. Domain modules should not depend
on generated Connect handlers.

## Reserved extension points

### Docker and Kubernetes

Adapters implement the Provider contract and their provider-specific typed
capabilities. Kubernetes does not replace the Aether resource vocabulary.

### Crossplane and OpenTofu

These are future external-infrastructure adapters. They must produce and
observe typed results through the same operation boundary.

### NATS

The initial event dispatcher remains in-process. NATS is introduced only when
multiple processes need durable event delivery.

### Temporal

The initial WorkflowEngine remains local. Temporal is introduced when
workflows are long-running, resumable, and compensating rather than ordinary
function calls.

### OpenBao

The control-plane database stores only secret metadata and references. OpenBao
will own actual secret values and dynamic credentials.

### AI providers

Model implementations sit behind a provider-neutral interface. Model output is
untrusted input: it must deserialize into a known Plan, pass validation and
policy, receive required approval, and execute through Operations.
