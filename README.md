# urandomtweets

This program is designed to randomly pick some words and tweet it. These words have to be connected to a url somehow.

By now, it only supports random wikipedia articles.

## Setup

To compile the program locally, install go and glide. Go has to be version 1.5 or newer in order to support the vendor directory.

Run `glide install` to install all dependencies.

Run `make install` to install the program, `make build` just for building and `make rpi` for crosscompiling it to run on a raspberry pi.