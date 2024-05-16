package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/live/internal/models"
	"log"
	"net/http"
)

const accessTokenCookie = "Access-Token"

func (h *Handler) SignInPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "signIn", nil)
}

func (h *Handler) SignIn(ctx echo.Context) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	token, err := h.s.UserService.GenerateTokens(email, password)
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

func (h *Handler) SignUpPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "signUp", nil)
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

	err := h.s.UserService.CreateUser(models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	})
	if err != nil {
		return err
	}

	token, err := h.s.UserService.GenerateTokens(email, password)
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

func (h *Handler) Index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", nil)
}

func (h *Handler) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie(accessTokenCookie)
		if err != nil {
			err = c.Redirect(http.StatusFound, "/sign-in")
			if err != nil {
				return err
			}

			return err
		}

		return next(c)
	}
}
