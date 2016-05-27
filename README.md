sup [![Build Status](https://travis-ci.org/voutasaurus/sup.svg?branch=master)](https://travis-ci.org/voutasaurus/sup)
=======

sup is ping for https.

# install

```
$ go get github.com/voutasaurus/sup
```

# details

It does the full request including getting the response so if the landing page is large then it will be as slow as that. In other words it doesn't measure the TLS negotiation, it measures the entire request/response operation.

# usage

Specify the domain name as the first and only argument to sup.

```
$ sup google.com
442.210126ms
102.355754ms
98.446453ms
100.002899ms
99.18354ms
...
```

# why?

Because it was easier than reading pages and pages of people saying "ping is over ICMP, HTTPS is different".

I know. We get it. Can we move on please?

# does it do x?

No. Feel free to submit PRs though.
