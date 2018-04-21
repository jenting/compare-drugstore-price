package poya

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

// Crawler implements poya crawler.
func Crawler(search string) {
	resp, err := http.Get("https://tw.search.mall.yahoo.com/search/mall/product?p=" + search)

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
		products := make([]*goquery.Selection, 1)
		doc.Find("div.item").Each(func(i int, s *goquery.Selection) {
			products = append(products, s)
		})

		for _, product := range products {
			if product != nil {
				id, e1 := product.Attr("data-ga-label")
				name, e2 := product.Find("div.wrap").Find("div.srp-pdtitle").ChildrenFiltered("a").Attr("title")
				price := product.Find("div.wrap").Find("div.srp-pdhead").Find("div.srp-pdprice").Find("em.yui3-u").Text()

				if e1 && e2 && price != "" {
					glog.Infof("%v %v %v\n", strings.TrimSpace(id), strings.TrimSpace(name), strings.TrimSpace(price))
				}
			}
		}
	}
}
