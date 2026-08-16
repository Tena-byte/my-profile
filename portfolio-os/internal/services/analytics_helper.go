package services

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return strings.TrimSpace(strings.Split(forwarded, ",")[0])
	}

	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return strings.TrimSpace(realIP)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func AnonymizeIP(ip string) string {
	if ip == "" {
		return "unknown"
	}

	parts := strings.Split(ip, ".")

	if len(parts) == 4 {
		return parts[0] + "." + parts[1] + "." + parts[2] + ".0"
	}

	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")

		if len(parts) > 4 {
			return strings.Join(parts[:4], ":") + "::"
		}

		return ip
	}

	return "unknown"
}

func DetectDevice(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "tablet") ||
		strings.Contains(ua, "ipad") {
		return "tablet"
	}

	if strings.Contains(ua, "mobile") ||
		strings.Contains(ua, "iphone") ||
		strings.Contains(ua, "android") {
		return "mobile"
	}

	return "desktop"
}
