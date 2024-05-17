package controllers

import (
    "github.com/revel/revel"
    "net/http"
    "time"
    "strconv"
    "math"
    "sociozat/app"
    "sociozat/app/models"
)

type Sitemap struct {
    App
    *revel.Controller
}

// Index generates the main sitemap.xml
func (c Sitemap) Index() revel.Result {
    var urls []models.SitemapURL

    //find total topic pages
    var topic []models.TopicModel
    var count int
    app.DB.Model(&topic).Count(&count)
    page := math.Ceil(float64(count)/float64(1000))
    for i := 1; i <= int(page); i++ {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/sitemap/topics/" + strconv.Itoa(i), // Customize URL format as needed
            LastMod:  time.Now().Format(time.RFC3339),        // Use topic's creation date as last modification date
            ChangeFreq: "weekly",                                  // Customize change frequency
            Priority:   0.8,                                       // Customize priority
        }
        urls = append(urls, url)
    }

    //find total topic pages
    var channels []models.ChannelModel
    app.DB.Model(&channels).Count(&count)
    page = math.Ceil(float64(count)/float64(1000))
    for i := 1; i <= int(page); i++ {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/sitemap/channels/" + strconv.Itoa(i), // Customize URL format as needed
            LastMod:  time.Now().Format(time.RFC3339),        // Use topic's creation date as last modification date
            ChangeFreq: "weekly",                                  // Customize change frequency
            Priority:   0.8,                                       // Customize priority
        }
        urls = append(urls, url)
    }

 // Generate XML response
    urlSet := models.SitemapURLSet{
        XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
        URLs:  urls,
    }

    // Render XML response
    c.Response.ContentType = "application/xml"
    return c.RenderXML(urlSet)
}

// Topics generates a paginated sitemap for topics
func (c Sitemap) Topics(page int) revel.Result {
    const pageSize = 1000 // Number of topics per page

    // Calculate offset based on page number
    offset := (page - 1) * pageSize

    // Fetch topics for the specified page
    var topics []models.TopicModel
    if err := app.DB.Offset(offset).Limit(pageSize).Find(&topics).Error; err != nil {
        // Handle error (e.g., log error, return error response)
        c.Response.Status = http.StatusInternalServerError
        return c.RenderText("Error fetching topics")
    }

    // Prepare URLs for the sitemap
    var urls []models.SitemapURL
    for _, topic := range topics {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/t/" + topic.Slug, // Customize URL format as needed
            LastMod:  topic.UpdatedAt.Format(time.RFC3339),        // Use topic's creation date as last modification date
            ChangeFreq: "weekly",                                  // Customize change frequency
            Priority:   0.8,                                       // Customize priority
        }
        urls = append(urls, url)
    }

    // Generate XML response
    urlSet := models.SitemapURLSet{
        XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
        URLs:  urls,
    }

    // Render XML response
    c.Response.ContentType = "application/xml"
    return c.RenderXML(urlSet)
}

// Channels generates a paginated sitemap for channels
func (c Sitemap) Channels(page int) revel.Result {
    const pageSize = 1000 // Number of channels per page

    // Calculate offset based on page number
    offset := (page - 1) * pageSize


    // Fetch channels for the specified page
    var channels []models.ChannelModel
    if err := app.DB.Offset(offset).Limit(pageSize).Find(&channels).Error; err != nil {
        // Handle error (e.g., log error, return error response)
        c.Response.Status = http.StatusInternalServerError
        return c.RenderText("Error fetching channels")
    }

    // Prepare URLs for the sitemap
    var urls []models.SitemapURL
    for _, channel := range channels {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/") + "/c/" + channel.Slug, // Customize URL format as needed
            LastMod:  channel.CreatedAt.Format(time.RFC3339),           // Use channel's creation date as last modification date
            ChangeFreq: "weekly",                                       // Customize change frequency
            Priority:   0.8,                                            // Customize priority
        }
        urls = append(urls, url)
    }

    // Generate XML response
    urlSet := models.SitemapURLSet{
        XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
        URLs:  urls,
    }

    // Render XML response
    c.Response.ContentType = "application/xml"
    return c.RenderXML(urlSet)
}
