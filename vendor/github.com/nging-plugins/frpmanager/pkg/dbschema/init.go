// @generated Do not edit a file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

var WithPrefix = func(tableName string) string {
	return "" + tableName
}

var DBI = factory.DefaultDBI

func init() {

	DBI.Fields.Register(map[string]map[string]*factory.FieldInfo{"nging_frp_client": {"admin_addr": {Name: "admin_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "AdminAddr"}, "admin_port": {Name: "admin_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "AdminPort"}, "admin_pwd": {Name: "admin_pwd", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "AdminPwd"}, "admin_user": {Name: "admin_user", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "AdminUser"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Created"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁用", GoType: "string", MyType: "", GoName: "Disabled"}, "dns_server": {Name: "dns_server", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "DnsServer"}, "extra": {Name: "extra", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "Extra"}, "group_id": {Name: "group_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "GroupId"}, "heartbeat_interval": {Name: "heartbeat_interval", DataType: "bigint", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: -0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "30", Comment: "", GoType: "int64", MyType: "", GoName: "HeartbeatInterval"}, "heartbeat_timeout": {Name: "heartbeat_timeout", DataType: "bigint", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: -0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "90", Comment: "", GoType: "int64", MyType: "", GoName: "HeartbeatTimeout"}, "http_proxy": {Name: "http_proxy", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "HttpProxy"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "log_file": {Name: "log_file", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "console", Comment: "", GoType: "string", MyType: "", GoName: "LogFile"}, "log_level": {Name: "log_level", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "info", Comment: "", GoType: "string", MyType: "", GoName: "LogLevel"}, "log_max_days": {Name: "log_max_days", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "3", Comment: "", GoType: "uint", MyType: "", GoName: "LogMaxDays"}, "log_way": {Name: "log_way", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "console", Comment: "console or file", GoType: "string", MyType: "", GoName: "LogWay"}, "login_fail_exit": {Name: "login_fail_exit", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "Y", Comment: "", GoType: "string", MyType: "", GoName: "LoginFailExit"}, "metas": {Name: "metas", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 2000, Options: []string{}, DefaultValue: "", Comment: "meta值", GoType: "string", MyType: "", GoName: "Metas"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 60, Options: []string{}, DefaultValue: "", Comment: "名称", GoType: "string", MyType: "", GoName: "Name"}, "pool_count": {Name: "pool_count", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "1", Comment: "", GoType: "uint", MyType: "", GoName: "PoolCount"}, "protocol": {Name: "protocol", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "tcp", Comment: "tcp or kcp", GoType: "string", MyType: "", GoName: "Protocol"}, "server_addr": {Name: "server_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "0.0.0.0", Comment: "", GoType: "string", MyType: "", GoName: "ServerAddr"}, "server_port": {Name: "server_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "7000", Comment: "", GoType: "uint", MyType: "", GoName: "ServerPort"}, "start": {Name: "start", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 2000, Options: []string{}, DefaultValue: "", Comment: "要启动的代理节点名称，留空代表全部，多个用半角逗号隔开", GoType: "string", MyType: "", GoName: "Start"}, "tcp_mux": {Name: "tcp_mux", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "Y", Comment: "必须跟服务端一致", GoType: "string", MyType: "", GoName: "TcpMux"}, "token": {Name: "token", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "Token"}, "type": {Name: "type", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "web", Comment: "", GoType: "string", MyType: "", GoName: "Type"}, "uid": {Name: "uid", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Uid"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Updated"}, "user": {Name: "user", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "User"}}, "nging_frp_group": {"created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "description": {Name: "description", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "说明", GoType: "string", MyType: "", GoName: "Description"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "uint", MyType: "", GoName: "Id"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 60, Options: []string{}, DefaultValue: "", Comment: "组名", GoType: "string", MyType: "", GoName: "Name"}, "uid": {Name: "uid", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "用户ID", GoType: "uint", MyType: "", GoName: "Uid"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "修改时间", GoType: "uint", MyType: "", GoName: "Updated"}}, "nging_frp_server": {"addr": {Name: "addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "0.0.0.0", Comment: "", GoType: "string", MyType: "", GoName: "Addr"}, "allow_ports": {Name: "allow_ports", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 2000, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "AllowPorts"}, "auth_timeout": {Name: "auth_timeout", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "900", Comment: "", GoType: "uint64", MyType: "", GoName: "AuthTimeout"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Created"}, "dashboard_addr": {Name: "dashboard_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "0.0.0.0", Comment: "", GoType: "string", MyType: "", GoName: "DashboardAddr"}, "dashboard_port": {Name: "dashboard_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "DashboardPort"}, "dashboard_pwd": {Name: "dashboard_pwd", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "admin", Comment: "", GoType: "string", MyType: "", GoName: "DashboardPwd"}, "dashboard_user": {Name: "dashboard_user", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "admin", Comment: "", GoType: "string", MyType: "", GoName: "DashboardUser"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁用", GoType: "string", MyType: "", GoName: "Disabled"}, "extra": {Name: "extra", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "Extra"}, "group_id": {Name: "group_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "GroupId"}, "heart_beat_timeout": {Name: "heart_beat_timeout", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "90", Comment: "", GoType: "uint", MyType: "", GoName: "HeartBeatTimeout"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "kcp_port": {Name: "kcp_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "KcpPort"}, "log_file": {Name: "log_file", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "console", Comment: "", GoType: "string", MyType: "", GoName: "LogFile"}, "log_level": {Name: "log_level", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "info", Comment: "", GoType: "string", MyType: "", GoName: "LogLevel"}, "log_max_days": {Name: "log_max_days", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "3", Comment: "", GoType: "uint", MyType: "", GoName: "LogMaxDays"}, "log_way": {Name: "log_way", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "console", Comment: "console or file", GoType: "string", MyType: "", GoName: "LogWay"}, "max_pool_count": {Name: "max_pool_count", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "5", Comment: "", GoType: "uint", MyType: "", GoName: "MaxPoolCount"}, "max_ports_per_client": {Name: "max_ports_per_client", DataType: "bigint", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: -0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "int64", MyType: "", GoName: "MaxPortsPerClient"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 60, Options: []string{}, DefaultValue: "", Comment: "名称", GoType: "string", MyType: "", GoName: "Name"}, "plugins": {Name: "plugins", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 2000, Options: []string{}, DefaultValue: "", Comment: "启用插件(半角逗号分隔)", GoType: "string", MyType: "", GoName: "Plugins"}, "port": {Name: "port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "7000", Comment: "", GoType: "uint", MyType: "", GoName: "Port"}, "proxy_addr": {Name: "proxy_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "0.0.0.0", Comment: "", GoType: "string", MyType: "", GoName: "ProxyAddr"}, "subdomain_host": {Name: "subdomain_host", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "SubdomainHost"}, "tcp_mux": {Name: "tcp_mux", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "Y", Comment: "", GoType: "string", MyType: "", GoName: "TcpMux"}, "token": {Name: "token", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 128, Options: []string{}, DefaultValue: "", Comment: "", GoType: "string", MyType: "", GoName: "Token"}, "udp_port": {Name: "udp_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "UdpPort"}, "uid": {Name: "uid", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Uid"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "Updated"}, "user_conn_timeout": {Name: "user_conn_timeout", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "10", Comment: "", GoType: "uint", MyType: "", GoName: "UserConnTimeout"}, "vhost_http_port": {Name: "vhost_http_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "VhostHttpPort"}, "vhost_http_timeout": {Name: "vhost_http_timeout", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint64", MyType: "", GoName: "VhostHttpTimeout"}, "vhost_https_port": {Name: "vhost_https_port", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "", GoType: "uint", MyType: "", GoName: "VhostHttpsPort"}}, "nging_frp_user": {"bandwidth": {Name: "bandwidth", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 500, Options: []string{}, DefaultValue: "", Comment: "带宽(MB/KB) ", GoType: "string", MyType: "", GoName: "Bandwidth"}, "banned": {Name: "banned", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁止连接", GoType: "string", MyType: "", GoName: "Banned"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "customer_id": {Name: "customer_id", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "客户ID", GoType: "uint64", MyType: "", GoName: "CustomerId"}, "end": {Name: "end", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "过期时间", GoType: "uint", MyType: "", GoName: "End"}, "id": {Name: "id", DataType: "bigint", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "uint64", MyType: "", GoName: "Id"}, "ip_blacklist": {Name: "ip_blacklist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP黑名单(一行一个) ", GoType: "string", MyType: "", GoName: "IpBlacklist"}, "ip_whitelist": {Name: "ip_whitelist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP白名单(一行一个) ", GoType: "string", MyType: "", GoName: "IpWhitelist"}, "password": {Name: "password", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 150, Options: []string{}, DefaultValue: "", Comment: "密码", GoType: "string", MyType: "", GoName: "Password"}, "server_id": {Name: "server_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "服务ID", GoType: "uint", MyType: "", GoName: "ServerId"}, "start": {Name: "start", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "生效时间", GoType: "uint", MyType: "", GoName: "Start"}, "traffic_total": {Name: "traffic_total", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "流量总量", GoType: "uint64", MyType: "", GoName: "TrafficTotal"}, "traffic_used": {Name: "traffic_used", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "已用流量", GoType: "uint64", MyType: "", GoName: "TrafficUsed"}, "uid": {Name: "uid", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "管理员ID", GoType: "uint", MyType: "", GoName: "Uid"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "修改时间", GoType: "uint", MyType: "", GoName: "Updated"}, "username": {Name: "username", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "", Comment: "用户名", GoType: "string", MyType: "", GoName: "Username"}}})

	DBI.Columns = map[string][]string{"nging_frp_client": {"id", "name", "disabled", "server_addr", "server_port", "http_proxy", "pool_count", "tcp_mux", "user", "dns_server", "login_fail_exit", "protocol", "heartbeat_interval", "heartbeat_timeout", "log_file", "log_way", "log_level", "log_max_days", "token", "admin_addr", "admin_port", "admin_user", "admin_pwd", "start", "metas", "extra", "uid", "group_id", "type", "created", "updated"}, "nging_frp_group": {"id", "uid", "name", "description", "created", "updated"}, "nging_frp_server": {"id", "name", "disabled", "tcp_mux", "addr", "port", "udp_port", "kcp_port", "proxy_addr", "vhost_http_port", "vhost_http_timeout", "vhost_https_port", "log_file", "log_way", "log_level", "log_max_days", "token", "auth_timeout", "subdomain_host", "max_ports_per_client", "max_pool_count", "heart_beat_timeout", "user_conn_timeout", "dashboard_addr", "dashboard_port", "dashboard_user", "dashboard_pwd", "allow_ports", "extra", "plugins", "uid", "group_id", "created", "updated"}, "nging_frp_user": {"id", "server_id", "customer_id", "uid", "username", "password", "banned", "bandwidth", "traffic_used", "traffic_total", "ip_whitelist", "ip_blacklist", "start", "end", "created", "updated"}}

	DBI.Models.Register(factory.ModelInstancers{`NgingFrpClient`: factory.NewMI("nging_frp_client", func(connID int) factory.Model { return &NgingFrpClient{base: *((&factory.Base{}).SetConnID(connID))} }, "FRP客户端设置"), `NgingFrpGroup`: factory.NewMI("nging_frp_group", func(connID int) factory.Model { return &NgingFrpGroup{base: *((&factory.Base{}).SetConnID(connID))} }, "FRP服务组"), `NgingFrpServer`: factory.NewMI("nging_frp_server", func(connID int) factory.Model { return &NgingFrpServer{base: *((&factory.Base{}).SetConnID(connID))} }, "FRP服务器设置"), `NgingFrpUser`: factory.NewMI("nging_frp_user", func(connID int) factory.Model { return &NgingFrpUser{base: *((&factory.Base{}).SetConnID(connID))} }, "网络用户")})

}