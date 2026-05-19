Swapzy
======
Swapzy is a small CLI that converts JSON to YAML or YAML to JSON. And also it can save the data in a PostgreSQL database. It is built using the Cobra library for creating powerful modern CLI applications.

Setup
-----
Install or tidy dependencies:

```bash
go mod tidy
```

### Note
Please make sure you have ```PostgreSQL``` installed on your system to run this CLI. And make a database named ```goku``` in your PostgreSQL server.

### Database Migration:
Before migrate, make sure you have ```.env``` file in the root directory of the project with the following ```.env.example``` file:

```bash
make migrate-up
```

Run
---
You can run the CLI using `go run` command:

```bash
- Add or Dump data
go run main.go -i save ./data/config.json
```
```bash
- List all data
go run main.go list
```
```bash
- Update data
go run main.go update -k APP_PORT -v 8080
```
```bash
- Delete data
go run main.go delete -k APP_PORT
```
```bash
- Drop the table
go run main.go drop
```

OR
===
Build the binary and run it:

```bash
go build
``` 
```bash
- Add or Dump data
./swapzy -i save ./data/config.json
```
```bash
- List all data
./swapzy list
```
```bash
- Update data
./swapzy update -k APP_PORT -v 8080
```     
```bash
- Delete data
./swapzy delete -k APP_PORT 
```
```bash
- Drop the table
./swapzy drop
```