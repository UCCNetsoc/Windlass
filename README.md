# Windlass

## UCC Netsoc Containers as a Service

## Developing Locally

Please checkout our [Netsoc Developer Environment](https://github.com/UCCNetworkingSociety/dev-env)

Add this to your Windlass override for debugging and live reload in the Docker container
```yaml
version: '3.7'
services:
windlass:
    command: task debug -w
    security_opt:
    - seccomp:unconfined
    ports:
    - 3456:3456
    volumes:
    - /path/to/your/Windlass:/windlass
```