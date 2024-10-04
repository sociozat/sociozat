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
    var sitemaps []models.Sitemap

    //find total topic pages
    var topic []models.TopicModel
    var count int
    app.DB.Model(&topic).Count(&count)
    page := math.Ceil(float64(count)/float64(1000))
    for i := 1; i <= int(page); i++ {
        url := models.Sitemap{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/sitemap/topics/" + strconv.Itoa(i), // Customize URL format as needed
            LastMod:  time.Now().Format(time.RFC3339),        // Use topic's creation date as last modification date
        }
        sitemaps = append(sitemaps, url)
    }

    //find total topic pages
    var channels []models.ChannelModel
    app.DB.Model(&channels).Count(&count)
    page = math.Ceil(float64(count)/float64(1000))
    for i := 1; i <= int(page); i++ {
        url := models.Sitemap{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/sitemap/channels/" + strconv.Itoa(i), // Customize URL format as needed
            LastMod:  time.Now().Format(time.RFC3339),        // Use topic's creation date as last modification date
        }
        sitemaps = append(sitemaps, url)
    }

    //find total user pages
    var users []models.UserModel
    app.DB.Model(&users).Count(&count)
    page = math.Ceil(float64(count)/float64(1000))
    for i := 1; i <= int(page); i++ {
        url := models.Sitemap{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/sitemap/users/" + strconv.Itoa(i), // Customize URL format as needed
            LastMod:  time.Now().Format(time.RFC3339),        // Use topic's creation date as last modification date
        }
        sitemaps = append(sitemaps, url)
    }

    // Generate XML response
    sitemapIndex := models.SitemapIndex{
        XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
        Sitemaps:  sitemaps,
    }

    // Render XML response
    c.Response.ContentType = "application/xml"
    return c.RenderXML(sitemapIndex)
}

// Topics generates a paginated sitemap for topics
func (c Sitemap) Topics(page int) revel.Result {
    const pageSize = 1000 // Number of topics per page

    // Calculate offset based on page number
    offset := (page - 1) * pageSize

    // Fetch topics for the specified page
    var topics []models.TopicModel
    if err := app.DB.Order("updated_at desc").Offset(offset).Limit(pageSize).Find(&topics).Error; err != nil {
        // Handle error (e.g., log error, return error response)
        c.Response.Status = http.StatusInternalServerError
        return c.RenderText("Error fetching topics")
    }

    freq := "weekly";
    if page < 2 {
        freq = "daily"
    }
    // Prepare URLs for the sitemap
    var urls []models.SitemapURL
    for _, topic := range topics {

//         if i < 50 {
//             freq : "hourly"
//         }

        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/")  + "/t/" + topic.Slug, // Customize URL format as needed
            LastMod:  topic.UpdatedAt.Format(time.RFC3339),        // Use topic's creation date as last modification date
            ChangeFreq: freq,                                  // Customize change frequency
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
    if err := app.DB.Order("updated_at desc").Offset(offset).Limit(pageSize).Find(&channels).Error; err != nil {
        // Handle error (e.g., log error, return error response)
        c.Response.Status = http.StatusInternalServerError
        return c.RenderText("Error fetching channels")
    }

    freq := "weekly";
    if page < 2 {
        freq = "daily"
    }
    // Prepare URLs for the sitemap
    var urls []models.SitemapURL
    for _, channel := range channels {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/") + "/c/" + channel.Slug, // Customize URL format as needed
            LastMod:  channel.CreatedAt.Format(time.RFC3339),           // Use channel's creation date as last modification date
            ChangeFreq: freq,                                       // Customize change frequency
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

func (c Sitemap) Users(page int) revel.Result {
    const pageSize = 1000 // Number of channels per page

    // Calculate offset based on page number
    offset := (page - 1) * pageSize


    // Fetch channels for the specified page
    var users []models.UserModel
    if err := app.DB.Order("updated_at desc").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
        // Handle error (e.g., log error, return error response)
        c.Response.Status = http.StatusInternalServerError
        return c.RenderText("Error fetching users")
    }

    // Prepare URLs for the sitemap
    var urls []models.SitemapURL
    for _, user := range users {
        url := models.SitemapURL{
            Loc:      revel.Config.StringDefault("app.url", "/") + "/u/" + user.Slug, // Customize URL format as needed
            LastMod:  user.CreatedAt.Format(time.RFC3339),           // Use user's creation date as last modification date
            ChangeFreq: "daily",                                       // Customize change frequency
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

func (c Sitemap) RobotsTxt() revel.Result {
    c.ViewArgs["sitemap"] = revel.Config.StringDefault("app.url", "/") + "/sitemap.xml"
    c.Response.ContentType = "text/plain"
    return c.RenderTemplate("Sitemap/RobotsTxt.html")
}