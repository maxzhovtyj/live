package handler

import (
	"github.com/labstack/echo/v4"
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
	log.Println(email, password)

	//TODO
	//token, err := h.s.UserService.GenerateTokens(email, password)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}

	ctx.SetCookie(&http.Cookie{
		Name:  accessTokenCookie,
		Value: "123",
		Path:  "/",
	})

	ctx.Response().Header().Set("HX-Redirect", "http://localhost:6789/")

	return nil
}

func (h *Handler) Index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", nil)
}

func (h *Handler) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie(accessTokenCookie)
		if err != nil {
			err = c.Redirect(http.StatusFound, "http://localhost:6789/sign-in")
			if err != nil {
				return err
			}

			return err
		}

		return next(c)
	}
}
