## What is cligen
This is a code generator for fastspeed Go kubectl plugin development based on client-go. 

## Usage
Create a kubectl plugin `test-cli` using cligen.
```bash
go install github.com/slaise/cligen@latest

./cligen kubectlgen -n test-cli 
```

## Files generated
cligen generates a list of files, include:
* go mod file
* cobra command: help, version, root, yourcmd
* client-go setup: config, client
