package post_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/require"
	"net/http"
	"pact-contract-go/client/post"
	"testing"
)

func TestClient(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "MyConsumer",
		Provider: "MyProvider",
		Host:     "localhost",
	}

	posts := []post.Post{
		{
			UserID: 1,
			ID:     3,
			Title:  "ea molestias quasi exercitationem repellat qui ipsa sit aut",
			Body:   "ullam et saepe reiciendis voluptatem ad",
		},
		{
			UserID: 1,
			ID:     4,
			Title:  "eum et est occaecati",
			Body:   "ullam et saepe reiciendis voluptatem adipisci\\nsit amet",
		},
	}

	pact.Setup(true)
	defer pact.Teardown()

	var test = func() (err error) {
		url := fmt.Sprintf("http://localhost:%d/posts", pact.Server.Port)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := http.DefaultClient.Do(req)
		require.Nil(t, err)
		require.NotNil(t, resp)
		return
	}

	t.Run("success", func(t *testing.T) {
		pact.AddInteraction().
			Given("Server is up").
			UponReceiving(" GET /posts").
			WithRequest(dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String("/posts"),
				Headers: dsl.MapMatcher{
					fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
				},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusOK,
				Body:   dsl.Like(posts),
			})

		require.Nil(t, pact.Verify(test))
	})

	require.Nil(t, pact.WritePact())
}
