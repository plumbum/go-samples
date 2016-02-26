Install Go 1.5+
===============

  * [Download Go 1.5+ version](https://golang.org/dl/)
  * Unpack archive here
  * Setup workspace environment ``. do_go1.5.sh``
  * Install **gomobile**:

    $ go get golang.org/x/mobile/cmd/gomobile
    $ gomobile init

Build and install ``apk`` packet:

    $ gomobile build -target android

or native application

    $ go build

