# Lights

Web app for running on Raspberry Pi with WS2811 light strip connected on GPIO pin 18.

Requires [rpi_ws281x](https://github.com/samueljmurray/rpi_ws281x) package installed in your Go PATH.

Built with Go's [Revel](https://revel.github.io/) web app framework.

Needs to be run with root access to control lights:

```
$ sudo -b GOPATH=/home/pi/Workspace /home/pi/Workspace/bin/revel run github.com/samueljmurray/lights
```

This will run in the background, so you can Ctrl-C after running the command.
