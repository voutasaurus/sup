package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	flagv := flag.Bool("v", false, "be really verbose, only do one request")
	flag.Parse()

	req, err := http.NewRequest("GET", "https://"+flag.Arg(0), nil)
	if err != nil {
		log.Fatalf("could not build request: %v", err)
	}
	if *flagv {
		trace := verboseTrace()
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	}

	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			start := time.Now()
			_, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(time.Since(start))
		}
		if *flagv { // only do once for verbose
			return
		}
	}
}

func verboseTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		// GetConn is called before a connection is created or
		// retrieved from an idle pool. The hostPort is the
		// "host:port" of the target or proxy. GetConn is called even
		// if there's already an idle cached connection available.
		GetConn: func(hostPort string) {
			log.Printf("state: GetConn, hostPort: %q", hostPort)
		},

		// GotConn is called after a successful connection is
		// obtained. There is no hook for failure to obtain a
		// connection; instead, use the error from
		// Transport.RoundTrip.
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("state: GotConn, GotConnInfo: %+v", info)
		},

		// PutIdleConn is called when the connection is returned to
		// the idle pool. If err is nil, the connection was
		// successfully returned to the idle pool. If err is non-nil,
		// it describes why not. PutIdleConn is not called if
		// connection reuse is disabled via Transport.DisableKeepAlives.
		// PutIdleConn is called before the caller's Response.Body.Close
		// call returns.
		// For HTTP/2, this hook is not currently used.
		PutIdleConn: func(err error) {
			log.Printf("state: PutIdleConn, error: %v", err)
		},

		// GotFirstResponseByte is called when the first byte of the response
		// headers is available.
		GotFirstResponseByte: func() {
			log.Println("state: GotFirstResponseByte")
		},

		// Got100Continue is called if the server replies with a "100
		// Continue" response.
		Got100Continue: func() {
			log.Println("state: Got100Continue")
		},

		// DNSStart is called when a DNS lookup begins.
		DNSStart: func(info httptrace.DNSStartInfo) {
			log.Printf("state: DNSStart, info: %+v", info)
		},

		// DNSDone is called when a DNS lookup ends.
		DNSDone: func(info httptrace.DNSDoneInfo) {
			log.Printf("state: DNSDone, info: %+v", info)
		},

		// ConnectStart is called when a new connection's Dial begins.
		// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
		// enabled, this may be called multiple times.
		ConnectStart: func(network, addr string) {
			log.Printf("state: ConnectStart, network: %q, addr: %q", network, addr)
		},

		// ConnectDone is called when a new connection's Dial
		// completes. The provided err indicates whether the
		// connection completedly successfully.
		// If net.Dialer.DualStack ("Happy Eyeballs") support is
		// enabled, this may be called multiple times.
		ConnectDone: func(network, addr string, err error) {
			log.Printf("state: ConnectDone, network: %q, addr: %q, err: %v", network, addr, err)
		},

		// TLSHandshakeStart is called when the TLS handshake is started. When
		// connecting to a HTTPS site via a HTTP proxy, the handshake happens after
		// the CONNECT request is processed by the proxy.
		TLSHandshakeStart: func() {
			log.Println("state: TLSHandshakeStart")
		},

		// TLSHandshakeDone is called after the TLS handshake with either the
		// successful handshake's connection state, or a non-nil error on handshake
		// failure.
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			log.Printf("state: TLSHandshakeDone, tlsConnectionState: %+v, error: %v", state, err)
		},

		// WroteHeaders is called after the Transport has written
		// the request headers.
		WroteHeaders: func() {
			log.Println("state: WroteHeaders")
		},

		// Wait100Continue is called if the Request specified
		// "Expected: 100-continue" and the Transport has written the
		// request headers but is waiting for "100 Continue" from the
		// server before writing the request body.
		Wait100Continue: func() {
			log.Println("state: Wait100Continue")
		},

		// WroteRequest is called with the result of writing the
		// request and any body. It may be called multiple times
		// in the case of retried requests.
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			log.Printf("state: WroteRequest, WroteRequestInfo: %+v", info)
		},
	}
}
