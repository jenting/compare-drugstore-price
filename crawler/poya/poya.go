package poya

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/data"
)

// Crawler implements poya crawler.
func Crawler(search string, timeout time.Duration) []data.ProductInfo {
	// Add HTTP timeout mechanism.
	c := &http.Client{
		Timeout: timeout,
	}
	resp, err := c.Get("https://tw.search.mall.yahoo.com/search/mall/product?p=" + search)

	if err != nil {
		glog.Fatal(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		glog.Errorf("http status code error: %v", resp.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil
	}

	products := make([]*goquery.Selection, 1)
	doc.Find("div.item").Each(func(i int, s *goquery.Selection) {
		products = append(products, s)
	})

	var rets []data.ProductInfo
	for _, product := range products {
		if product == nil {
			continue
		}

		id, e1 := product.Attr("data-ga-label")
		name, e2 := product.Find("div.wrap").Find("div.srp-pdtitle").ChildrenFiltered("a").Attr("title")
		priceS := product.Find("div.wrap").Find("div.srp-pdhead").Find("div.srp-pdprice").Find("em.yui3-u").Text()
		if e1 && e2 && priceS != "" {
			priceS = strings.Trim(strings.TrimSpace(priceS), "$")
			price, _ := strconv.Atoi(priceS)

			p := data.ProductInfo{
				ID:    strings.TrimSpace(id),
				Name:  strings.TrimSpace(name),
				Price: price,
				Shop:  data.POYA_SHOP,
			}
			rets = append(rets, p)
			//glog.Infof("%v %v %v\n", strings.TrimSpace(id), strings.TrimSpace(name), strings.TrimSpace(price))
		}
	}

	return rets
}
