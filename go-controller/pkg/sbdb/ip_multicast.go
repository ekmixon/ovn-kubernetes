// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

// IPMulticast defines an object in IP_Multicast table
type IPMulticast struct {
	UUID          string `ovsdb:"_uuid"`
	Datapath      string `ovsdb:"datapath"`
	Enabled       *bool  `ovsdb:"enabled"`
	EthSrc        string `ovsdb:"eth_src"`
	IdleTimeout   *int   `ovsdb:"idle_timeout"`
	Ip4Src        string `ovsdb:"ip4_src"`
	Ip6Src        string `ovsdb:"ip6_src"`
	Querier       *bool  `ovsdb:"querier"`
	QueryInterval *int   `ovsdb:"query_interval"`
	QueryMaxResp  *int   `ovsdb:"query_max_resp"`
	SeqNo         int    `ovsdb:"seq_no"`
	TableSize     *int   `ovsdb:"table_size"`
}
