// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	scheme "awslabs/aws-eks-cluster-controller/pkg/clientset/versioned/scheme"

	v1alpha1 "github.com/awslabs/aws-eks-cluster-controller/pkg/apis/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeGroupsGetter has a method to return a NodeGroupInterface.
// A group's client should implement this interface.
type NodeGroupsGetter interface {
	NodeGroups(namespace string) NodeGroupInterface
}

// NodeGroupInterface has methods to work with NodeGroup resources.
type NodeGroupInterface interface {
	Create(*v1alpha1.NodeGroup) (*v1alpha1.NodeGroup, error)
	Update(*v1alpha1.NodeGroup) (*v1alpha1.NodeGroup, error)
	UpdateStatus(*v1alpha1.NodeGroup) (*v1alpha1.NodeGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodeGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.NodeGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeGroup, err error)
	NodeGroupExpansion
}

// nodeGroups implements NodeGroupInterface
type nodeGroups struct {
	client rest.Interface
	ns     string
}

// newNodeGroups returns a NodeGroups
func newNodeGroups(c *ClusterV1alpha1Client, namespace string) *nodeGroups {
	return &nodeGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeGroup, and returns the corresponding nodeGroup object, and an error if there is any.
func (c *nodeGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeGroup, err error) {
	result = &v1alpha1.NodeGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeGroups that match those selectors.
func (c *nodeGroups) List(opts v1.ListOptions) (result *v1alpha1.NodeGroupList, err error) {
	result = &v1alpha1.NodeGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeGroups.
func (c *nodeGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nodeGroup and creates it.  Returns the server's representation of the nodeGroup, and an error, if there is any.
func (c *nodeGroups) Create(nodeGroup *v1alpha1.NodeGroup) (result *v1alpha1.NodeGroup, err error) {
	result = &v1alpha1.NodeGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodegroups").
		Body(nodeGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeGroup and updates it. Returns the server's representation of the nodeGroup, and an error, if there is any.
func (c *nodeGroups) Update(nodeGroup *v1alpha1.NodeGroup) (result *v1alpha1.NodeGroup, err error) {
	result = &v1alpha1.NodeGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodegroups").
		Name(nodeGroup.Name).
		Body(nodeGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodeGroups) UpdateStatus(nodeGroup *v1alpha1.NodeGroup) (result *v1alpha1.NodeGroup, err error) {
	result = &v1alpha1.NodeGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodegroups").
		Name(nodeGroup.Name).
		SubResource("status").
		Body(nodeGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeGroup and deletes it. Returns an error if one occurs.
func (c *nodeGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeGroup.
func (c *nodeGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeGroup, err error) {
	result = &v1alpha1.NodeGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}