package actions

import (
	"errors"
	"fmt"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/buffalo"
	"golang.org/x/crypto/bcrypt"
)

// SigninHandler : Handles signin.
func SigninHandler(c buffalo.Context) error {
	fmt.Println("GET params were:", c.Request().URL.Query())
	users := []models.User{}
	err := models.DB.All(&users)
	if err != nil {
		fmt.Println(err)
	}
	contests := []models.Contest{}
	err = models.DB.All(&contests)
	if err != nil {
		fmt.Println(err)
	}
	questions := []models.Question{}
	err = models.DB.All(&questions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(questions)
	fmt.Println(users)
	fmt.Println(contests)
	username := c.Request().URL.Query().Get("username")
	password := c.Request().URL.Query().Get("password")
	if username != "" && password != "" {
		fmt.Println("IN")
		err := verifyPassword(c, username, password)
		if err != nil {
			return c.Render(403, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		err = GenerateTokenAndSetCookie(c, username)
		if err != nil {
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": "Internal Server error",
			}))
		}
		// Success.
		user, err1 := getUser(username)
		if err1 != nil {
			fmt.Println(err)
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": "Internal Server error",
			}))
		}
		return c.Render(200, r.JSON(user))
	}
	return c.Render(400, r.JSON(map[string]interface{}{
		"message": "Bad request",
	}))
}

func getUser(username string) (models.User, error) {
	users := &[]models.User{}
	user := models.User{}
	err := models.DB.Where("username = (?)", username).All(users)

	if err != nil {
		fmt.Println("getUserInfo error", err)
		return user, errors.New("User doesn't exist")
	}

	fmt.Println((*users)[0])
	return (*users)[0], nil
}

func verifyPassword(c buffalo.Context, username string, password string) error {
	candidateUser := &[]models.User{}
	fmt.Println("ad", models.DB.Where("username = ?", username).All(candidateUser))
	if len((*candidateUser)) == 0 {
		return errors.New("User Not Found")
	}
	user := (*candidateUser)[0]
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return errors.New("Wrong username/password")
	}
	return nil
}
