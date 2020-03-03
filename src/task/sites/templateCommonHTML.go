/*
  Package sites for parse site template
*/

package sites

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"

	cm "../../common"
)

// ParseInfoCommonHTML for common html template which do request get returns pointer of ProInfo instance
func (s *SiteService) ParseInfoCommonHTML(pageURL string, doc *goquery.Document, orderDoc *goquery.Document, labels *cm.LabelsParse) *cm.ProInfo {
	var pi cm.ProInfo
	pi.PageURL = pageURL
	pi.Template = "templateCommonHTML"

	// get page html document
	if doc == nil {
		log.WithFields(log.Fields{
			"pageURL":	pageURL,
		}).Error("doc is nil, log by templateCommonHTML")

		return &pi
	}

	// cover image
	imageDir := cm.ImageDir + time.Now().Format("2006/01/02")
	pi.Cover = s.ParseCoverImagesHTML(doc, pageURL, imageDir, labels.Cover)
	//pi.Images = images
	//if len(pi.Images) > 0 {
	//	pi.Cover = pi.Images[0].URL
	//}

	// title
	pi.Title = s.ParseTitleHTML(doc, labels.Title)

	// price
	pi.Price = s.ParsePriceHTML(doc, labels.Price)

	// description
	pi.Desc = s.parseHTMLDesc(doc, pageURL, imageDir, labels.Desc)

	if orderDoc == nil {  // main and order at the same page or can find order page
		// specifications
		pi.Spec = *s.parseHTMLSpec(doc, pageURL, imageDir, labels.Spec)

		// set meal
		pi.Set = s.parseHTMLSet(doc, pageURL, imageDir, labels.Set)
	} else {  // order at another page, request order page to get order document
		// specifications
		pi.Spec = *s.parseHTMLSpec(orderDoc, pageURL, imageDir, labels.Spec)

		// set meal
		pi.Set = s.parseHTMLSet(orderDoc, pageURL, imageDir, labels.Set)
	}

	return &pi
}