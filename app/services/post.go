package services

import (
	"sozluk/app"
	"sozluk/app/models"
	"sozluk/app/repositories"

	"github.com/revel/revel"
)

//PostService struct
type PostService struct {
	UserRepository repositories.UserRepository
	PostRepository repositories.PostRepository
	Validation     *revel.Validation
}

//CreateNewPost validates post model and sends to repository
func (p PostService) CreateNewPost(name string, content string, user *models.UserModel) (*models.PostModel, map[string]*revel.ValidationError, error) {
	var err error
	model := models.CreateNewPost(name, content, user)
	//validate
	v := p.Validate(model)

	//insert db
	if v == nil {
		m, err := p.PostRepository.Create(model)
		if err != nil {
			return model, v, err
		}
		return m, v, err
	}
	return model, v, err

}

//Validate validates user model form
func (p PostService) Validate(m *models.PostModel) map[string]*revel.ValidationError {
	p.Validation.Check(m.Topic.Name,
		revel.Required{},
		revel.MaxSize{90},
		revel.MinSize{4},
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
