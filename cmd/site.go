package cmd

/*
Copyright © 2020 Peter Howe <pnhowe@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"errors"
	"fmt"
	"os"

	cinp "github.com/cinp/go"
	"github.com/spf13/cobra"
	//"github.com/t3kton/contractorcli/contractor"
)

// var configSetName, configSetValue, configDeleteName string
// var configFull bool
var detailName, detailDescription, detailParent, detailZone string

func siteArgCheck(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Requires a Site Id/Name Argument")
	}
	return nil
}

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Work with sites",
}

var siteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Sites",
	Run: func(cmd *cobra.Command, args []string) {
		c := getContractor()
		defer c.Logout()

		rl := []cinp.Object{}
		for v := range c.SiteSiteList("", map[string]interface{}{}) {
			rl = append(rl, v)
		}
		outputList(rl, "Id	Name	Description	Created	Updated\n", "{{.GetID | extractID}}	{{.Name}}	{{.Description}}	{{.Created}}	{{.Updated}}\n")
	},
}

var siteGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Site",
	Args:  siteArgCheck,
	Run: func(cmd *cobra.Command, args []string) {
		siteID := args[0]
		c := getContractor()
		defer c.Logout()

		r, err := c.SiteSiteGet(siteID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputDetail(r, `Site:          {{.Name}}
Description:   {{.Description}}
Parent:        {{.Parent | extractID}}
Zone:          {{.Zone | extractID}}
Config Values: {{.ConfigValues}}
Created:       {{.Created}}
Updated:       {{.Updated}}
`)
	},
}

var siteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create New Site",
	Run: func(cmd *cobra.Command, args []string) {
		c := getContractor()
		defer c.Logout()

		o := c.SiteSiteNew()
		o.Name = detailName
		o.Description = detailDescription

		if detailParent != "" {
			r, err := c.SiteSiteGet(detailParent)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o.Parent = r.GetID()
		}

		if detailZone != "" {
			r, err := c.DirectoryZoneGet(detailZone)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o.Zone = r.GetID()
		}

		if err := o.Create(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var siteUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Site",
	Args:  siteArgCheck,
	Run: func(cmd *cobra.Command, args []string) {
		fieldList := []string{}
		siteID := args[0]
		c := getContractor()
		defer c.Logout()

		o, err := c.SiteSiteGet(siteID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if detailDescription != "" {
			o.Description = detailDescription
			fieldList = append(fieldList, "description")
		}

		if detailParent != "" {
			r, err := c.SiteSiteGet(detailParent)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o.Parent = r.GetID()
			fieldList = append(fieldList, "parent")
		}

		if detailZone != "" {
			r, err := c.DirectoryZoneGet(detailZone)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o.Zone = r.GetID()
			fieldList = append(fieldList, "zone")
		}

		if err := o.Update(fieldList); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var siteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Site",
	Args:  siteArgCheck,
	Run: func(cmd *cobra.Command, args []string) {
		siteID := args[0]
		c := getContractor()
		defer c.Logout()

		r, err := c.SiteSiteGet(siteID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := r.Delete(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var siteConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Work With Site Config",
	Args:  siteArgCheck,
	Run: func(cmd *cobra.Command, args []string) {
		siteID := args[0]
		c := getContractor()
		defer c.Logout()

		if configSetName != "" {
			o, err := c.SiteSiteGet(siteID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o.ConfigValues[configSetName] = configSetValue
			if err := o.Update([]string{"config_values"}); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputKV(o.ConfigValues)

		} else if configDeleteName != "" {
			o, err := c.SiteSiteGet(siteID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			delete(o.ConfigValues, configDeleteName)
			if err := o.Update([]string{"config_values"}); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputKV(o.ConfigValues)

		} else if configFull {
			o := c.SiteSiteNewWithID(siteID)
			r, err := o.CallGetConfig()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputKV(r)

		} else {
			o, err := c.SiteSiteGet(siteID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputKV(o.ConfigValues)
		}
	},
}

func init() {
	siteConfigCmd.Flags().BoolVarP(&configFull, "full", "f", false, "Display the Full/Compiled config")
	siteConfigCmd.Flags().StringVarP(&configSetName, "set-name", "n", "", "Set Config Value Key Name, if set-value is not specified, the value will be set to ''")
	siteConfigCmd.Flags().StringVarP(&configSetValue, "set-value", "v", "", "Set Config Value, ignored if set-name is not specified")
	siteConfigCmd.Flags().StringVarP(&configDeleteName, "delete", "d", "", "Delete Config Value Key Name")

	siteCreateCmd.Flags().StringVarP(&detailName, "name", "n", "", "Name of New Site")
	siteCreateCmd.Flags().StringVarP(&detailDescription, "description", "d", "", "Description of New Site")
	siteCreateCmd.Flags().StringVarP(&detailParent, "parent", "p", "", "Parent of New Site")
	siteCreateCmd.Flags().StringVarP(&detailZone, "zone", "z", "", "Zone of New Site")

	siteUpdateCmd.Flags().StringVarP(&detailDescription, "description", "d", "", "Update the Description of Site with value")
	siteUpdateCmd.Flags().StringVarP(&detailParent, "parent", "p", "", "Update the Parent of Site with value")
	siteUpdateCmd.Flags().StringVarP(&detailZone, "zone", "z", "", "Update the Zone of Site with value")

	rootCmd.AddCommand(siteCmd)
	siteCmd.AddCommand(siteListCmd)
	siteCmd.AddCommand(siteGetCmd)
	siteCmd.AddCommand(siteCreateCmd)
	siteCmd.AddCommand(siteUpdateCmd)
	siteCmd.AddCommand(siteDeleteCmd)
	siteCmd.AddCommand(siteConfigCmd)
}