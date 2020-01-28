# status

This is a simple Linux CLI tool to control a Luxafor Flag. It can
display the color red or the color green.

## Setup and installation

Add the following to `/etc/udev/rules.d/luxafor.rules` to enable Flag
access for all users.

```
SUBSYSTEM=="usb", ATTR{idVendor}=="04d8", ATTR{idProduct}=="f372" MODE="0666"
```

udev can be restarted like this:

```
$ sudo udevadm control --reload
$ sudo udevadm trigger
```

If you're compiling from source, you will require a working Go environment.

```bash
$ go install github.com/abrander/status
```

# Usage

`status` can either turn the flag red - or green.

Use `status red` to turn the flag red.
Use `status green` to turn the flag green.
