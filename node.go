package main

import (
	"context"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
)

// Node Proxy
type ProxyNodeServer struct {
	csi.UnimplementedNodeServer
	backend csi.NodeClient
}

func (p *ProxyNodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	resp, err := p.backend.NodeGetInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.MaxVolumesPerNode = *flagMaxVolumesPerNode
	return resp, nil
}

func (p *ProxyNodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return p.backend.NodePublishVolume(ctx, req)
}

func (p *ProxyNodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return p.backend.NodeUnpublishVolume(ctx, req)
}

func (p *ProxyNodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return p.backend.NodeStageVolume(ctx, req)
}

func (p *ProxyNodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return p.backend.NodeUnstageVolume(ctx, req)
}

func (p *ProxyNodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return p.backend.NodeGetCapabilities(ctx, req)
}

func (p *ProxyNodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return p.backend.NodeGetVolumeStats(ctx, req)
}

func (p *ProxyNodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return p.backend.NodeExpandVolume(ctx, req)
}
