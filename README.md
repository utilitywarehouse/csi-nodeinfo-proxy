# csi-nodeinfo-proxy

This is a CSI Node sidecar to proxy all the grpc traffic to the target, but
intercept `NodeGetInfo` responses to patch the maximum volumes per node
returned value.

This is especially made to work around Netapp/Trident current behaviour:
https://github.com/NetApp/trident/issues/859
