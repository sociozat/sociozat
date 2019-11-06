package services

import (
	"sociozat/app"
	"sociozat/app/models"
	"sociozat/app/repositories"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

//PostService struct
type PostService struct {
	UserRepository    repositories.UserRepository
	PostRepository    repositories.PostRepository
	ChannelRepository repositories.ChannelRepository
	Validation        *revel.Validation
}

//CreateNewPost validates post model and sends to repository
func (p PostService) CreateNewPost(name string, content string, channels string, user *models.UserModel) (*models.PostModel, map[string]*revel.ValidationError, error) {
	var err error

	var topicChannels = p.GenerateChannels(channels)

	post := models.CreateNewPost(name, content, user)
	post.Topic.Channels = topicChannels

	//validate
	v := p.Validate(post)

	//insert db
	if v == nil {
		m, err := p.PostRepository.Create(post)
		if err != nil {
			return post, v, err
		}
		return m, v, err
	}
	return post, v, err

}

func (p PostService) GenerateChannels(channels string) []models.ChannelModel {
	c := []models.ChannelModel{}
	if channels == "" {
		return c
	}

	var tags = strings.Split(channels, ",")

	for _, v := range tags {
		//check if its not an id create new channel
		j, err := strconv.Atoi(string(v))
		if err != nil {
			//create new channel by name
			channel := models.NewChannel(string(v))
			dbChannel, _ := p.ChannelRepository.Create(channel)
			c = append(c, *dbChannel)
		} else {
			//add channel by id
			dbChannel, _ := p.ChannelRepository.FindByID(j)
			c = append(c, *dbChannel)
		}
	}

	return c
}

func (p PostService) FindByID(id int) (*models.PostModel, error) {
	return p.PostRepository.FindByID(id)
}

//Validate validates post model form
func (p PostService) Validate(m *models.PostModel) map[string]*revel.ValidationError {
	p.Validation.Check(m.Topic.Name,
		revel.Required{},
		revel.MaxSize{90},
		revel.MinSize{2},
	).Message(app.Trans("post.create.validation.name"))

	p.Validation.Check(m.Content,
		revel.Required{},
		revel.MinSize{5},
	).Message(app.Trans("post.create.validation.content"))

	if p.Validation.HasErrors() {
		p.Validation.Keep()
		return p.Validation.ErrorMap()
	}

	return nil
}
