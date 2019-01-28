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

package fake

import (
	v1alpha1 "github.com/awslabs/aws-eks-cluster-controller/pkg/apis/components/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoles implements ClusterRoleInterface
type FakeClusterRoles struct {
	Fake *FakeComponentsV1alpha1
	ns   string
}

var clusterrolesResource = schema.GroupVersionResource{Group: "components.eks.amazonaws.com", Version: "v1alpha1", Resource: "clusterroles"}

var clusterrolesKind = schema.GroupVersionKind{Group: "components.eks.amazonaws.com", Version: "v1alpha1", Kind: "ClusterRole"}

// Get takes name of the clusterRole, and returns the corresponding clusterRole object, and an error if there is any.
func (c *FakeClusterRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterrolesResource, c.ns, name), &v1alpha1.ClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors.
func (c *FakeClusterRoles) List(opts v1.ListOptions) (result *v1alpha1.ClusterRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterrolesResource, clusterrolesKind, c.ns, opts), &v1alpha1.ClusterRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterRoleList{ListMeta: obj.(*v1alpha1.ClusterRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoles.
func (c *FakeClusterRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterrolesResource, c.ns, opts))

}

// Create takes the representation of a clusterRole and creates it.  Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *FakeClusterRoles) Create(clusterRole *v1alpha1.ClusterRole) (result *v1alpha1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterrolesResource, c.ns, clusterRole), &v1alpha1.ClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// Update takes the representation of a clusterRole and updates it. Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *FakeClusterRoles) Update(clusterRole *v1alpha1.ClusterRole) (result *v1alpha1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterrolesResource, c.ns, clusterRole), &v1alpha1.ClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRoles) UpdateStatus(clusterRole *v1alpha1.ClusterRole) (*v1alpha1.ClusterRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterrolesResource, "status", c.ns, clusterRole), &v1alpha1.ClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// Delete takes name of the clusterRole and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterrolesResource, c.ns, name), &v1alpha1.ClusterRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterRoleList{})
	return err
}

// Patch applies the patch and returns the patched clusterRole.
func (c *FakeClusterRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterrolesResource, c.ns, name, data, subresources...), &v1alpha1.ClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}
