// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumLoadBalancerIPPoolLister helps list CiliumLoadBalancerIPPools.
// All objects returned here must be treated as read-only.
type CiliumLoadBalancerIPPoolLister interface {
	// List lists all CiliumLoadBalancerIPPools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.CiliumLoadBalancerIPPool, err error)
	// Get retrieves the CiliumLoadBalancerIPPool from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.CiliumLoadBalancerIPPool, error)
	CiliumLoadBalancerIPPoolListerExpansion
}

// ciliumLoadBalancerIPPoolLister implements the CiliumLoadBalancerIPPoolLister interface.
type ciliumLoadBalancerIPPoolLister struct {
	indexer cache.Indexer
}

// NewCiliumLoadBalancerIPPoolLister returns a new CiliumLoadBalancerIPPoolLister.
func NewCiliumLoadBalancerIPPoolLister(indexer cache.Indexer) CiliumLoadBalancerIPPoolLister {
	return &ciliumLoadBalancerIPPoolLister{indexer: indexer}
}

// List lists all CiliumLoadBalancerIPPools in the indexer.
func (s *ciliumLoadBalancerIPPoolLister) List(selector labels.Selector) (ret []*v2alpha1.CiliumLoadBalancerIPPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CiliumLoadBalancerIPPool))
	})
	return ret, err
}

// Get retrieves the CiliumLoadBalancerIPPool from the index for a given name.
func (s *ciliumLoadBalancerIPPoolLister) Get(name string) (*v2alpha1.CiliumLoadBalancerIPPool, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("ciliumloadbalancerippool"), name)
	}
	return obj.(*v2alpha1.CiliumLoadBalancerIPPool), nil
}
