# Windlass

## UCC Netsoc Containers as a Service

## Developing Locally

To run Windlass, please checkout our [Netsoc Developer Environment](https://github.com/UCCNetworkingSociety/dev-env)

To build the LXD image:

1. Install [Packer](http://packer.io/)
2. Run `packer build packer.json`
3. $$ Profit $$

### Debugging

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

### Live Reload

Add this to your Windlass override for live reload only in the Docker container

```yaml
version: '3.7'
services:
windlass:
    command: task live-reload -w
    security_opt:
    - seccomp:unconfined
    ports:
    - 3456:3456
    volumes:
    - /path/to/your/Windlass:/windlass
```
