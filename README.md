
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

[![Build Status](https://travis-ci.org/jkassismz/digits.svg?branch=master)](https://travis-ci.org/jkassismz/digits)

digits is a tiny program that sends some load to a web application.

digits was originally called boom and was influenced from Tarek Ziade's
tool at [tarekziade/boom](https://github.com/tarekziade/boom). Using the same name was a mistake as it resulted in cases
where binary name conflicts created confusion.
To preserve the name for its original owner, we renamed this project to digits.

## Installation

* Linux 64-bit: https://storage.googleapis.com/digits-release/digits_linux_amd64
* Mac 64-bit: https://storage.googleapis.com/digits-release/digits_darwin_amd64
* Windows 64-bit: https://storage.googleapis.com/digits-release/digits_windows_amd64

### Package Managers

macOS:
-  [Homebrew](https://brew.sh/) users can use `brew install digits`.

## Usage

digits runs provided number of requests in the provided concurrency level and prints stats.

It also supports HTTP2 endpoints.

```
Usage: digits [options...] <url>

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

  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  -A  HTTP Accept header.
  -d  HTTP request body.
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  -T  Content-type, defaults to "text/html".
  -a  Basic authentication, username:password.
  -x  HTTP Proxy address as host:port.
  -h2 Enable HTTP/2.

  -host	HTTP Host header.

  -disable-compression  Disable compression.
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  -disable-redirects    Disable following of HTTP redirects
  -cpus                 Number of used cpu cores.
                        (default for current machine is 8 cores)
```

Previously known as [github.com/rakyll/boom](https://github.com/rakyll/boom).
