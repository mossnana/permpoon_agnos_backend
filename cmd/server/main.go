package main

import (
	passwordHandler "agnos/backend/handlers/password"
	"agnos/backend/pkgs/log"
	passwordValidator "agnos/backend/pkgs/password"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%v:%v@%v:%v/%v", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), "5432", os.Getenv("POSTGRES_DB")))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger := log.New(conn)
	logger.CreateTable(ctx)

	app := gin.Default()
	pwdValidator := passwordValidator.NewMPA()
	pwdHandler := passwordHandler.NewPasswordHandler(pwdValidator, logger)

	app.POST("/api/strong_password_steps", pwdHandler.RecommendMinimumPasswordAction)

	app.Run(os.Getenv("PORT"))
}
