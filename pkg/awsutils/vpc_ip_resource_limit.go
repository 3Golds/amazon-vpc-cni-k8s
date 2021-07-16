// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2021-05-11T19:31:10Z

package awsutils

// InstanceNetworkingLimits contains a mapping from instance type to networking limits for the type. Documentation found at
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
var InstanceNetworkingLimits = map[string]InstanceTypeLimits{
	"a1.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"a1.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"a1.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"a1.medium":     {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"a1.metal":      {ENILimit: 8, IPv4Limit: 30, HypervisorType: ""},
	"a1.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c1.medium":     {ENILimit: 2, IPv4Limit: 6, HypervisorType: "xen"},
	"c1.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"c3.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"c3.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"c3.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"c3.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"c3.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"c4.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"c4.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"c4.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"c4.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"c4.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"c5.12xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5.18xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5.24xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5.9xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c5.metal":      {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"c5.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5a.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5a.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5a.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5a.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5a.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5a.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5a.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c5a.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: "unkown"},
	"c5a.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5ad.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5ad.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5ad.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5ad.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5ad.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5ad.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5ad.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c5ad.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: "unkown"},
	"c5ad.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5d.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5d.18xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5d.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5d.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5d.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5d.9xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5d.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c5d.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"c5d.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5n.18xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c5n.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c5n.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5n.9xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c5n.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c5n.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"c5n.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6g.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6g.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c6g.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6g.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6g.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6g.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c6g.medium":    {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"c6g.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"c6g.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6gd.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gd.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c6gd.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6gd.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gd.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gd.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c6gd.medium":   {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"c6gd.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"c6gd.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6gn.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gn.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"c6gn.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"c6gn.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gn.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"c6gn.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"c6gn.medium":   {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"c6gn.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"cc2.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"cr1.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "unkown"},
	"d2.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"d2.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"d2.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"d2.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"d3.2xlarge":    {ENILimit: 4, IPv4Limit: 5, HypervisorType: "nitro"},
	"d3.4xlarge":    {ENILimit: 4, IPv4Limit: 10, HypervisorType: "nitro"},
	"d3.8xlarge":    {ENILimit: 3, IPv4Limit: 20, HypervisorType: "nitro"},
	"d3.xlarge":     {ENILimit: 4, IPv4Limit: 3, HypervisorType: "nitro"},
	"d3en.12xlarge": {ENILimit: 3, IPv4Limit: 30, HypervisorType: "nitro"},
	"d3en.2xlarge":  {ENILimit: 4, IPv4Limit: 5, HypervisorType: "nitro"},
	"d3en.4xlarge":  {ENILimit: 4, IPv4Limit: 10, HypervisorType: "nitro"},
	"d3en.6xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"d3en.8xlarge":  {ENILimit: 4, IPv4Limit: 20, HypervisorType: "nitro"},
	"d3en.xlarge":   {ENILimit: 4, IPv4Limit: 3, HypervisorType: "nitro"},
	"f1.16xlarge":   {ENILimit: 8, IPv4Limit: 50, HypervisorType: "xen"},
	"f1.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"f1.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"g2.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"g2.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"g3.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "xen"},
	"g3.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"g3.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"g3s.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"g4ad.16xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"g4ad.4xlarge":  {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"g4ad.8xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"g4dn.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"g4dn.16xlarge": {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"g4dn.2xlarge":  {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"g4dn.4xlarge":  {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"g4dn.8xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"g4dn.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"g4dn.xlarge":   {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"h1.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "xen"},
	"h1.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"h1.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"h1.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"hs1.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "unkown"},
	"i2.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"i2.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"i2.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"i2.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"i3.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "xen"},
	"i3.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"i3.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"i3.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"i3.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"i3.metal":      {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"i3.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"i3en.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"i3en.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"i3en.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"i3en.3xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"i3en.6xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"i3en.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"i3en.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"i3en.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"inf1.24xlarge": {ENILimit: 11, IPv4Limit: 30, HypervisorType: "nitro"},
	"inf1.2xlarge":  {ENILimit: 4, IPv4Limit: 10, HypervisorType: "nitro"},
	"inf1.6xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"inf1.xlarge":   {ENILimit: 4, IPv4Limit: 10, HypervisorType: "nitro"},
	"m1.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"m1.medium":     {ENILimit: 2, IPv4Limit: 6, HypervisorType: "xen"},
	"m1.small":      {ENILimit: 2, IPv4Limit: 4, HypervisorType: "xen"},
	"m1.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"m2.2xlarge":    {ENILimit: 4, IPv4Limit: 30, HypervisorType: "xen"},
	"m2.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"m2.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"m3.2xlarge":    {ENILimit: 4, IPv4Limit: 30, HypervisorType: "xen"},
	"m3.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"m3.medium":     {ENILimit: 2, IPv4Limit: 6, HypervisorType: "xen"},
	"m3.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"m4.10xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"m4.16xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"m4.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"m4.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"m4.large":      {ENILimit: 2, IPv4Limit: 10, HypervisorType: "xen"},
	"m4.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"m5.12xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5.24xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5.metal":      {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m5.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5a.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5a.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5a.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5a.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5a.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5a.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5a.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5a.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5ad.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5ad.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5ad.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5ad.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5ad.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5ad.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5ad.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5ad.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5d.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5d.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5d.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5d.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5d.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5d.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5d.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5d.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m5d.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5dn.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5dn.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5dn.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5dn.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5dn.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5dn.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5dn.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5dn.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m5dn.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5n.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5n.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5n.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5n.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5n.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5n.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5n.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5n.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m5n.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5zn.12xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m5zn.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m5zn.3xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5zn.6xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m5zn.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m5zn.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m5zn.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m6g.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6g.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m6g.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m6g.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6g.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6g.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m6g.medium":    {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"m6g.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m6g.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m6gd.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6gd.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"m6gd.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"m6gd.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6gd.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"m6gd.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"m6gd.medium":   {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"m6gd.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"m6gd.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"mac1.metal":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: ""},
	"p2.16xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"p2.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"p2.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"p3.16xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"p3.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"p3.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"p3dn.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"p4d.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "unkown"},
	"r3.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"r3.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"r3.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"r3.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"r3.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"r4.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "xen"},
	"r4.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"r4.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"r4.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"r4.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"r4.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"r5.12xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5.16xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5.24xlarge":   {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5.4xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5.8xlarge":    {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5.large":      {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5.metal":      {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r5.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5a.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5a.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5a.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5a.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5a.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5a.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5a.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5a.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5ad.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5ad.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5ad.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5ad.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5ad.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5ad.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5ad.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5ad.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5b.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5b.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5b.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5b.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5b.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5b.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5b.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5b.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r5b.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5d.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5d.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5d.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5d.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5d.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5d.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5d.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5d.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r5d.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5dn.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5dn.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5dn.24xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5dn.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5dn.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5dn.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5dn.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5dn.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r5dn.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5n.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5n.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5n.24xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r5n.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r5n.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5n.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r5n.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r5n.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r5n.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r6g.12xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6g.16xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r6g.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r6g.4xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6g.8xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6g.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r6g.medium":    {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"r6g.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r6g.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r6gd.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6gd.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"r6gd.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"r6gd.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6gd.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"r6gd.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"r6gd.medium":   {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"r6gd.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"r6gd.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t1.micro":      {ENILimit: 2, IPv4Limit: 2, HypervisorType: "xen"},
	"t2.2xlarge":    {ENILimit: 3, IPv4Limit: 15, HypervisorType: "xen"},
	"t2.large":      {ENILimit: 3, IPv4Limit: 12, HypervisorType: "xen"},
	"t2.medium":     {ENILimit: 3, IPv4Limit: 6, HypervisorType: "xen"},
	"t2.micro":      {ENILimit: 2, IPv4Limit: 2, HypervisorType: "xen"},
	"t2.nano":       {ENILimit: 2, IPv4Limit: 2, HypervisorType: "xen"},
	"t2.small":      {ENILimit: 3, IPv4Limit: 4, HypervisorType: "xen"},
	"t2.xlarge":     {ENILimit: 3, IPv4Limit: 15, HypervisorType: "xen"},
	"t3.2xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t3.large":      {ENILimit: 3, IPv4Limit: 12, HypervisorType: "nitro"},
	"t3.medium":     {ENILimit: 3, IPv4Limit: 6, HypervisorType: "nitro"},
	"t3.micro":      {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t3.nano":       {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t3.small":      {ENILimit: 3, IPv4Limit: 4, HypervisorType: "nitro"},
	"t3.xlarge":     {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t3a.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t3a.large":     {ENILimit: 3, IPv4Limit: 12, HypervisorType: "nitro"},
	"t3a.medium":    {ENILimit: 3, IPv4Limit: 6, HypervisorType: "nitro"},
	"t3a.micro":     {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t3a.nano":      {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t3a.small":     {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"t3a.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t4g.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"t4g.large":     {ENILimit: 3, IPv4Limit: 12, HypervisorType: "nitro"},
	"t4g.medium":    {ENILimit: 3, IPv4Limit: 6, HypervisorType: "nitro"},
	"t4g.micro":     {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t4g.nano":      {ENILimit: 2, IPv4Limit: 2, HypervisorType: "nitro"},
	"t4g.small":     {ENILimit: 3, IPv4Limit: 4, HypervisorType: "nitro"},
	"t4g.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"u-12tb1.metal": {ENILimit: 5, IPv4Limit: 30, HypervisorType: "unkown"},
	"u-18tb1.metal": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "unkown"},
	"u-24tb1.metal": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "unkown"},
	"u-6tb1.metal":  {ENILimit: 5, IPv4Limit: 30, HypervisorType: "unkown"},
	"u-9tb1.metal":  {ENILimit: 5, IPv4Limit: 30, HypervisorType: "unkown"},
	"x1.16xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"x1.32xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"x1e.16xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"x1e.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"x1e.32xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "xen"},
	"x1e.4xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"x1e.8xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "xen"},
	"x1e.xlarge":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "xen"},
	"x2gd.12xlarge": {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"x2gd.16xlarge": {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"x2gd.2xlarge":  {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"x2gd.4xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"x2gd.8xlarge":  {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"x2gd.large":    {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"x2gd.medium":   {ENILimit: 2, IPv4Limit: 4, HypervisorType: "nitro"},
	"x2gd.metal":    {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"x2gd.xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"z1d.12xlarge":  {ENILimit: 15, IPv4Limit: 50, HypervisorType: "nitro"},
	"z1d.2xlarge":   {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
	"z1d.3xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"z1d.6xlarge":   {ENILimit: 8, IPv4Limit: 30, HypervisorType: "nitro"},
	"z1d.large":     {ENILimit: 3, IPv4Limit: 10, HypervisorType: "nitro"},
	"z1d.metal":     {ENILimit: 15, IPv4Limit: 50, HypervisorType: ""},
	"z1d.xlarge":    {ENILimit: 4, IPv4Limit: 15, HypervisorType: "nitro"},
}
