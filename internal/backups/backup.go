package backups

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer      Servicer
	Service       cherrygo.BackupsService
	ServerService cherrygo.ServersService
	Out           outputs.Outputer
}

func (c *Client) NewCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     `backup`,
		Aliases: []string{"backups"},
		Short:   "Server backup operations. For more information on backups, check the Product Docs: https://docs.cherryservers.com/knowledge/backup-storage",
		Long:    "Server backup storage operations: plans, get, list, create, update, methods, update-method, method-whitelist and remove.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Backups
			c.ServerService = c.Servicer.API(cmd).Servers
		},
	}

	plans, err := c.Plans()
	if err != nil {
		return nil, err
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

	update, err := c.Update()
	if err != nil {
		return nil, err
	}

	upgrade, err := c.Upgrade()
	if err != nil {
		return nil, err
	}

	methods, err := c.Methods()
	if err != nil {
		return nil, err
	}

	methodsWhitelist, err := c.MethodsWhitelist()
	if err != nil {
		return nil, err
	}

	updateMethod, err := c.UpdateMethod()
	if err != nil {
		return nil, err
	}

	delete, err := c.Delete()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(
		plans,
		get,
		list,
		create,
		update,
		upgrade,
		methods,
		methodsWhitelist,
		updateMethod,
		delete,
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
