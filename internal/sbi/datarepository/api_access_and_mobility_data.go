/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/free5gc/openapi"
	udr_context "github.com/free5gc/udr/internal/context"
	"github.com/gin-gonic/gin"
)

// CreateAccessAndMobilityData - Creates and updates the access and mobility exposure data for a UE
func CreateAccessAndMobilityData(c *gin.Context) {
	scopes := []string{"nudr-dr"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && udr_context.UDR_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteAccessAndMobilityData - Deletes the access and mobility exposure data for a UE
func DeleteAccessAndMobilityData(c *gin.Context) {
	scopes := []string{"nudr-dr"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && udr_context.UDR_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// QueryAccessAndMobilityData - Retrieves the access and mobility exposure data for a UE
func QueryAccessAndMobilityData(c *gin.Context) {
	scopes := []string{"nudr-dr"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && udr_context.UDR_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
