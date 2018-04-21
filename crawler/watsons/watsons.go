package watsons

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

// Crawler implements watsons crawler.
func Crawler(search string) {
	resp, err := http.Get("http://www.watsons.com.tw/search2?text=" + search)
	if err != nil {
		glog.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		glog.Errorf("http status code error: %v", resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		glog.Error(err)
	} else {
		product := doc.Find("div.productNameInfo")
		product.Each(func(index int, sql *goquery.Selection) {
			name := sql.Find("div.h1")
			price := sql.Find("div.h2")
			id, _ := sql.Find("div.btn-add-to-bag").Attr("data-code")

			glog.Infof("%v %v %v\n", strings.TrimSpace(id), strings.TrimSpace(name.Text()), strings.TrimSpace(price.Text()))
		})
	}
}
