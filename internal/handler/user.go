package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/pkg/templates"
	"log"
	"net/http"
	"time"
)

const accessTokenCookie = "Access-Token"

func (h *Handler) SignInPage(ctx echo.Context) error {
	return templates.SignIn().Render(context.Background(), ctx.Response().Writer)
}

func (h *Handler) SignIn(ctx echo.Context) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	token, err := h.s.User.GenerateTokens(email, password)
	if err != nil {
		err = ctx.String(http.StatusBadRequest, err.Error())
		if err != nil {
			return err
		}

		return nil
	}

	ctx.SetCookie(&http.Cookie{
		Name:  accessTokenCookie,
		Value: token,
		Path:  "/",
	})

	ctx.Response().Header().Set("HX-Redirect", "/")

	return nil
}

func (h *Handler) SignOut(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:    accessTokenCookie,
		Path:    "/",
		Expires: time.Now(),
		MaxAge:  -1,
	})

	ctx.Response().Header().Set("HX-Redirect", "/sign-in")

	return nil
}

func (h *Handler) SignUpPage(ctx echo.Context) error {
	return templates.SignUp().Render(context.Background(), ctx.Response().Writer)
}

func (h *Handler) SignUp(ctx echo.Context) error {
	firstName := ctx.FormValue("firstName")
	lastName := ctx.FormValue("lastName")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	repeatPassword := ctx.FormValue("repeat-password")

	if password != repeatPassword {
		log.Println(password, repeatPassword)
		err := ctx.String(http.StatusBadRequest, "паролі не співпадають")
		if err != nil {
			return err
		}
	}

	err := h.s.User.CreateUser(models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	})
	if err != nil {
		return err
	}

	token, err := h.s.User.GenerateTokens(email, password)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:  accessTokenCookie,
		Value: token,
		Path:  "/",
	})

	ctx.Response().Header().Set("HX-Redirect", "/")

	return nil
}

func (h *Handler) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		t, err := ctx.Cookie(accessTokenCookie)
		if err != nil {
			err = ctx.Redirect(http.StatusFound, "/sign-in")
			if err != nil {
				return err
			}

			return err
		}

		u, err := h.s.ParseToken(t.Value)
		if err != nil {
			ctx.SetCookie(&http.Cookie{
				Name:    accessTokenCookie,
				Path:    "/",
				Expires: time.Now(),
				MaxAge:  -1,
			})

			err = ctx.Redirect(http.StatusFound, "/sign-in")
			if err != nil {
				return err
			}

			return err
		}

		ctx.Set("currentUser", u)

		return next(ctx)
	}
}

func (h *Handler) getUserFromContext(ctx echo.Context) db.User {
	return ctx.Get("currentUser").(db.User)
}
