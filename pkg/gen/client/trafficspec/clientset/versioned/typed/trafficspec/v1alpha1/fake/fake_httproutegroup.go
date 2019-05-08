/*
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
	v1alpha1 "github.com/deislabs/smi-sdk-go/pkg/apis/trafficspec/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPRouteGroups implements HTTPRouteGroupInterface
type FakeHTTPRouteGroups struct {
	Fake *FakeSmispecV1alpha1
	ns   string
}

var httproutegroupsResource = schema.GroupVersionResource{Group: "smispec.io", Version: "v1alpha1", Resource: "httproutegroups"}

var httproutegroupsKind = schema.GroupVersionKind{Group: "smispec.io", Version: "v1alpha1", Kind: "HTTPRouteGroup"}

// Get takes name of the hTTPRouteGroup, and returns the corresponding hTTPRouteGroup object, and an error if there is any.
func (c *FakeHTTPRouteGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.HTTPRouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httproutegroupsResource, c.ns, name), &v1alpha1.HTTPRouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPRouteGroup), err
}

// List takes label and field selectors, and returns the list of HTTPRouteGroups that match those selectors.
func (c *FakeHTTPRouteGroups) List(opts v1.ListOptions) (result *v1alpha1.HTTPRouteGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httproutegroupsResource, httproutegroupsKind, c.ns, opts), &v1alpha1.HTTPRouteGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HTTPRouteGroupList{ListMeta: obj.(*v1alpha1.HTTPRouteGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.HTTPRouteGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPRouteGroups.
func (c *FakeHTTPRouteGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httproutegroupsResource, c.ns, opts))

}

// Create takes the representation of a hTTPRouteGroup and creates it.  Returns the server's representation of the hTTPRouteGroup, and an error, if there is any.
func (c *FakeHTTPRouteGroups) Create(hTTPRouteGroup *v1alpha1.HTTPRouteGroup) (result *v1alpha1.HTTPRouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httproutegroupsResource, c.ns, hTTPRouteGroup), &v1alpha1.HTTPRouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPRouteGroup), err
}

// Update takes the representation of a hTTPRouteGroup and updates it. Returns the server's representation of the hTTPRouteGroup, and an error, if there is any.
func (c *FakeHTTPRouteGroups) Update(hTTPRouteGroup *v1alpha1.HTTPRouteGroup) (result *v1alpha1.HTTPRouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httproutegroupsResource, c.ns, hTTPRouteGroup), &v1alpha1.HTTPRouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPRouteGroup), err
}

// Delete takes name of the hTTPRouteGroup and deletes it. Returns an error if one occurs.
func (c *FakeHTTPRouteGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httproutegroupsResource, c.ns, name), &v1alpha1.HTTPRouteGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPRouteGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httproutegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HTTPRouteGroupList{})
	return err
}

// Patch applies the patch and returns the patched hTTPRouteGroup.
func (c *FakeHTTPRouteGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HTTPRouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httproutegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HTTPRouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPRouteGroup), err
}
