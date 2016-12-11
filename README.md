brightnessctl
=============

A simple brightness controller. It helps change device values like keyboards and displays. You point to the directory
with a `max_brightness` and `brightness` interface, and specify an `amount`. Everything else is :ok_hand:.

```
$ ./brightnessctl --help
Usage of ./brightnessctl:
  -amount int
        Increate or decreate by this percentage. Ex: 5 or -5. (default 5)
  -device string
        The options are 'keyboard' or 'display' or '/sys/class/etcetcetc'.
```

### Building

```
$ go build brightness.go
```
