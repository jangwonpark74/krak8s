// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application Contexts
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// CreateApplicationContext provides the application create action context.
type CreateApplicationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
	Payload   *ApplicationPostBody
}

// NewCreateApplicationContext parses the incoming request URL and body, performs validations and creates the
// context used by the application controller create action.
func NewCreateApplicationContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateApplicationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateApplicationContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// Accepted sends a HTTP response with status code 202.
func (ctx *CreateApplicationContext) Accepted(r *Application) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/application+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 202, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateApplicationContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateApplicationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateApplicationContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteApplicationContext provides the application delete action context.
type DeleteApplicationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Appid     string
	Projectid string
}

// NewDeleteApplicationContext parses the incoming request URL and body, performs validations and creates the
// context used by the application controller delete action.
func NewDeleteApplicationContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteApplicationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteApplicationContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAppid := req.Params["appid"]
	if len(paramAppid) > 0 {
		rawAppid := paramAppid[0]
		rctx.Appid = rawAppid
	}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteApplicationContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteApplicationContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteApplicationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetApplicationContext provides the application get action context.
type GetApplicationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Appid     string
	Projectid string
}

// NewGetApplicationContext parses the incoming request URL and body, performs validations and creates the
// context used by the application controller get action.
func NewGetApplicationContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetApplicationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetApplicationContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAppid := req.Params["appid"]
	if len(paramAppid) > 0 {
		rawAppid := paramAppid[0]
		rctx.Appid = rawAppid
	}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetApplicationContext) OK(r *Application) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/application+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetApplicationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListApplicationContext provides the application list action context.
type ListApplicationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
	Payload   *ListApplicationPayload
}

