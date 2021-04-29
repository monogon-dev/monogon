package client

import (
	"context"
	"errors"
	"io"

	"go.etcd.io/etcd/clientv3"
)

var (
	// UnimplementedInNamespaced will be raised by panic() any time a method
	// from the Cluster, Auth and Maintenance APIs gets called on a
	// clientv3.Client returned by ThinClient or Namespaced.ThinClient.
	UnimplementedInNamespaced = errors.New("interface not implemented in Namespaced etcd client")
)

// unimplementedCluster implements clientv3.Cluster, but panics on any call.
type unimplementedCluster struct {
}

func (c *unimplementedCluster) MemberList(_ context.Context) (*clientv3.MemberListResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedCluster) MemberAdd(_ context.Context, _ []string) (*clientv3.MemberAddResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedCluster) MemberAddAsLearner(_ context.Context, _ []string) (*clientv3.MemberAddResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedCluster) MemberRemove(_ context.Context, _ uint64) (*clientv3.MemberRemoveResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedCluster) MemberUpdate(_ context.Context, _ uint64, _ []string) (*clientv3.MemberUpdateResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedCluster) MemberPromote(_ context.Context, _ uint64) (*clientv3.MemberPromoteResponse, error) {
	panic(UnimplementedInNamespaced)
}

// unimplementedAuth implements clientv3.Auth but panics on any call.
type unimplementedAuth struct {
}

func (c *unimplementedAuth) Authenticate(ctx context.Context, name string, password string) (*clientv3.AuthenticateResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) AuthEnable(ctx context.Context) (*clientv3.AuthEnableResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) AuthDisable(ctx context.Context) (*clientv3.AuthDisableResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserAdd(ctx context.Context, name string, password string) (*clientv3.AuthUserAddResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserAddWithOptions(ctx context.Context, name string, password string, opt *clientv3.UserAddOptions) (*clientv3.AuthUserAddResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserDelete(ctx context.Context, name string) (*clientv3.AuthUserDeleteResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserChangePassword(ctx context.Context, name string, password string) (*clientv3.AuthUserChangePasswordResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserGrantRole(ctx context.Context, user string, role string) (*clientv3.AuthUserGrantRoleResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserGet(ctx context.Context, name string) (*clientv3.AuthUserGetResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserList(ctx context.Context) (*clientv3.AuthUserListResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) UserRevokeRole(ctx context.Context, name string, role string) (*clientv3.AuthUserRevokeRoleResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleAdd(ctx context.Context, name string) (*clientv3.AuthRoleAddResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleGrantPermission(ctx context.Context, name string, key, rangeEnd string, permType clientv3.PermissionType) (*clientv3.AuthRoleGrantPermissionResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleGet(ctx context.Context, role string) (*clientv3.AuthRoleGetResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleList(ctx context.Context) (*clientv3.AuthRoleListResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleRevokePermission(ctx context.Context, role string, key, rangeEnd string) (*clientv3.AuthRoleRevokePermissionResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedAuth) RoleDelete(ctx context.Context, role string) (*clientv3.AuthRoleDeleteResponse, error) {
	panic(UnimplementedInNamespaced)
}

// unimplementedMaintenance implements clientv3.Maintenance but panics on any call.
type unimplementedMaintenance struct {
}

func (c *unimplementedMaintenance) AlarmList(ctx context.Context) (*clientv3.AlarmResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) AlarmDisarm(ctx context.Context, m *clientv3.AlarmMember) (*clientv3.AlarmResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) Defragment(ctx context.Context, endpoint string) (*clientv3.DefragmentResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) Status(ctx context.Context, endpoint string) (*clientv3.StatusResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) HashKV(ctx context.Context, endpoint string, rev int64) (*clientv3.HashKVResponse, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) Snapshot(ctx context.Context) (io.ReadCloser, error) {
	panic(UnimplementedInNamespaced)
}

func (c *unimplementedMaintenance) MoveLeader(ctx context.Context, transfereeID uint64) (*clientv3.MoveLeaderResponse, error) {
	panic(UnimplementedInNamespaced)
}
