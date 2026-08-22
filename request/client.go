package request

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry"
)

func NewHttpClient() (*http.Client, error) {
	proxyURL, err := getSystemProxy()
	if err != nil {
		return nil, err
	}
	if proxyURL != nil {
		return &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}, nil
	} else {
		return &http.Client{
			Timeout: 10 * time.Second,
		}, nil
	}

}

// parseProxyServer 处理 Windows 注册表中 ProxyServer 的两种格式：
//   - 简单格式: "127.0.0.1:8888"
//   - 扩展格式: "http=127.0.0.1:8888;https=127.0.0.1:8888"
func parseProxyServer(proxyServer string) string {
	// 不包含 '='，说明是简单格式，直接返回
	if !strings.Contains(proxyServer, "=") {
		return proxyServer
	}

	// 扩展格式：按 ';' 分割，按 '=' 拆分协议和地址
	// 优先取 http= 对应的地址，其次取 https=
	var httpAddr, httpsAddr string
	for _, part := range strings.Split(proxyServer, ";") {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			continue
		}
		protocol, addr := strings.ToLower(kv[0]), kv[1]
		switch protocol {
		case "http":
			httpAddr = addr
		case "https":
			httpsAddr = addr
		}
	}

	// HTTP 代理优先；HTTPS 代理作为 HTTP 隧道代理同样可用
	if httpAddr != "" {
		return httpAddr
	}
	if httpsAddr != "" {
		return httpsAddr
	}

	// 都没取到，兜底返回原始值
	return proxyServer
}

func getSystemProxy() (*url.URL, error) {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Internet Settings`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	proxyEnable, _, err := key.GetIntegerValue("ProxyEnable")
	if err != nil {
		return nil, err
	}

	if proxyEnable == 1 {
		proxyServer, _, err := key.GetStringValue("ProxyServer")
		if err != nil {
			return nil, err
		}
		if proxyServer != "" {
			// 解析扩展格式，如 "http=127.0.0.1:8888;https=127.0.0.1:8888"
			proxyAddr := parseProxyServer(proxyServer)
			proxyURL, err := url.Parse("http://" + proxyAddr)
			if err != nil {
				return nil, err
			}
			return proxyURL, nil
		}
	}

	return nil, nil
}
