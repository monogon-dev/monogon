package rpc

import (
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	epb "source.monogon.dev/metropolis/proto/ext"
)

// methodInfo is the parsed information for a given RPC method, as configured by
// the metropolis.common.ext.authorization extension.
type methodInfo struct {
	// unauthenticated is true if the method is defined as 'unauthenticated', ie.
	// that all requests should be passed to the gRPC handler without any
	// authentication or authorization performed.
	unauthenticated bool
	// need is a map of permissions that the caller needs to have in order to be
	// allowed to call this method. If not empty, unauthenticated cannot be set to
	// true.
	need map[epb.Permission]bool
}

var (
	// reMethodName matches a /some.service/Method string from
	// {Stream,Unary}ServerInfo.FullMethod.
	reMethodName = regexp.MustCompile(`^/([^/]+)/([^/.]+)$`)
)

// getMethodInfo returns the methodInfo for a given method name, as retrieved
// from grpc.{Stream,Unary}ServerInfo.FullMethod, or nil if the method could not
// be found.
//
// SECURITY: If the given method does not have any
// metropolis.common.ext.authorization annotations, a methodInfo which requires
// authorization but no permissions is returned, defaulting to a mildly secure
// default of a method that can be called by any authenticated user.
func getMethodInfo(methodName string) (*methodInfo, error) {
	m := reMethodName.FindStringSubmatch(methodName)
	if len(m) != 3 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid method name %q", methodName)
	}
	// Convert /foo.bar/Method to foo.bar.Method, which is used by the protoregistry.
	methodName = fmt.Sprintf("%s.%s", m[1], m[2])
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(methodName))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not retrieve descriptor for method: %v", err)
	}
	method, ok := desc.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "querying method name did not yield a MethodDescriptor")
	}

	// Get authorization extension, defaults to no options set.
	if !proto.HasExtension(method.Options(), epb.E_Authorization) {
		return nil, status.Errorf(codes.Internal, "method does not provide Authorization extension, failing safe")
	}
	authz, ok := proto.GetExtension(method.Options(), epb.E_Authorization).(*epb.Authorization)
	if !ok {
		return nil, status.Errorf(codes.Internal, "method contains Authorization extension with wrong type, failing safe")
	}
	if authz == nil {
		return nil, status.Errorf(codes.Internal, "method contains nil Authorization extension, failing safe")
	}

	// If unauthenticated connections are allowed, return immediately.
	if authz.AllowUnauthenticated && len(authz.Need) == 0 {
		return &methodInfo{
			unauthenticated: true,
		}, nil
	}

	// Otherwise, return needed permissions.
	res := &methodInfo{
		need: make(map[epb.Permission]bool),
	}
	for _, n := range authz.Need {
		res.need[n] = true
	}
	return res, nil
}
