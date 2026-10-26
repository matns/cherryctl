package teams

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer Servicer
	Service  cherrygo.TeamsService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     `team`,
		Aliases: []string{"teams", "organization", "organizations"},
		Short:   "Team operations. For more information on Teams, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#teams",
		Long:    "Team operations: get, list.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Teams
		},
	}

	get, err := c.Get()
	if err != nil {
		return nil, err
	}

	list, err := c.List()
	if err != nil {
		return nil, err
	}

	create, err := c.Create()
	if err != nil {
		return nil, err
	}

	delete, err := c.Delete()
	if err != nil {
		return nil, err
	}

	update, err := c.Update()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(
		get,
		list,
		create,
		delete,
		update,
	)

	return cmd, nil
}

type Servicer interface {
	API(*cobra.Command) *cherrygo.Client
	GetOptions() *cherrygo.GetOptions
	Config(cmd *cobra.Command) *viper.Viper
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
