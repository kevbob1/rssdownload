RSS download
===========

RSS attachment downloader written in Go


## Configuration

* create a config file 'rssdownload.conf'  put it in the current directory. This file contains a list of atom/rss feed urls.  One per line.


* create a /download directory use to store the cache and output files.

## Build

go build drule.org/rssdownload/rssdownload


## Usage

(after creating config file and download dir above)

./rssdownload 

status and errors are printed to stdout

## Acknowledgements

Thanks to https://github.com/jteeuwen for creating a cool, easy to use, rss fetch/parser library.

