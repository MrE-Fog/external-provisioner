/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumegroupsnapshot/v1alpha1"
	scheme "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeGroupSnapshotsGetter has a method to return a VolumeGroupSnapshotInterface.
// A group's client should implement this interface.
type VolumeGroupSnapshotsGetter interface {
	VolumeGroupSnapshots(namespace string) VolumeGroupSnapshotInterface
}

// VolumeGroupSnapshotInterface has methods to work with VolumeGroupSnapshot resources.
type VolumeGroupSnapshotInterface interface {
	Create(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.CreateOptions) (*v1alpha1.VolumeGroupSnapshot, error)
	Update(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.UpdateOptions) (*v1alpha1.VolumeGroupSnapshot, error)
	UpdateStatus(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.UpdateOptions) (*v1alpha1.VolumeGroupSnapshot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumeGroupSnapshot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeGroupSnapshotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeGroupSnapshot, err error)
	VolumeGroupSnapshotExpansion
}

// volumeGroupSnapshots implements VolumeGroupSnapshotInterface
type volumeGroupSnapshots struct {
	client rest.Interface
	ns     string
}

// newVolumeGroupSnapshots returns a VolumeGroupSnapshots
func newVolumeGroupSnapshots(c *GroupsnapshotV1alpha1Client, namespace string) *volumeGroupSnapshots {
	return &volumeGroupSnapshots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeGroupSnapshot, and returns the corresponding volumeGroupSnapshot object, and an error if there is any.
func (c *volumeGroupSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeGroupSnapshot, err error) {
	result = &v1alpha1.VolumeGroupSnapshot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeGroupSnapshots that match those selectors.
func (c *volumeGroupSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeGroupSnapshotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeGroupSnapshotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeGroupSnapshots.
func (c *volumeGroupSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeGroupSnapshot and creates it.  Returns the server's representation of the volumeGroupSnapshot, and an error, if there is any.
func (c *volumeGroupSnapshots) Create(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.CreateOptions) (result *v1alpha1.VolumeGroupSnapshot, err error) {
	result = &v1alpha1.VolumeGroupSnapshot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeGroupSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeGroupSnapshot and updates it. Returns the server's representation of the volumeGroupSnapshot, and an error, if there is any.
func (c *volumeGroupSnapshots) Update(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.UpdateOptions) (result *v1alpha1.VolumeGroupSnapshot, err error) {
	result = &v1alpha1.VolumeGroupSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		Name(volumeGroupSnapshot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeGroupSnapshot).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumeGroupSnapshots) UpdateStatus(ctx context.Context, volumeGroupSnapshot *v1alpha1.VolumeGroupSnapshot, opts v1.UpdateOptions) (result *v1alpha1.VolumeGroupSnapshot, err error) {
	result = &v1alpha1.VolumeGroupSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		Name(volumeGroupSnapshot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeGroupSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeGroupSnapshot and deletes it. Returns an error if one occurs.
func (c *volumeGroupSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeGroupSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeGroupSnapshot.
func (c *volumeGroupSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeGroupSnapshot, err error) {
	result = &v1alpha1.VolumeGroupSnapshot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumegroupsnapshots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
