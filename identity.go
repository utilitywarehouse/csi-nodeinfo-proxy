package main

import (
	"context"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
)

// Identity Proxy
type ProxyIdentityServer struct {
	csi.UnimplementedIdentityServer
	backend csi.IdentityClient
}

func (p *ProxyIdentityServer) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return p.backend.GetPluginInfo(ctx, req)
}

func (p *ProxyIdentityServer) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	return p.backend.GetPluginCapabilities(ctx, req)
}

func (p *ProxyIdentityServer) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return p.backend.Probe(ctx, req)
}
