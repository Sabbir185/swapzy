Swapzy
======
Swapzy is a small CLI that converts JSON to YAML or YAML to JSON.

Setup
-----
Install or tidy dependencies:

```bash
go mod tidy
```

Run
---

JSON to YAML:

```bash
go run main.go -i ./data/config.json -o yaml
```

YAML to JSON:

```bash
go run main.go -i ./data/config.yaml -o json
```

YAML to YAML (expected error because no conversion is needed, similar for json):

```bash
go run main.go -i ./data/config.yaml -o yaml
```

OR
===
Build the binary and run it:

```bash
go build
``` 
```bash
./swapzy -i ./data/config.json -o yaml
```
```bash
./swapzy -i ./data/config.yaml -o json
```
```bash
./swapzy -i ./data/config.yaml -o yaml
```
