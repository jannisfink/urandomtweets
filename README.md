# urandomtweets

This program is designed to randomly pick some words and tweet it. These words have to be connected to a url somehow.

By now, it only supports random wikipedia articles.

## Setup

To compile the program locally, install go and glide. Go has to be version 1.5 or newer in order to support the vendor directory.

Run `glide install` to install all dependencies.

Run `make install` to install the program, `make build` just for building and `make rpi` for crosscompiling it to run on a raspberry pi.

## Run the program

To actually run the program, you will have to copy the file named `config.default.toml` into the same directory the binary lives in and rename it to `config.toml`.

Open [https://apps.twitter.com/](https://apps.twitter.com/) and create a new application. Generate consumer key and access token and copy all of them into the config file. Then run the program. 
