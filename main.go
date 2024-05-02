package main

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		JWTToken string `json:"jwt_token"`
	}

	ProfileResponse struct {
		Email string `json:"email"`
	}

	User struct {
		Email    string
		password string
	}
)

var (
	webApiPort = ":8080"

	users = map[string]User{}

	secretKey = []byte("qwerty123456")

	contextKeyUser = "user"
)

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	webApp.Post("/signup", func(c *fiber.Ctx) error {
		reqStruct := SignUpRequest{}
		if err := c.BodyParser(&reqStruct); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}

		 if _, exist := users[reqStruct.Email]; exist {
			return c.Status(409).SendString("User already exists")
		}

		users[reqStruct.Email] = User{
			Email:    reqStruct.Email,
			password: reqStruct.Password,
		}

		return c.SendStatus(200)
	})

	webApp.Post("/signin", func(c *fiber.Ctx) error {
		reqStruct := SignInRequest{}
		if err := c.BodyParser(&reqStruct); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}

		user, exist := users[reqStruct.Email]

		if !exist {
			return c.Status(422).SendString("User does not exist")
		}

		if reqStruct.Password != user.password {
			return c.SendStatus(422)
		}

		payload := jwt.MapClaims{
			"sub": user.Email,
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		t, err := token.SignedString(secretKey)

		if err != nil {
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(SignInResponse{JWTToken: t})
	})

	authorizedGroup := webApp.Group("")
	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: secretKey,
		},
		ContextKey: contextKeyUser,
	}))

	authorizedGroup.Get("/profile", func(c *fiber.Ctx) error {
		jwtToken, ok := c.Context().Value(contextKeyUser).(*jwt.Token)

		if !ok {
			return c.SendStatus(401)
		}

		payload, ok := jwtToken.Claims.(jwt.MapClaims)

		if !ok {
			return c.SendStatus(401)
		}

		userInfo, ok := users[payload["sub"].(string)]

		if !ok {
			return c.SendStatus(422)
		}

		return c.Status(200).JSON(ProfileResponse{Email: userInfo.Email})
	})
	// END

	logrus.Fatal(webApp.Listen(webApiPort))
}
