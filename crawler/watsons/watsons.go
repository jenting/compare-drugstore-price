package watsons

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"

	"github.com/jenting/compare-drugstore-price/data"
)

// Crawler implements watsons crawler.
func Crawler(search string, timeout time.Duration) data.ProductInfoList {
	// Add HTTP timeout mechanism.
	c := &http.Client{
		Timeout: timeout,
	}
	resp, err := c.Get("http://www.watsons.com.tw/search2?text=" + search)
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

	var rets data.ProductInfoList
	product := doc.Find("div.productNameInfo")
	product.Each(func(index int, sql *goquery.Selection) {
		name := sql.Find("div.h1")
		priceS := sql.Find("div.h2").Text()
		id, _ := sql.Find("div.btn-add-to-bag").Attr("data-code")

		priceS = strings.Trim(strings.TrimSpace(priceS), "$")
		price, _ := strconv.Atoi(priceS)

		p := data.ProductInfo{
			ID:    strings.TrimSpace(id),
			Name:  strings.TrimSpace(name.Text()),
			Price: price,
			Shop:  data.WATSONS_SHOP,
		}
		rets = append(rets, p)
		//glog.Infof("%v %v %v\n", strings.TrimSpace(id), strings.TrimSpace(name.Text()), strings.TrimSpace(price.Text()))
	})

	return rets
}
