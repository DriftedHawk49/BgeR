# BgeR
### Background shuffler for Linux
A command-line tool for linux built using `Go (Golang)` which is able to change background as slideshow. Currently tested on Ubuntu 22.04 and GNOME Desktop Environment.


## How To Use
Download the binary from release section, or compile it yourself. Then add the executable's path to your `$PATH`

  It accepts 2 flags, `dir` for specifying the directory from where it will pick images from (make sure your current user has sufficient permission or run the tool as sudo, and provide absolute path) and `dur` for duration in seconds after which it should change the background. `dur` is optional,  defaults to 60 seconds.


  #### Example Usage : `bger -dir /home/user1/Wallpapers -dur 60`

  You can add new images to wallpaper folder while the tool is running and it will pick up new images automatically. No need to re-deploy the tool after adding or removing any images. Supports only `.jpeg OR .jpg` files.

  To run it in detached mode, use systemd service to manage auto startup, start, stop and other background functionalities.

  To use systemd service, you need to get the value of XDG_CURRENT_DESKTOP env variable and add its value in service file. replace `gnome` with the value that you get by running `echo $XDG_CURRENT_DESKTOP` in terminal.

  A simple systemd service file will look like this : 

```
[Unit]
Description=BackgroundR
After=network.target

[Service]
Environment="XDG_CURRENT_DESKTOP=gnome"
ExecStart=<path to bger exectuable> -dir <path to wallpapers folder> -dur 30
Type=simple
Restart=always

[Install]
WantedBy=default.target
```

## Contributing
Support for several other Desktop environments is still to be added, and tested. You are welcome to suggest better approach(es) and enhance the experience.  


