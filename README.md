```
DDDDDDDDDDDDD          iiii                        iiii          tttt
D::::::::::::DDD      i::::i                      i::::i      ttt:::t
D:::::::::::::::DD     iiii                        iiii       t:::::t
DDD:::::DDDDD:::::D                                           t:::::t
  D:::::D    D:::::D iiiiiii    ggggggggg   gggggiiiiiiittttttt:::::ttttttt        ssssssssss
  D:::::D     D:::::Di:::::i   g:::::::::ggg::::gi:::::it:::::::::::::::::t      ss::::::::::s
  D:::::D     D:::::D i::::i  g:::::::::::::::::g i::::it:::::::::::::::::t    ss:::::::::::::s
  D:::::D     D:::::D i::::i g::::::ggggg::::::gg i::::itttttt:::::::tttttt    s::::::ssss:::::s
  D:::::D     D:::::D i::::i g:::::g     g:::::g  i::::i      t:::::t           s:::::s  ssssss
  D:::::D     D:::::D i::::i g:::::g     g:::::g  i::::i      t:::::t             s::::::s
  D:::::D     D:::::D i::::i g:::::g     g:::::g  i::::i      t:::::t                s::::::s
  D:::::D    D:::::D  i::::i g::::::g    g:::::g  i::::i      t:::::t    ttttttssssss   s:::::s
DDD:::::DDDDD:::::D  i::::::ig:::::::ggggg:::::g i::::::i     t::::::tttt:::::ts:::::ssss::::::s
D:::::::::::::::DD   i::::::i g::::::::::::::::g i::::::i     tt::::::::::::::ts::::::::::::::s
D::::::::::::DDD     i::::::i  gg::::::::::::::g i::::::i       tt:::::::::::tt s:::::::::::ss
DDDDDDDDDDDDD        iiiiiiii    gggggggg::::::g iiiiiiii         ttttttttttt    sssssssssss
                                         g:::::g
                             gggggg      g:::::g
                             g:::::gg   gg:::::g
                              g::::::ggg:::::::g
                               gg:::::::::::::g
                                 ggg::::::ggg
                                    gggggg
```
[![Build Status](https://travis-ci.org/jkassismz/digits.svg?branch=master)](https://travis-ci.org/jkassismz/digits)

digits is a tiny program that load tests dns servers by querying for a single hostname over and over.

digits was forked and adapted from raykll/hey

## Installation

These links might work someday... for now... pull the source and do

> go run digits.go [options...]


* Linux 64-bit: https://storage.googleapis.com/digits-release/digits_linux_amd64
* Mac 64-bit: https://storage.googleapis.com/digits-release/digits_darwin_amd64
* Windows 64-bit: https://storage.googleapis.com/digits-release/digits_windows_amd64

### Package Managers
This also doesn't work right now...

macOS:
-  [Homebrew](https://brew.sh/) users can use `brew install digits`.

## Usage

digits runs provided number of requests in the provided concurrency level and prints stats.


```
Usage: digits [options...] <host>

Options:
  -n  Number of requests to run. Default is 200.
  -c  Number of workers to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  -q  Rate limit, in queries per second (QPS) per worker. Default is no rate limit.
  -z  Duration of application to send requests. When duration is reached,
      application stops and exits. If duration is specified, n is ignored.
      Examples: -z 10s -z 3m.
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.

  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.

  -cpus                 Number of used cpu cores.
```
