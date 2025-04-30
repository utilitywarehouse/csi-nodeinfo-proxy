package main

import (
	"context"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
)

type ProxyControllerServer struct {
	csi.UnimplementedControllerServer
	backend csi.ControllerClient
}

func (p *ProxyControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return p.backend.CreateVolume(ctx, req)
}

func (p *ProxyControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return p.backend.DeleteVolume(ctx, req)
}

func (p *ProxyControllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return p.backend.ControllerPublishVolume(ctx, req)
}

func (p *ProxyControllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return p.backend.ControllerUnpublishVolume(ctx, req)
}

func (p *ProxyControllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return p.backend.ValidateVolumeCapabilities(ctx, req)
}

func (p *ProxyControllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return p.backend.ListVolumes(ctx, req)
}

func (p *ProxyControllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return p.backend.GetCapacity(ctx, req)
}

func (p *ProxyControllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return p.backend.ControllerGetCapabilities(ctx, req)
}

func (p *ProxyControllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return p.backend.CreateSnapshot(ctx, req)
}

func (p *ProxyControllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return p.backend.DeleteSnapshot(ctx, req)
}

func (p *ProxyControllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return p.backend.ListSnapshots(ctx, req)
}

func (p *ProxyControllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return p.backend.ControllerExpandVolume(ctx, req)
}

func (p *ProxyControllerServer) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return p.backend.ControllerGetVolume(ctx, req)
}
