// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

type (
	LogicalFlowPipeline = string
)

var (
	LogicalFlowPipelineEgress  LogicalFlowPipeline = "egress"
	LogicalFlowPipelineIngress LogicalFlowPipeline = "ingress"
)

// LogicalFlow defines an object in Logical_Flow table
type LogicalFlow struct {
	UUID            string              `ovsdb:"_uuid"`
	Actions         string              `ovsdb:"actions"`
	ExternalIDs     map[string]string   `ovsdb:"external_ids"`
	LogicalDatapath *string             `ovsdb:"logical_datapath"`
	LogicalDpGroup  *string             `ovsdb:"logical_dp_group"`
	Match           string              `ovsdb:"match"`
	Pipeline        LogicalFlowPipeline `ovsdb:"pipeline"`
	Priority        int                 `ovsdb:"priority"`
	TableID         int                 `ovsdb:"table_id"`
}
