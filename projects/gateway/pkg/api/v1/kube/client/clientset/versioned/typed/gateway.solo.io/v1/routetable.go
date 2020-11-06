/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	scheme "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RouteTablesGetter has a method to return a RouteTableInterface.
// A group's client should implement this interface.
type RouteTablesGetter interface {
	RouteTables(namespace string) RouteTableInterface
}

// RouteTableInterface has methods to work with RouteTable resources.
type RouteTableInterface interface {
	Create(ctx context.Context, routeTable *v1.RouteTable, opts metav1.CreateOptions) (*v1.RouteTable, error)
	Update(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (*v1.RouteTable, error)
	UpdateStatus(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (*v1.RouteTable, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RouteTable, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RouteTableList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RouteTable, err error)
	RouteTableExpansion
}

// routeTables implements RouteTableInterface
type routeTables struct {
	client rest.Interface
	ns     string
}

// newRouteTables returns a RouteTables
func newRouteTables(c *GatewayV1Client, namespace string) *routeTables {
	return &routeTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the routeTable, and returns the corresponding routeTable object, and an error if there is any.
func (c *routeTables) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RouteTable, err error) {
	result = &v1.RouteTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RouteTables that match those selectors.
func (c *routeTables) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RouteTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RouteTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested routeTables.
func (c *routeTables) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("routetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a routeTable and creates it.  Returns the server's representation of the routeTable, and an error, if there is any.
func (c *routeTables) Create(ctx context.Context, routeTable *v1.RouteTable, opts metav1.CreateOptions) (result *v1.RouteTable, err error) {
	result = &v1.RouteTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("routetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(routeTable).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a routeTable and updates it. Returns the server's representation of the routeTable, and an error, if there is any.
func (c *routeTables) Update(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (result *v1.RouteTable, err error) {
	result = &v1.RouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routetables").
		Name(routeTable.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(routeTable).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *routeTables) UpdateStatus(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (result *v1.RouteTable, err error) {
	result = &v1.RouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routetables").
		Name(routeTable.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(routeTable).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the routeTable and deletes it. Returns an error if one occurs.
func (c *routeTables) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routetables").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *routeTables) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routetables").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched routeTable.
func (c *routeTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RouteTable, err error) {
	result = &v1.RouteTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("routetables").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
