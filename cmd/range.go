package cmd

import (
	"fmt"

	"github.com/3th1nk/cidr"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use: "print",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cidr.Parse(args[0])
		if err != nil {
			return err
		}
		c.Each(func(ip string) bool {
			fmt.Println(ip)
			return true
		})
		return nil
	},
}

var rangeCmd = &cobra.Command{
	Use: "range",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cidr.Parse(args[0])
		if err != nil {
			return err
		}
		begin, end := c.IPRange()
		fmt.Printf("%s %s\n", begin.String(), end.String())
		return nil
	},
}

var countCmd = &cobra.Command{
	Use: "count",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cidr.Parse(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%d\n", c.IPCount())
		return nil
	},
}

var subnetCmd = &cobra.Command{
	Use: "subnet",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cidr.Parse(args[0])
		if err != nil {
			return err
		}
		mode := cidr.MethodSubnetNum
		if ok, err := cmd.PersistentFlags().GetBool("host"); err == nil && ok {
			mode = cidr.MethodHostNum
		}
		if num, err := cmd.PersistentFlags().GetInt("num"); err == nil {
			ranges, err := c.SubNetting(mode, num)
			if err != nil {
				return err
			}
			for _, subrange := range ranges {
				fmt.Println(subrange.CIDR().String())
			}
		}

		return nil
	},
}

func init() {
	subnetCmd.PersistentFlags().Int("num", 1, "split the cidr into this many chunks")
	subnetCmd.PersistentFlags().Bool("host", false, "split the cidr into N hosts")
	rootCmd.AddCommand(
		printCmd,
		rangeCmd,
		countCmd,
		subnetCmd,
	)
}
