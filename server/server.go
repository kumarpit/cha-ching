package main

import (
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/plaid"
	"log"
	"net/http"
	"os"
)

var client *plaid.Client

func main() {
	client = plaid.NewClient(plaid.ClientOptions{
		ClientID:    os.Getenv("PLAID_CLIENT_ID"),
		Environment: plaid.Sandbox,
	})

	r := gin.Default()

	r.GET("/create-link-token", createLinkToken)
	r.Run(":8080")
}

func createLinkToken(c *gin.Context) {
	request := plaid.CreateLinkTokenRequest{
		User: plaid.LinkTokenUser{
			ClientUserID: "unique_user_id",
		},
		ClientName: "Your App Name",
		Products:   []plaid.Product{plaid.ProductTransactions},
		Languages:  []string{"en"},
	}

	linkTokenResponse, err := client.CreateLinkToken(c, request)
	if err != nil {
		log.Printf("Error creating link token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"link_token": linkTokenResponse.LinkToken,
	})
}
