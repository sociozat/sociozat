package repositories

import (
	"sociozat/app"
	"sociozat/app/models"
	"github.com/jinzhu/gorm"
)

type PostRepository struct{}

//Create add new post to db
func (c PostRepository) Create(p *models.PostModel) (*models.PostModel, error) {

	var err error
	if err := app.DB.Create(&p).Error; err != nil {
		return p, err
	}

	return p, err
}

//Update post by id
func (c PostRepository) Update(p *models.PostModel) (*models.PostModel, error) {

	var err error
	if err := app.DB.Save(&p).Error; err != nil {
		return p, err
	}

	return p, err
}


func (c PostRepository) Find(limit int) ([]models.PostModel, error) {

    posts := []models.PostModel{}
    var err error

    if err := app.DB.Limit(limit).Order("id desc").Preload("Topic").Preload("User").Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err

}

func (c PostRepository) FindByID(id int) (*models.PostModel, error) {
	post := models.PostModel{}

	var err error
	if err := app.DB.Where(&post).Preload("User").First(&post, id).Error; err != nil {
		return &post, err
	}

	//add topic channels
	topic := models.TopicModel{}
	app.DB.Where("id = ?", post.TopicID).Preload("Channels").Find(&topic)
	post.Topic = topic

	return &post, err
}

func (c PostRepository) SaveAction(post int, action string, user uint) (*models.PostActionResponse, error) {
	p := models.PostActionModel{}

    err := app.DB.Transaction(func(tx *gorm.DB) error {
      // do some database operations in the transaction (use 'tx' from this point, not 'db')
      tx.Where("post_id = ?", post).Where("user_id = ?", user).First(&p)

      p.UserID = user
      p.PostID = post

      if p.ID > 0 {
        if action != p.Action && action == "like" {
            tx.Exec("UPDATE posts SET likes = likes + 1, dislikes = dislikes - 1 WHERE id = ?", post)
        }

        if action != p.Action && action == "dislike" {
            tx.Exec("UPDATE posts SET likes = likes - 1, dislikes = dislikes + 1 WHERE id = ?", post)
        }

        p.Action = action
        if err :=  tx.Save(&p).Error; err != nil {
            return err
        }

      }else{
        p.Action = action
        if err := tx.Create(&p).Error; err != nil {
            return err
        }

        if action == "like" {
            tx.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", post)
        }

        if action == "dislike" {
            tx.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?", post)
        }
      }

      return nil
    })

    postModel := models.PostModel{}
    app.DB.First(&postModel, post)

    actionResponse := models.PostActionResponse{}
    actionResponse.Likes = postModel.Likes
    actionResponse.Dislikes = postModel.Dislikes

    //return
	return &actionResponse, err;
}

