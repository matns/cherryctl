/*
Copyright © 2022 Cherry Severs <support@cherryservers.com>
*/
package cmd

import (
	"fmt"

	"github.com/cherryservers/cherryctl/internal/backups"
	root "github.com/cherryservers/cherryctl/internal/cli"
	"github.com/cherryservers/cherryctl/internal/docs"
	"github.com/cherryservers/cherryctl/internal/images"
	initPck "github.com/cherryservers/cherryctl/internal/init"
	"github.com/cherryservers/cherryctl/internal/ips"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherryctl/internal/plans"
	"github.com/cherryservers/cherryctl/internal/projects"
	"github.com/cherryservers/cherryctl/internal/regions"
	"github.com/cherryservers/cherryctl/internal/servers"
	"github.com/cherryservers/cherryctl/internal/sshkeys"
	"github.com/cherryservers/cherryctl/internal/storages"
	"github.com/cherryservers/cherryctl/internal/teams"
	"github.com/cherryservers/cherryctl/internal/users"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
)

var Version string = "dev"

type Cli struct {
	MainCmd  *cobra.Command
	Outputer outputs.Outputer
}

func NewCli() (*Cli, error) {
	cli := &Cli{
		Outputer: &outputs.Standard{},
	}

	rootClient := root.NewClient(Version)

	rootCmd, err := rootClient.NewCommand()
	if err != nil {
		return nil, err
	}

	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd

	err = cli.RegisterCommands(rootClient)
	if err != nil {
		return nil, err
	}

	cobra.OnInitialize(
		func() {
			cli.Outputer.SetFormat(rootClient.Format())
		},
	)
	return cli, nil
}

type sharedDeps struct {
	out    outputs.Outputer
	client *root.Client
}

func (d sharedDeps) GetOpts() *cherrygo.GetOptions {
	return d.client.GetOptions()
}

func (d sharedDeps) Outputer() outputs.Outputer {
	return d.out
}

type planDeps struct {
	sharedDeps
}

func (d *planDeps) Client() cherrygo.PlansService {
	// The API method doesn't actually use the command for anything, so we can pass nil.
	// Should refactor it out at some point.
	return d.client.API(nil).Plans
}

type serverDeps struct {
	sharedDeps
}

func (d *serverDeps) Client() cherrygo.ServersService {
	return d.client.API(nil).Servers
}

func (cli *Cli) RegisterCommands(client *root.Client) error {
	shared := sharedDeps{
		out:    cli.Outputer,
		client: client,
	}

	registerCommandIfValid := func(cmd *cobra.Command, err error) error {
		if err != nil {
			return err
		}

		cli.MainCmd.AddCommand(cmd)
		return nil
	}

	if err := registerCommandIfValid(docs.NewCommand()); err != nil {
		return fmt.Errorf("register docs command: %w", err)
	}

	if err := registerCommandIfValid(initPck.NewClient(client).NewCommand()); err != nil {
		return fmt.Errorf("register init command: %w", err)
	}

	if err := registerCommandIfValid(servers.NewCommand(&serverDeps{sharedDeps: shared}).CobraCommand()); err != nil {
		return fmt.Errorf("register servers command: %w", err)
	}

	if err := registerCommandIfValid(ips.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register ips command: %w", err)
	}

	if err := registerCommandIfValid(storages.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register storages command: %w", err)
	}

	if err := registerCommandIfValid(backups.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register backups command: %w", err)
	}

	if err := registerCommandIfValid(regions.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register regions command: %w", err)
	}

	// We don't have the dependencies initialized yet, as that's done on pre-execution, so we need an injector interface.
	if err := registerCommandIfValid(plans.NewCommand(&planDeps{sharedDeps: shared}).CobraCommand()); err != nil {
		return fmt.Errorf("register plans command: %w", err)
	}

	if err := registerCommandIfValid(projects.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register projects command: %w", err)
	}

	if err := registerCommandIfValid(teams.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register teams command: %w", err)
	}

	if err := registerCommandIfValid(sshkeys.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register sshkeys command: %w", err)
	}

	if err := registerCommandIfValid(images.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register images command: %w", err)
	}

	if err := registerCommandIfValid(users.NewClient(client, cli.Outputer).NewCommand()); err != nil {
		return fmt.Errorf("register users command: %w", err)
	}

	return nil
}
