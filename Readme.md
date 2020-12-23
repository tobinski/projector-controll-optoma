# Cli for optoma projector
Very simple cli to control (switch on and off) optoma eh460st projector. 

## Usage
```bash
# list possible RS232 ports
beamer list
# start beamer  switch on
beamer start [port]
# stop beamer switch power off
beamer stop [port]
```
## Raspberry 4
To build the cli for a raspberry 4 you can cross-compile
```bash
env GOOS=linux GOARCH=arm GOARM=5 go build
```