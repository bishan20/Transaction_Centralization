package controllers

import (
	"Centralized_transaction/database"
	"Centralized_transaction/models"
	"Centralized_transaction/repository"
	"Centralized_transaction/repository/crud"
	"Centralized_transaction/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	transaction := models.Transaction{}
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewRepositoryTransactionCRUD(db)
	func(transactionsRepository repository.TransactionRepository) {
		posts, err := transactionsRepository.FindAll(transaction)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, posts)
	}(repo)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Transaction{}
	err = json.Unmarshal(body, &post)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewRepositoryTransactionCRUD(db)
	func(transactionsRepository repository.TransactionRepository) {
		post, err = transactionsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

// func GetPost(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	pid, err := strconv.ParseUint(vars["id"], 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()
// 	repo := crud.NewRepositoryTransactionCRUD(db)
// 	func(transactionsRepository repository.TransactionRepository) {
// 		post, err := transactionsRepository.FindById(pid)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusBadRequest, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, post)
// 	}(repo)
// }

// func UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	pid, err := strconv.ParseUint(vars["id"], 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	post := models.Post{}
// 	err = json.Unmarshal(body, &post)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	uid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	if uid != post.AuthorID {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()
// 	repo := crud.NewRepositoryPostsCRUD(db)
// 	func(postsRepository repository.PostRepository) {
// 		rows, err := postsRepository.Update(pid, post)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusBadRequest, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, rows)
// 	}(repo)
// }

// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	pid, err := strconv.ParseUint(vars["id"], 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	uid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	//fmt.Println("USER:", uid)
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()
// 	repo := crud.NewRepositoryPostsCRUD(db)
// 	func(postsRepository repository.PostRepository) {
// 		_, err := postsRepository.Delete(pid, uid)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusBadRequest, err)
// 			return
// 		}
// 		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
// 		responses.JSON(w, http.StatusNoContent, "")
// 	}(repo)
// }
