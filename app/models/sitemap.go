package models

import "encoding/xml"

type SitemapURL struct {
    XMLName    xml.Name `xml:"url"`
    Loc        string   `xml:"loc"`
    LastMod    string   `xml:"lastmod,omitempty"`
    ChangeFreq string   `xml:"changefreq,omitempty"`
    Priority   float64  `xml:"priority,omitempty"`
}

type SitemapURLSet struct {
    XMLName xml.Name `xml:"urlset"`
    XMLNS   string   `xml:"xmlns,attr"`
    URLs    []SitemapURL    `xml:"url"`
}
