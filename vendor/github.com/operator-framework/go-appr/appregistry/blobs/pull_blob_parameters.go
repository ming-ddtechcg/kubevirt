// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPullBlobParams creates a new PullBlobParams object
// with the default values initialized.
func NewPullBlobParams() *PullBlobParams {
	var ()
	return &PullBlobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPullBlobParamsWithTimeout creates a new PullBlobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPullBlobParamsWithTimeout(timeout time.Duration) *PullBlobParams {
	var ()
	return &PullBlobParams{

		timeout: timeout,
	}
}

// NewPullBlobParamsWithContext creates a new PullBlobParams object
// with the default values initialized, and the ability to set a context for a request
func NewPullBlobParamsWithContext(ctx context.Context) *PullBlobParams {
	var ()
	return &PullBlobParams{

		Context: ctx,
	}
}

// NewPullBlobParamsWithHTTPClient creates a new PullBlobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPullBlobParamsWithHTTPClient(client *http.Client) *PullBlobParams {
	var ()
	return &PullBlobParams{
		HTTPClient: client,
	}
}

/*PullBlobParams contains all the parameters to send to the API endpoint
for the pull blob operation typically these are written to a http.Request
*/
type PullBlobParams struct {

	/*Digest
	  content digest

	*/
	Digest string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pull blob params
func (o *PullBlobParams) WithTimeout(timeout time.Duration) *PullBlobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pull blob params
func (o *PullBlobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pull blob params
func (o *PullBlobParams) WithContext(ctx context.Context) *PullBlobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pull blob params
func (o *PullBlobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pull blob params
func (o *PullBlobParams) WithHTTPClient(client *http.Client) *PullBlobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pull blob params
func (o *PullBlobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the pull blob params
func (o *PullBlobParams) WithDigest(digest string) *PullBlobParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the pull blob params
func (o *PullBlobParams) SetDigest(digest string) {
	o.Digest = digest
}

// WithNamespace adds the namespace to the pull blob params
func (o *PullBlobParams) WithNamespace(namespace string) *PullBlobParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the pull blob params
func (o *PullBlobParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the pull blob params
func (o *PullBlobParams) WithPackage(packageVar string) *PullBlobParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the pull blob params
func (o *PullBlobParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WriteToRequest writes these params to a swagger request
func (o *PullBlobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param digest
	if err := r.SetPathParam("digest", o.Digest); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
