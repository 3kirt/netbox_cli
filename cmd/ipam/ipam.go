package ipam

import (
  "github.com/spf13/cobra"

  "github.com/kirtis/netbox-cli/internal/clientctx"
  "github.com/kirtis/netbox-cli/internal/cmdutil"
)

func Command() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "ipam",
    Short: "Commands for the NetBox IPAM API",
  }

  cmd.AddCommand(
    aggregatesCmd(),
    prefixesCmd(),
    ipAddressesCmd(),
    ipRangesCmd(),
    vlansCmd(),
    vlanGroupsCmd(),
    vrfsCmd(),
    rolesCmd(),
    servicesCmd(),
  )

  return cmd
}

// aggregatesCmd -------------------------------------------------------

func aggregatesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "aggregates", Short: "Manage aggregates"}
  cmd.AddCommand(
    listCmd("aggregates", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamAggregatesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("aggregate", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamAggregatesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// prefixesCmd -------------------------------------------------------

func prefixesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "prefixes", Short: "Manage prefixes"}
  cmd.AddCommand(
    listCmd("prefixes", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamPrefixesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("prefix", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamPrefixesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// ipAddressesCmd -------------------------------------------------------

func ipAddressesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "ip-addresses", Short: "Manage IP addresses"}
  cmd.AddCommand(
    listCmd("ip-addresses", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamIpAddressesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("ip-address", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamIpAddressesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// ipRangesCmd -------------------------------------------------------

func ipRangesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "ip-ranges", Short: "Manage IP ranges"}
  cmd.AddCommand(
    listCmd("ip-ranges", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamIpRangesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("ip-range", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamIpRangesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// vlansCmd -------------------------------------------------------

func vlansCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "vlans", Short: "Manage VLANs"}
  cmd.AddCommand(
    listCmd("vlans", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVlansList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("vlan", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVlansRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// vlanGroupsCmd -------------------------------------------------------

func vlanGroupsCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "vlan-groups", Short: "Manage VLAN groups"}
  cmd.AddCommand(
    listCmd("vlan-groups", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVlanGroupsList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("vlan-group", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVlanGroupsRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// vrfsCmd -------------------------------------------------------

func vrfsCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "vrfs", Short: "Manage VRFs"}
  cmd.AddCommand(
    listCmd("vrfs", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVrfsList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("vrf", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamVrfsRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// rolesCmd -------------------------------------------------------

func rolesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "roles", Short: "Manage IP/VLAN roles"}
  cmd.AddCommand(
    listCmd("roles", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamRolesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("role", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamRolesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// servicesCmd -------------------------------------------------------

func servicesCmd() *cobra.Command {
  cmd := &cobra.Command{Use: "services", Short: "Manage services"}
  cmd.AddCommand(
    listCmd("services", func(cmd *cobra.Command) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamServicesList(cmd.Context()).Limit(0).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp.GetResults())
    }),
    getCmd("service", func(cmd *cobra.Command, id int32) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.IpamAPI.IpamServicesRetrieve(cmd.Context(), id).Execute()
      if err != nil {
        return cmdutil.APIError(err)
      }
      return cmdutil.OutputJSON(resp)
    }),
  )
  return cmd
}

// helpers -------------------------------------------------------

func listCmd(noun string, run func(cmd *cobra.Command) error) *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all " + noun,
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(cmd)
    },
  }
}

func getCmd(noun string, run func(cmd *cobra.Command, id int32) error) *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a " + noun + " by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(cmd, id)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, noun+" ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}
