package initbundles

import (
	"context"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	pkgCommon "github.com/stackrox/stackrox/pkg/roxctl/common"
	"github.com/stackrox/stackrox/pkg/utils"
	"github.com/stackrox/stackrox/roxctl/common/environment"
	"github.com/stackrox/stackrox/roxctl/common/util"
)

func listInitBundles(cliEnvironment environment.Environment) error {
	ctx, cancel := context.WithTimeout(pkgCommon.Context(), contextTimeout)
	defer cancel()

	conn, err := cliEnvironment.GRPCConnection()
	if err != nil {
		return err
	}
	defer utils.IgnoreError(conn.Close)
	svc := v1.NewClusterInitServiceClient(conn)

	tabWriter := tabwriter.NewWriter(os.Stdout, 4, 8, 2, '\t', 0)

	rsp, err := svc.GetInitBundles(ctx, &v1.Empty{})
	if err != nil {
		return errors.Wrap(err, "getting all init bundles")
	}

	bundles := rsp.GetItems()
	sort.Slice(bundles, func(i, j int) bool { return bundles[i].GetName() < bundles[j].GetName() })

	fmt.Fprintln(tabWriter, "Name\tCreated at\tExpires at\tCreated By\tID")
	fmt.Fprintln(tabWriter, "====\t==========\t==========\t==========\t==")

	for _, meta := range bundles {
		name := meta.GetName()
		if name == "" {
			name = "(empty)"
		}
		fmt.Fprintf(tabWriter, "%s\t%s\t%v\t%s\t%v\n",
			name,
			meta.GetCreatedAt(),
			meta.GetExpiresAt(),
			getPrettyUser(meta.GetCreatedBy()),
			meta.GetId(),
		)
	}
	return errors.Wrap(tabWriter.Flush(), "flushing tabular output")
}

// listCommand implements the command for listing init bundles.
func listCommand(cliEnvironment environment.Environment) *cobra.Command {
	c := &cobra.Command{
		Use: "list",
		RunE: util.RunENoArgs(func(c *cobra.Command) error {
			return listInitBundles(cliEnvironment)
		}),
	}
	return c
}
