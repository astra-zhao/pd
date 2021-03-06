// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	regionsPrefix          = "pd/api/v1/regions"
	regionsStorePrefix     = "pd/api/v1/regions/store"
	regionsCheckPrefix     = "pd/api/v1/regions/check"
	regionsWriteflowPrefix = "pd/api/v1/regions/writeflow"
	regionsReadflowPrefix  = "pd/api/v1/regions/readflow"
	regionsConfVerPrefix   = "pd/api/v1/regions/confver"
	regionsVersionPrefix   = "pd/api/v1/regions/version"
	regionsSizePrefix      = "pd/api/v1/regions/size"
	regionsSiblingPrefix   = "pd/api/v1/regions/sibling"
	regionIDPrefix         = "pd/api/v1/region/id"
	regionKeyPrefix        = "pd/api/v1/region/key"
)

// NewRegionCommand returns a region subcommand of rootCmd
func NewRegionCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   `region <region_id> [-jq="<query string>"]`,
		Short: "show the region status",
		Run:   showRegionCommandFunc,
	}
	r.AddCommand(NewRegionWithKeyCommand())
	r.AddCommand(NewRegionWithCheckCommand())
	r.AddCommand(NewRegionWithSiblingCommand())
	r.AddCommand(NewRegionWithStoreCommand())
	r.AddCommand(NewRegionsWithStartKeyCommand())

	topRead := &cobra.Command{
		Use:   "topread <limit>",
		Short: "show regions with top read flow",
		Run:   showRegionTopReadCommandFunc,
	}
	r.AddCommand(topRead)

	topWrite := &cobra.Command{
		Use:   "topwrite <limit>",
		Short: "show regions with top write flow",
		Run:   showRegionTopWriteCommandFunc,
	}
	r.AddCommand(topWrite)

	topConfVer := &cobra.Command{
		Use:   "topconfver <limit>",
		Short: "show regions with top conf version",
		Run:   showRegionTopConfVerCommandFunc,
	}
	r.AddCommand(topConfVer)

	topVersion := &cobra.Command{
		Use:   "topversion <limit>",
		Short: "show regions with top version",
		Run:   showRegionTopVersionCommandFunc,
	}
	r.AddCommand(topVersion)

	topSize := &cobra.Command{
		Use:   "topsize <limit>",
		Short: "show regions with top size",
		Run:   showRegionTopSizeCommandFunc,
	}
	r.AddCommand(topSize)
	r.Flags().String("jq", "", "jq query")

	return r
}

func showRegionCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsPrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("region_id should be a number")
			return
		}
		prefix = regionIDPrefix + "/" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get region: %s\n", err)
		return
	}
	if flag := cmd.Flag("jq"); flag != nil && flag.Value.String() != "" {
		printWithJQFilter(r, flag.Value.String())
		return
	}

	fmt.Println(r)
}

func showRegionTopWriteCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsWriteflowPrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions: %s\n", err)
		return
	}
	fmt.Println(r)
}

func showRegionTopReadCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsReadflowPrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions: %s\n", err)
		return
	}
	fmt.Println(r)
}

func showRegionTopConfVerCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsConfVerPrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions: %s\n", err)
		return
	}
	fmt.Println(r)
}

func showRegionTopVersionCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsVersionPrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions: %s\n", err)
		return
	}
	fmt.Println(r)
}

func showRegionTopSizeCommandFunc(cmd *cobra.Command, args []string) {
	prefix := regionsSizePrefix
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[0]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions: %s\n", err)
		return
	}
	fmt.Println(r)
}

// NewRegionWithKeyCommand return a region with key subcommand of regionCmd
func NewRegionWithKeyCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   "key [--format=raw|encode] <key>",
		Short: "show the region with key",
		Run:   showRegionWithTableCommandFunc,
	}
	r.Flags().String("format", "raw", "the key format")
	return r
}

func showRegionWithTableCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.UsageString())
		return
	}

	var (
		key string
		err error
	)

	format := cmd.Flags().Lookup("format").Value.String()
	switch format {
	case "raw":
		key = args[0]
	case "encode":
		key, err = decodeKey(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	default:
		fmt.Println("Error: unknown format")
		return
	}

	key = url.QueryEscape(key)
	prefix := regionKeyPrefix + "/" + key
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get region: %s\n", err)
		return
	}
	fmt.Println(r)
}

