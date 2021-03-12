package torclient

import (
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/net/proxy"
)

func New(port int) (*http.Client, error) {
	// Create a transport that uses Tor Browser's SocksPort.
	tbProxyURL, err := url.Parse("socks5://127.0.0.1:" + strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	// Get a proxy Dialer that will create the connection on our
	// behalf via the SOCKS5 proxy.
	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		return nil, err
	}

	// Make a http.Transport that uses the proxy dialer, and a
	// http.Client that uses the transport.
	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}

	return client, nil
}
