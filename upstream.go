package hebo

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
)

type node struct {
	Address string
	Weight  uint32
}

type upstream struct {
	NodeList []node
}

func InitUpStream(cfg *Config) *upstream {
	list := make([]node, 0, len(cfg.UpStream))
	for _, v := range cfg.UpStream {
		n := node{
			Address: v.Address,
			Weight:  v.Weight,
		}
		list = append(list, n)
	}
	return &upstream{
		NodeList: list,
	}
}

func (u *upstream) Forward(w http.ResponseWriter, r *http.Request) error {
	var totalWeight int
	for i := range u.NodeList {
		totalWeight += int(u.NodeList[i].Weight)
	}
	n := rand.Intn(totalWeight)

	var node *node
	for i := range u.NodeList {
		if n <= int(u.NodeList[i].Weight) {
			node = &u.NodeList[i]
			break
		}
		n -= int(u.NodeList[i].Weight)
	}

	if node == nil {
		return errors.New("failed to calculate weight")
	}

	url := r.URL
	url.Scheme = "http"
	url.Host = node.Address
	req, err := http.NewRequest(r.Method, url.String(), r.Body)
	if err != nil {
		return err
	}

	for k, v := range r.Header {
		for i := range v {
			req.Header.Add(k, v[i])
		}
	}
	req.Header.Set("Host", r.Host)
	req.Header.Set("X-Forwarded-For", r.RemoteAddr)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
	// w.WriteHeader(resp.StatusCode)

	return nil
}
