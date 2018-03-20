sup [![Build Status](https://travis-ci.org/voutasaurus/sup.svg?branch=master)](https://travis-ci.org/voutasaurus/sup)
=======

sup is ping for https.

# install

Install the latest version of Go: https://golang.org/dl/

Then install sup from a terminal:
```
$ go get github.com/voutasaurus/sup
```

# details

It does the full request including getting the response so if the landing page is large then it will be as slow as that. In other words it doesn't measure only the TLS negotiation, it measures the entire request/response operation.

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

If you want to be verbose sup will do one request:
```
$ sup -v google.com
2018/03/20 11:19:47.557200 state: GetConn, hostPort: "google.com:443"
2018/03/20 11:19:47.557823 state: DNSStart, info: {Host:google.com}
2018/03/20 11:19:47.570227 state: DNSDone, info: {Addrs:[{IP:172.217.6.174 Zone:} {IP:2607:f8b0:4000:804::200e Zone:}] Err:<nil> Coalesced:false}
2018/03/20 11:19:47.570326 state: ConnectStart, network: "tcp", addr: "172.217.6.174:443"
2018/03/20 11:19:47.630020 state: ConnectDone, network: "tcp", addr: "172.217.6.174:443", err: <nil>
2018/03/20 11:19:47.630898 state: TLSHandshakeStart
2018/03/20 11:19:47.876416 state: TLSHandshakeDone, tlsConnectionState: {Version:771 HandshakeComplete:true DidResume:false CipherSuite:49195 NegotiatedProtocol:h2 NegotiatedProtocolIsMutual:true ServerName: PeerCertificates:[0xc42019a000 0xc42019a580 0xc42019ab00] VerifiedChains:[[0xc42019a000 0xc42019a580 0xc420308100]] SignedCertificateTimestamps:[[0 238 75 189 183 117 206 96 186 225 66 105 31 171 225 158 102 163 15 126 95 176 114 216 131 0 196 123 137 122 168 253 203 0 0 1 97 222 200 45 93 0 0 4 3 0 71 48 69 2 33 0 153 68 76 186 150 48 115 197 224 123 115 188 75 40 109 189 245 208 200 52 92 41 197 111 126 97 124 147 122 16 127 240 2 32 63 147 189 30 134 151 10 95 55 176 227 171 12 61 91 250 233 62 37 24 63 209 149 229 155 210 228 62 131 0 175 191] [0 111 83 118 172 49 240 49 25 216 153 0 164 81 21 255 119 21 28 17 217 2 193 0 41 6 141 178 8 154 55 217 19 0 0 1 97 222 200 45 89 0 0 4 3 0 71 48 69 2 32 18 252 239 126 227 248 66 224 210 236 58 156 117 133 134 242 78 124 48 59 12 101 112 255 244 42 199 159 112 97 234 234 2 33 0 185 32 73 254 175 174 228 62 41 215 84 241 164 124 243 54 5 235 0 74 115 224 193 173 61 123 217 14 147 170 109 239]] OCSPResponse:[] TLSUnique:[35 5 177 124 243 70 77 171 234 66 50 99]}, error: <nil>
2018/03/20 11:19:47.876808 state: GotConn, GotConnInfo: {Conn:0xc42016e000 Reused:false WasIdle:false IdleTime:0s}
2018/03/20 11:19:47.876955 state: WroteHeaders
2018/03/20 11:19:47.876981 state: WroteRequest, WroteRequestInfo: {Err:<nil>}
2018/03/20 11:19:47.930778 state: GotFirstResponseByte
2018/03/20 11:19:47.931069 state: GetConn, hostPort: "www.google.com:443"
2018/03/20 11:19:47.931129 state: DNSStart, info: {Host:www.google.com}
2018/03/20 11:19:47.952980 state: DNSDone, info: {Addrs:[{IP:172.217.6.4 Zone:} {IP:2607:f8b0:4000:804::2004 Zone:}] Err:<nil> Coalesced:false}
2018/03/20 11:19:47.953079 state: ConnectStart, network: "tcp", addr: "172.217.6.4:443"
2018/03/20 11:19:48.005764 state: ConnectDone, network: "tcp", addr: "172.217.6.4:443", err: <nil>
2018/03/20 11:19:48.005838 state: TLSHandshakeStart
2018/03/20 11:19:48.127789 state: TLSHandshakeDone, tlsConnectionState: {Version:771 HandshakeComplete:true DidResume:false CipherSuite:49195 NegotiatedProtocol:h2 NegotiatedProtocolIsMutual:true ServerName: PeerCertificates:[0xc420194b00 0xc420195080 0xc420195600] VerifiedChains:[[0xc420194b00 0xc420195080 0xc420308100]] SignedCertificateTimestamps:[[0 164 185 9 144 180 24 88 20 135 187 19 162 204 103 112 10 60 53 152 4 249 27 223 184 227 119 205 14 200 13 220 16 0 0 1 97 222 199 12 176 0 0 4 3 0 71 48 69 2 32 22 215 174 177 248 8 14 181 235 3 22 64 146 71 16 203 41 96 201 93 130 214 188 2 161 199 183 200 192 77 21 14 2 33 0 255 15 46 200 118 81 46 116 94 164 22 141 176 151 91 52 239 73 140 170 43 184 160 155 216 117 12 9 138 34 23 102] [0 111 83 118 172 49 240 49 25 216 153 0 164 81 21 255 119 21 28 17 217 2 193 0 41 6 141 178 8 154 55 217 19 0 0 1 97 222 199 10 203 0 0 4 3 0 71 48 69 2 33 0 171 193 64 141 144 246 188 190 104 128 29 65 17 144 206 68 246 143 52 51 53 239 55 222 58 191 147 253 164 51 184 155 2 32 76 106 200 154 243 185 180 131 151 60 74 39 33 172 0 148 79 130 131 54 215 50 176 224 53 19 74 75 165 79 115 188]] OCSPResponse:[] TLSUnique:[70 44 7 207 211 100 83 105 29 75 72 237]}, error: <nil>
2018/03/20 11:19:48.127985 state: GotConn, GotConnInfo: {Conn:0xc42016ea80 Reused:false WasIdle:false IdleTime:0s}
2018/03/20 11:19:48.128070 state: WroteHeaders
2018/03/20 11:19:48.128082 state: WroteRequest, WroteRequestInfo: {Err:<nil>}
2018/03/20 11:19:48.242261 state: GotFirstResponseByte
685.564456ms
```

# why?

Because it was easier than reading pages and pages of people saying "ping is over ICMP, HTTPS is different".

I know. We get it. Can we move on please?

# does it do x?

No. Feel free to submit PRs though.
