package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-docker-crud/internal/handler"
	"go-docker-crud/internal/repository"
	"go-docker-crud/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// データベース接続の設定
	dsn := "host=db user=postgres dbname=crudapp password=secret sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}

	// スキーマのマイグレーション
	err = db.AutoMigrate(&repository.User{})
	if err != nil {
		log.Fatal("スキーマのマイグレーションに失敗しました:", err)
	}

	// リポジトリ、サービス、ハンドラー層の初期化
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ルーターとルートの設定
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")

	// HTTPサーバーの起動
	log.Println("サーバーがポート8080で実行中です...")
	http.ListenAndServe(":8080", r)
}
