package client

import (
	"bytes"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/jukree/iris/core/netutil"
)

const host = "http://live.iris-go.com"

// PostForm performs the PostForm with a secure client.
func PostForm(p string, data url.Values) (*http.Response, error) {
	client := netutil.Client(25 * time.Second)

	if len(data) == 0 {
		data = make(url.Values, 1)
	}
	data.Set("X-Auth", a)

	u := host + p
	r, err := client.PostForm(u, data)
	return r, err
}

var a string

func init() {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, f := range interfaces {
			if f.Flags&net.FlagUp != 0 && bytes.Compare(f.HardwareAddr, nil) != 0 {
				a = f.HardwareAddr.String()
				break
			}
		}
	}
}