func decodeKey(text string) (string, error) {
	var buf []byte
	r := bytes.NewBuffer([]byte(text))
	for {
		c, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		if c != '\\' {
			buf = append(buf, c)
			continue
		}
		n := r.Next(1)
		switch n[0] {
		case '"':
			buf = append(buf, '"')
		case '\'':
			buf = append(buf, '\'')
		case '\\':
			buf = append(buf, '\\')
		case 'n':
			buf = append(buf, '\n')
		case 't':
			buf = append(buf, '\t')
		case 'r':
			buf = append(buf, '\r')
		case 'x':
			fmt.Sscanf(string(r.Next(2)), "%02x", &c)
			buf = append(buf, c)
		default:
			n = append(n, r.Next(2)...)
			_, err := fmt.Sscanf(string(n), "%03o", &c)
			if err != nil {
				return "", err
			}
			buf = append(buf, c)
		}
	}
	return string(buf), nil
}

// NewRegionsWithStartKeyCommand returns regions from startkey subcommand of regionCmd.
func NewRegionsWithStartKeyCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   "startkey [--format=raw|encode] <key> <limit>",
		Short: "show regions from start key",
		Run:   showRegionsFromStartKeyCommandFunc,
	}

	r.Flags().String("format", "raw", "the key format")
	return r
}

func showRegionsFromStartKeyCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 2 {
		fmt.Println(cmd.UsageString())
		return
	}

	var (
		key string
		err error
	)

	format := cmd.Flags().Lookup("format").Value.String()
	switch format {
	case "raw":
		key = args[0]
	case "encode":
		key, err = decodeKey(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	default:
		fmt.Println("Error: unknown format")
		return
	}
	key = url.QueryEscape(key)
	prefix := regionKeyPrefix + "/" + key
	if len(args) == 2 {
		if _, err = strconv.Atoi(args[1]); err != nil {
			fmt.Println("limit should be a number")
			return
		}
		prefix += "?limit=" + args[1]
	}
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get region: %s\n", err)
		return
	}
	fmt.Println(r)
}

// NewRegionWithCheckCommand returns a region with check subcommand of regionCmd
func NewRegionWithCheckCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   "check [miss-peer|extra-peer|down-peer|pending-peer|incorrect-ns]",
		Short: "show the region with check specific status",
		Run:   showRegionWithCheckCommandFunc,
	}
	return r
}

func showRegionWithCheckCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.UsageString())
		return
	}
	state := args[0]
	prefix := regionsCheckPrefix + "/" + state
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get region: %s\n", err)
		return
	}
	fmt.Println(r)
}

// NewRegionWithSiblingCommand returns a region with sibling subcommand of regionCmd
func NewRegionWithSiblingCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   "sibling <region_id>",
		Short: "show the sibling regions of specific region",
		Run:   showRegionWithSiblingCommandFunc,
	}
	return r
}

func showRegionWithSiblingCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.UsageString())
		return
	}
	regionID := args[0]
	prefix := regionsSiblingPrefix + "/" + regionID
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get region sibling: %s\n", err)
		return
	}
	fmt.Println(r)
}

// NewRegionWithStoreCommand returns regions with store subcommand of regionCmd
func NewRegionWithStoreCommand() *cobra.Command {
	r := &cobra.Command{
		Use:   "store <store_id>",
		Short: "show the regions of a specific store",
		Run:   showRegionWithStoreCommandFunc,
	}
	return r
}

func showRegionWithStoreCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.UsageString())
		return
	}
	storeID := args[0]
	prefix := regionsStorePrefix + "/" + storeID
	r, err := doRequest(cmd, prefix, http.MethodGet)
	if err != nil {
		fmt.Printf("Failed to get regions with the given storeID: %s\n", err)
		return
	}
	fmt.Println(r)
}

func printWithJQFilter(data, filter string) {
	cmd := exec.Command("jq", "-c", filter)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, data)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out), err)
		return
	}

	fmt.Printf("%s\n", out)
}
