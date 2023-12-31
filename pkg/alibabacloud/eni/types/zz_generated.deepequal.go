//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package types

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ENI) DeepEqual(other *ENI) bool {
	if other == nil {
		return false
	}

	if in.NetworkInterfaceID != other.NetworkInterfaceID {
		return false
	}
	if in.MACAddress != other.MACAddress {
		return false
	}
	if in.Type != other.Type {
		return false
	}
	if in.InstanceID != other.InstanceID {
		return false
	}
	if ((in.SecurityGroupIDs != nil) && (other.SecurityGroupIDs != nil)) || ((in.SecurityGroupIDs == nil) != (other.SecurityGroupIDs == nil)) {
		in, other := &in.SecurityGroupIDs, &other.SecurityGroupIDs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if !in.VPC.DeepEqual(&other.VPC) {
		return false
	}

	if in.ZoneID != other.ZoneID {
		return false
	}
	if in.VSwitch != other.VSwitch {
		return false
	}

	if in.PrimaryIPAddress != other.PrimaryIPAddress {
		return false
	}
	if ((in.PrivateIPSets != nil) && (other.PrivateIPSets != nil)) || ((in.PrivateIPSets == nil) != (other.PrivateIPSets == nil)) {
		in, other := &in.PrivateIPSets, &other.PrivateIPSets
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Tags != nil) && (other.Tags != nil)) || ((in.Tags == nil) != (other.Tags == nil)) {
		in, other := &in.Tags, &other.Tags
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ENIStatus) DeepEqual(other *ENIStatus) bool {
	if other == nil {
		return false
	}

	if ((in.ENIs != nil) && (other.ENIs != nil)) || ((in.ENIs == nil) != (other.ENIs == nil)) {
		in, other := &in.ENIs, &other.ENIs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PrivateIPSet) DeepEqual(other *PrivateIPSet) bool {
	if other == nil {
		return false
	}

	if in.PrivateIpAddress != other.PrivateIpAddress {
		return false
	}
	if in.Primary != other.Primary {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Spec) DeepEqual(other *Spec) bool {
	if other == nil {
		return false
	}

	if in.InstanceType != other.InstanceType {
		return false
	}
	if in.AvailabilityZone != other.AvailabilityZone {
		return false
	}
	if in.VPCID != other.VPCID {
		return false
	}
	if in.CIDRBlock != other.CIDRBlock {
		return false
	}
	if ((in.VSwitches != nil) && (other.VSwitches != nil)) || ((in.VSwitches == nil) != (other.VSwitches == nil)) {
		in, other := &in.VSwitches, &other.VSwitches
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.VSwitchTags != nil) && (other.VSwitchTags != nil)) || ((in.VSwitchTags == nil) != (other.VSwitchTags == nil)) {
		in, other := &in.VSwitchTags, &other.VSwitchTags
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.SecurityGroups != nil) && (other.SecurityGroups != nil)) || ((in.SecurityGroups == nil) != (other.SecurityGroups == nil)) {
		in, other := &in.SecurityGroups, &other.SecurityGroups
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SecurityGroupTags != nil) && (other.SecurityGroupTags != nil)) || ((in.SecurityGroupTags == nil) != (other.SecurityGroupTags == nil)) {
		in, other := &in.SecurityGroupTags, &other.SecurityGroupTags
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *VPC) DeepEqual(other *VPC) bool {
	if other == nil {
		return false
	}

	if in.VPCID != other.VPCID {
		return false
	}
	if in.CIDRBlock != other.CIDRBlock {
		return false
	}
	if in.IPv6CIDRBlock != other.IPv6CIDRBlock {
		return false
	}
	if ((in.SecondaryCIDRs != nil) && (other.SecondaryCIDRs != nil)) || ((in.SecondaryCIDRs == nil) != (other.SecondaryCIDRs == nil)) {
		in, other := &in.SecondaryCIDRs, &other.SecondaryCIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *VSwitch) DeepEqual(other *VSwitch) bool {
	if other == nil {
		return false
	}

	if in.VSwitchID != other.VSwitchID {
		return false
	}
	if in.CIDRBlock != other.CIDRBlock {
		return false
	}
	if in.IPv6CIDRBlock != other.IPv6CIDRBlock {
		return false
	}

	return true
}
