package main

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	net_html "golang.org/x/net/html"
)

type UrlMetadata struct {
	Url             string
	StatusCode      int               `json:"StatusCode,omitempty"`
	ResponseHeasers map[string]string `json:"ResponseHeasers,omitempty"`
	ResponseTime    int64             `json:"ResponseTime,omitempty"`
	RegexMatches    []string          `json:"RegexMatches,omitempty"`
	FaviconHash     string            `json:"FaviconHash,omitempty"`
}

type Extractor struct {
	RequestOptions    RequestOptions
	ExtractionOptions ExtractionOptions
}

func (ext Extractor) Extract(url string) (UrlMetadata, error) {
	var metadata = UrlMetadata{Url: url}
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := http.Client{Transport: transCfg}

	req, _ := http.NewRequest(ext.RequestOptions.Method, url, nil)
	for _, header := range ext.RequestOptions.Headers {
		req.Header.Add(strings.Split(header, ":")[0], strings.Split(header, ":")[1])
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return metadata, err
	}
	t := time.Now()
	elapsed := t.Sub(start)

	metadata.StatusCode = resp.StatusCode
	metadata.ResponseTime = elapsed.Milliseconds()
	metadata.ResponseHeasers = map[string]string{}
	for _, header := range ext.ExtractionOptions.Headers {
		metadata.ResponseHeasers[header] = resp.Header.Get(header)
	}
	if ext.ExtractionOptions.FaviconHash {
		metadata.FaviconHash = ext.ExtractFavicon(client, url, resp.Body)
	}

	// TODO: implement regex extraction

	return metadata, nil
}

func (ext Extractor) ExtractFavicon(client http.Client, requestUrl string, bodyReader io.ReadCloser) string {
	var getAttribute func(n *net_html.Node, key string) string
	getAttribute = func(n *net_html.Node, key string) string {

		for _, attr := range n.Attr {

			if attr.Key == key {
				return attr.Val
			}
		}

		return ""
	}

	body, _ := ioutil.ReadAll(bodyReader)
	htmlText := string(body)
	doc, _ := net_html.Parse(strings.NewReader(htmlText))
	var crawler func(*net_html.Node) string
	crawler = func(node *net_html.Node) string {
		if node.Type == net_html.ElementNode && node.Data == "link" && strings.Contains(getAttribute(node, "rel"), "icon") {
			return getAttribute(node, "href")
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			faviconUrl := crawler(child)
			if faviconUrl != "" {
				return faviconUrl
			}
		}
		return ""
	}

	faviconUrl := crawler(doc)
	if faviconUrl == "" {
		faviconUrl = "/favicon.ico" // TODO: should i bruteforce other image extensions?
	}

	if !strings.HasPrefix(faviconUrl, "http") && !strings.HasPrefix(faviconUrl, "//") {
		base, _ := url.Parse(requestUrl)
		favUrl, _ := base.Parse(faviconUrl)
		faviconUrl = favUrl.String()
	}

	req, _ := http.NewRequest(ext.RequestOptions.Method, faviconUrl, nil)
	for _, header := range ext.RequestOptions.Headers {
		req.Header.Add(strings.Split(header, ":")[0], strings.Split(header, ":")[1])
	}

	resp, _ := client.Do(req)
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return ""
	}
	body, _ = ioutil.ReadAll(resp.Body)
	hash := md5.Sum(body)

	return hex.EncodeToString(hash[:])
}
