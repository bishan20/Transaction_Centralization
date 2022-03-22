package crud

import (
	"Centralized_transaction/models"
	"Centralized_transaction/utils/channels"

	"github.com/jinzhu/gorm"
)

type repositoryTransactiontsCRUD struct {
	db *gorm.DB
}

func NewRepositoryTransactionCRUD(db *gorm.DB) *repositoryTransactiontsCRUD {
	return &repositoryTransactiontsCRUD{db}
}

func (r *repositoryTransactiontsCRUD) Save(post models.Transaction) (models.Transaction, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Transaction{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Transaction{}, err
}

func (r *repositoryTransactiontsCRUD) FindAll(trancitizen models.Transaction) ([]models.Transaction, error) {
	var err error

	posts := []models.Transaction{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Transaction{}).Where("citizenship = ?", trancitizen.Citizenship).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		if len(posts) > 0 {
			for i := range posts {
				err = r.db.Debug().Model(&models.User{}).Where("citizenship = ?", posts[i].Citizenship).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

// func (r *repositoryTransactiontsCRUD) FindById(pid uint64) (models.Transaction, error) {
// 	var err error
// 	post := models.Transaction{}
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		err = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", pid).Find(&post).Error
// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// 		if post.ID != 0 {
// 			err = r.db.Debug().Model(&models.User{}).Where("citizenship = ?", post.Citizenship).Take(&post.Author).Error
// 			if err != nil {
// 				ch <- false
// 				return
// 			}
// 		}
// 		ch <- true
// 	}(done)
// 	if channels.OK(done) {
// 		return post, nil
// 	}
// 	return models.Transaction{}, err
// }

// func (r *repositoryTransactiontsCRUD) Update(pid uint64, post models.Transaction) (int64, error) {
// 	var rs *gorm.DB
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		rs = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", pid).Take(&models.Transaction{}).UpdateColumns(
// 			map[string]interface{}{
// 				"title":      post.Title,
// 				"content":    post.Content,
// 				"updated_at": time.Now(),
// 			},
// 		)
// 		ch <- true
// 	}(done)
// 	if channels.OK(done) {
// 		if rs.Error != nil {
// 			if gorm.IsRecordNotFoundError(rs.Error) {
// 				return 0, errors.New("post not found")
// 			}
// 			return 0, rs.Error
// 		}
// 		return rs.RowsAffected, nil
// 	}
// 	return 0, rs.Error
// }

// func (r *repositoryTransactiontsCRUD) Delete(pid uint64, uid uint32) (int64, error) {

// 	var rs *gorm.DB
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		rs = r.db.Debug().Model(&models.Transaction{}).Where("id = ? and author_id=?", pid, uid).Take(&models.Transaction{}).Delete(&models.Post{})
// 		ch <- true
// 	}(done)
// 	if channels.OK(done) {
// 		if rs.Error != nil {
// 			return 0, rs.Error
// 		}
// 		return rs.RowsAffected, nil
// 	}
// 	return 0, rs.Error
// }
