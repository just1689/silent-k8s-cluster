# silent-k8s-cluster

## Initial setup

```shell
# Generate the config files
go run app.go -generateAll true
```

## Update each config file

- `job-config.json` specifies the overall job requirements
- `machine-specs-config.json` details each spec of a machine
- `router-config.json` is for specifying the Mikrotik router connection details

## Run

```shell
go run app.go
```