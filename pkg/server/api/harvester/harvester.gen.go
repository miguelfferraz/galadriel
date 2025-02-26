// Package harvester provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package harvester

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// GetFederationRelationshipsParams defines parameters for GetFederationRelationships.
type GetFederationRelationshipsParams struct {
	// filter relationships by spireServer
	SpireServer *string `form:"spireServer,omitempty" json:"spireServer,omitempty"`

	// filter relationships by status
	Status *string `form:"status,omitempty" json:"status,omitempty"`

	// filter relationships by status
	FederationGroupId *int64 `form:"federationGroupId,omitempty" json:"federationGroupId,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /FederationRelationship)
	GetFederationRelationships(ctx echo.Context, params GetFederationRelationshipsParams) error

	// (GET /FederationRelationship/{relationshipID})
	GetRelationshipbyID(ctx echo.Context, relationshipID int64) error

	// (PUT /FederationRelationship/{relationshipID})
	UpdateFederatedRelationshipStatus(ctx echo.Context, relationshipID int64) error

	// (PUT /trustBundles/{trustBundleId})
	UpdateTrustBundle(ctx echo.Context, trustBundleId int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFederationRelationships converts echo context to params.
func (w *ServerInterfaceWrapper) GetFederationRelationships(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFederationRelationshipsParams
	// ------------- Optional query parameter "spireServer" -------------

	err = runtime.BindQueryParameter("form", true, false, "spireServer", ctx.QueryParams(), &params.SpireServer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter spireServer: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Optional query parameter "federationGroupId" -------------

	err = runtime.BindQueryParameter("form", true, false, "federationGroupId", ctx.QueryParams(), &params.FederationGroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter federationGroupId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFederationRelationships(ctx, params)
	return err
}

// GetRelationshipbyID converts echo context to params.
func (w *ServerInterfaceWrapper) GetRelationshipbyID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "relationshipID" -------------
	var relationshipID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "relationshipID", runtime.ParamLocationPath, ctx.Param("relationshipID"), &relationshipID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter relationshipID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRelationshipbyID(ctx, relationshipID)
	return err
}

// UpdateFederatedRelationshipStatus converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateFederatedRelationshipStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "relationshipID" -------------
	var relationshipID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "relationshipID", runtime.ParamLocationPath, ctx.Param("relationshipID"), &relationshipID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter relationshipID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateFederatedRelationshipStatus(ctx, relationshipID)
	return err
}

// UpdateTrustBundle converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTrustBundle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "trustBundleId" -------------
	var trustBundleId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "trustBundleId", runtime.ParamLocationPath, ctx.Param("trustBundleId"), &trustBundleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter trustBundleId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTrustBundle(ctx, trustBundleId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/FederationRelationship", wrapper.GetFederationRelationships)
	router.GET(baseURL+"/FederationRelationship/:relationshipID", wrapper.GetRelationshipbyID)
	router.PUT(baseURL+"/FederationRelationship/:relationshipID", wrapper.UpdateFederatedRelationshipStatus)
	router.PUT(baseURL+"/trustBundles/:trustBundleId", wrapper.UpdateTrustBundle)

}