// NewListApplicationContext parses the incoming request URL and body, performs validations and creates the
// context used by the application controller list action.
func NewListApplicationContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListApplicationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListApplicationContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// listApplicationPayload is the application list action payload.
type listApplicationPayload struct {
	Namespaceid *string `form:"namespaceid,omitempty" json:"namespaceid,omitempty" xml:"namespaceid,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *listApplicationPayload) Validate() (err error) {
	if payload.Namespaceid == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "namespaceid"))
	}
	return
}

// Publicize creates ListApplicationPayload from listApplicationPayload
func (payload *listApplicationPayload) Publicize() *ListApplicationPayload {
	var pub ListApplicationPayload
	if payload.Namespaceid != nil {
		pub.Namespaceid = *payload.Namespaceid
	}
	return &pub
}

// ListApplicationPayload is the application list action payload.
type ListApplicationPayload struct {
	Namespaceid string `form:"namespaceid" json:"namespaceid" xml:"namespaceid"`
}

// Validate runs the validation rules defined in the design.
func (payload *ListApplicationPayload) Validate() (err error) {
	if payload.Namespaceid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "namespaceid"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ListApplicationContext) OK(r ApplicationCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/application+json; type=collection")
	if r == nil {
		r = ApplicationCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListApplicationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateClusterContext provides the cluster create action context.
type CreateClusterContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
	Payload   *ClusterPostBody
}

// NewCreateClusterContext parses the incoming request URL and body, performs validations and creates the
// context used by the cluster controller create action.
func NewCreateClusterContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateClusterContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateClusterContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// Accepted sends a HTTP response with status code 202.
func (ctx *CreateClusterContext) Accepted(r *Cluster) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/cluster+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 202, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateClusterContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateClusterContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// Conflict sends a HTTP response with status code 409.
func (ctx *CreateClusterContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateClusterContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteClusterContext provides the cluster delete action context.
type DeleteClusterContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid  string
	ResourceID string
}

// NewDeleteClusterContext parses the incoming request URL and body, performs validations and creates the
// context used by the cluster controller delete action.
func NewDeleteClusterContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteClusterContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteClusterContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	paramResourceID := req.Params["resource_id"]
	if len(paramResourceID) > 0 {
		rawResourceID := paramResourceID[0]
		rctx.ResourceID = rawResourceID
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteClusterContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteClusterContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteClusterContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetClusterContext provides the cluster get action context.
type GetClusterContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid  string
	ResourceID string
}

// NewGetClusterContext parses the incoming request URL and body, performs validations and creates the
// context used by the cluster controller get action.
func NewGetClusterContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetClusterContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetClusterContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	paramResourceID := req.Params["resource_id"]
	if len(paramResourceID) > 0 {
		rawResourceID := paramResourceID[0]
		rctx.ResourceID = rawResourceID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetClusterContext) OK(r *Cluster) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/cluster+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetClusterContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// HealthHealthContext provides the health health action context.
type HealthHealthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHealthHealthContext parses the incoming request URL and body, performs validations and creates the
// context used by the health controller health action.
func NewHealthHealthContext(ctx context.Context, r *http.Request, service *goa.Service) (*HealthHealthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := HealthHealthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *HealthHealthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CreateNamespaceContext provides the namespace create action context.
type CreateNamespaceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
	Payload   *CreateNamespacePayload
}

// NewCreateNamespaceContext parses the incoming request URL and body, performs validations and creates the
// context used by the namespace controller create action.
func NewCreateNamespaceContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateNamespaceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateNamespaceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// createNamespacePayload is the namespace create action payload.
type createNamespacePayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createNamespacePayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates CreateNamespacePayload from createNamespacePayload
func (payload *createNamespacePayload) Publicize() *CreateNamespacePayload {
	var pub CreateNamespacePayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateNamespacePayload is the namespace create action payload.
type CreateNamespacePayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateNamespacePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateNamespaceContext) Created(r *Namespace) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/namespace+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateNamespaceContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateNamespaceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateNamespaceContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteNamespaceContext provides the namespace delete action context.
type DeleteNamespaceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Namespaceid string
	Projectid   string
}

// NewDeleteNamespaceContext parses the incoming request URL and body, performs validations and creates the
// context used by the namespace controller delete action.
func NewDeleteNamespaceContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteNamespaceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteNamespaceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramNamespaceid := req.Params["namespaceid"]
	if len(paramNamespaceid) > 0 {
		rawNamespaceid := paramNamespaceid[0]
		rctx.Namespaceid = rawNamespaceid
	}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteNamespaceContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteNamespaceContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteNamespaceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetNamespaceContext provides the namespace get action context.
type GetNamespaceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Namespaceid string
	Projectid   string
}

// NewGetNamespaceContext parses the incoming request URL and body, performs validations and creates the
// context used by the namespace controller get action.
func NewGetNamespaceContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetNamespaceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetNamespaceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramNamespaceid := req.Params["namespaceid"]
	if len(paramNamespaceid) > 0 {
		rawNamespaceid := paramNamespaceid[0]
		rctx.Namespaceid = rawNamespaceid
	}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetNamespaceContext) OK(r *Namespace) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/namespace+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetNamespaceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListNamespaceContext provides the namespace list action context.
type ListNamespaceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
}

// NewListNamespaceContext parses the incoming request URL and body, performs validations and creates the
// context used by the namespace controller list action.
func NewListNamespaceContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListNamespaceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListNamespaceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListNamespaceContext) OK(r NamespaceCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/namespace+json; type=collection")
	if r == nil {
		r = NamespaceCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateProjectContext provides the project create action context.
type CreateProjectContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateProjectPayload
}

// NewCreateProjectContext parses the incoming request URL and body, performs validations and creates the
// context used by the project controller create action.
func NewCreateProjectContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateProjectContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateProjectContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createProjectPayload is the project create action payload.
type createProjectPayload struct {
	// name of project
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createProjectPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 2, true))
		}
	}
	return
}

// Publicize creates CreateProjectPayload from createProjectPayload
func (payload *createProjectPayload) Publicize() *CreateProjectPayload {
	var pub CreateProjectPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateProjectPayload is the project create action payload.
type CreateProjectPayload struct {
	// name of project
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateProjectPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if utf8.RuneCountInString(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, utf8.RuneCountInString(payload.Name), 2, true))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateProjectContext) Created(r *Project) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/project+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateProjectContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateProjectContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteProjectContext provides the project delete action context.
type DeleteProjectContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
}

// NewDeleteProjectContext parses the incoming request URL and body, performs validations and creates the
// context used by the project controller delete action.
func NewDeleteProjectContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteProjectContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteProjectContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteProjectContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteProjectContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProjectContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetProjectContext provides the project get action context.
type GetProjectContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Projectid string
}

// NewGetProjectContext parses the incoming request URL and body, performs validations and creates the
// context used by the project controller get action.
func NewGetProjectContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetProjectContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetProjectContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectid := req.Params["projectid"]
	if len(paramProjectid) > 0 {
		rawProjectid := paramProjectid[0]
		rctx.Projectid = rawProjectid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetProjectContext) OK(r *Project) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/project+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetProjectContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetProjectContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListProjectContext provides the project list action context.
type ListProjectContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListProjectContext parses the incoming request URL and body, performs validations and creates the
// context used by the project controller list action.
func NewListProjectContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListProjectContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListProjectContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProjectContext) OK(r ProjectCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/project+json; type=collection")
	if r == nil {
		r = ProjectCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
