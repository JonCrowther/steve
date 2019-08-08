package urlbuilder

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func ParseRequestURL(r *http.Request) string {
	scheme := getScheme(r)
	host := getHost(r, scheme)
	return fmt.Sprintf("%s://%s%s%s", scheme, host, r.Header.Get(PrefixHeader), r.URL.Path)
}

func getHost(r *http.Request, scheme string) string {
	host := strings.Split(r.Header.Get(ForwardedHostHeader), ",")[0]
	if host == "" {
		host = r.Host
	}

	port := r.Header.Get(ForwardedPortHeader)
	if port == "" {
		return host
	}

	if port == "80" && scheme == "http" {
		return host
	}

	if port == "443" && scheme == "http" {
		return host
	}

	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return host
	}

	return strings.Join([]string{hostname, port}, ":")
}

func getScheme(r *http.Request) string {
	scheme := r.Header.Get(ForwardedProtoHeader)
	if scheme != "" {
		return scheme
	} else if r.TLS != nil {
		return "https"
	}
	return "http"
}

func ParseResponseURLBase(currentURL string, r *http.Request) (string, error) {
	path := r.URL.Path

	index := strings.LastIndex(currentURL, path)
	if index == -1 {
		// Fallback, if we can't find path in currentURL, then we just assume the base is the root of the web request
		u, err := url.Parse(currentURL)
		if err != nil {
			return "", err
		}

		buffer := bytes.Buffer{}
		buffer.WriteString(u.Scheme)
		buffer.WriteString("://")
		buffer.WriteString(u.Host)
		return buffer.String(), nil
	}

	return currentURL[0:index], nil
}