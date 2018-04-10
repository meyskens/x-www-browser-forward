x-www-browser-forward
====================

If you're running apps inside Docker containers you might have noticed URLs won't open. This is of course because it is looking for a browser in it's own container. The simplest way to set a browser is by setting `/usr/bin/x-www-browser`, which is where to put the client binary. 
On the host the server runs which listens to the socket for browser calls and forwards it to it's own /usr/bin/x-www-browser` which is a browser on the host or a script to call the container of your choice!

## How to use
1) Compile both the client and server with `go build`
2) Place the client as `/usr/bin/x-www-browser` in the containers
3) Make the server run as a service (insert SystemD here)
4) Link the `/var/run/browser.sock` socket to the container

Note: it doesn't work in every app yet due some diversity. It works in Slack! Some apps, like Thunderbird, can also be modified to point to the correct application that opens the browser.

### Optional command line arguments for the server part

```
Usage of x-www-browser-forwarder:
  -browser-cmd string
        Command to open URL (default "x-www-browser")
  -socket-file string
        Socket address (default "/var/run/browser.sock")
```

Note: The client defaults are not adjustable since you have full control with Docker on how to map the browser command and socket into the container.

