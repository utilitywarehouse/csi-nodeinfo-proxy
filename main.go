package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"path"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// Socket permissions as Trident sets them in:
	// https://github.com/NetApp/trident/blob/master/config/config.go
	// CSIUnixSocketPermissions CSI socket file needs rw access only for user
	CSIUnixSocketPermissions = 0o600
	// CSISocketDirPermissions CSI socket directory needs rwx access only for user
	CSISocketDirPermissions = 0o700
)

var (
	flagListenEndpoint    = flag.String("listen-endpoint", getEnv("CSI_LISTEN_ENDPOINT", "/plugin/csi.sock"), "Endpoint to listen for requests")
	flagProxyEndpoint     = flag.String("proxy-endpoint", getEnv("CSI_PROXY_ENDPOINT", "/plugin/csi-proxy.sock"), "Endpoint to proxy requests")
	flagMaxVolumesPerNode = flag.Int64("max-volumes-per-node", int64(15), "Maximum number of volumes that controller can publish to the node.")
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(
		*flagProxyEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.Dial("unix", addr)
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to backend CSI: %v", err)
	}
	defer conn.Close()
	// Ensure socket dir exists - This is how NetApp/Trident managing
	// sockets and permissions, so let's do the same.
	addr := *flagListenEndpoint
	socketDir := path.Dir(addr)
	if err := os.MkdirAll(socketDir, CSISocketDirPermissions); err != nil {
		log.Fatalf("Failed to make socket dir %s, error: %v", socketDir, err)
	}
	// Ensure socket dir has minimum permissions if it already existed
	if err := os.Chmod(socketDir, CSISocketDirPermissions); err != nil {
		log.Fatalf("Failed to chmod socket dir %s, error: %v", socketDir, err)
	}
	// Remove existing socket
	if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to remove %s, error: %v", addr, err)
	}

	lis, err := net.Listen("unix", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	identityClient := csi.NewIdentityClient(conn)
	nodeClient := csi.NewNodeClient(conn)
	controllerClient := csi.NewControllerClient(conn)
	csi.RegisterIdentityServer(grpcServer, &ProxyIdentityServer{backend: identityClient})
	csi.RegisterNodeServer(grpcServer, &ProxyNodeServer{backend: nodeClient})
	csi.RegisterControllerServer(grpcServer, &ProxyControllerServer{backend: controllerClient})

	log.Printf("CSI proxy listening on %s", *flagListenEndpoint)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
