package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// Credential struct to store json response
type Credential struct {
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

func getPrivateIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Error getting private IP address")
}

func getCredential(url, roleARN string, duration, window int) ([]byte, error) {

	var err error

	client := http.Client{
		Timeout: time.Second * 10,
	}

	if roleARN == "" {
		log.Fatalln(err)
	}

	privateIP, err := getPrivateIP()

	if err != nil {
		log.Fatalln(err)
	}

	hostname, err := os.Hostname()

	if err != nil {
		log.Fatalln(errors.New("Error getting hostname"))
	}

	request, err := json.Marshal(map[string]interface{}{
		"assumed_role_arn": roleARN,
		"token_duration":   duration,
		"expiry_window":    window,
		"private_ip":       privateIP,
		"hostname":         hostname,
	})
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(request))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var creds *Credential

	json.NewDecoder(res.Body).Decode(&creds)

	return json.Marshal(creds)

}
