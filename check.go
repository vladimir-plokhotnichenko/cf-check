package main

import "net"

func isCloudflare(ip net.IP) bool {
        for _, c := range cidrs {
                if c == "" {
                        continue
                }

                _, cidrNet, _ := net.ParseCIDR(c)
                if cidrNet.Contains(ip) {
                        return true
                }
        }
        return false
}
