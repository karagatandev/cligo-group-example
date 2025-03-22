package main

import (
	"go.arpabet.com/cligo"
	"go.arpabet.com/glue"
)

type Plane struct {
	Parent cligo.CliGroup `cli:"group=cli"`
}

func (g *Plane) Group() string {
	return "plane"
}

func (g *Plane) Help() (string, string) {
	return `Manages planes.`, ""
}

type Launch struct {
	Parent cligo.CliGroup `cli:"group=plane"`
	Name   string         `cli:"argument=name"`
}

func (cmd *Launch) Command() string {
	return "launch"
}

func (cmd *Launch) Help() (string, string) {
	return "Launch the plane.", `This command launch the plane.
Do not launch the already launched planes.`
}

func (cmd *Launch) Run(ctx glue.Context) error {
	cligo.Echo("Launch the plane '%s'", cmd.Name)
	return nil
}

type Land struct {
	Parent cligo.CliGroup `cli:"group=plane"`
	Name   string         `cli:"argument=name"`
}

func (cmd *Land) Command() string {
	return "land"
}

func (cmd *Land) Help() (string, string) {
	return "Land the plane.", `This command lands the plane.
Do not land the already landed planes.`
}

func (cmd *Land) Run(ctx glue.Context) error {
	cligo.Echo("Land the plane '%s'", cmd.Name)
	return nil
}

type Version struct {
	Parent      cligo.CliGroup       `cli:"group=cli"`
	Application cligo.CliApplication `inject:""`
}

func (cmd *Version) Command() string {
	return "version"
}

func (cmd *Version) Help() (string, string) {
	return "Prints version of the app.", ""
}

func (cmd *Version) Run(ctx glue.Context) error {
	cligo.Echo("Version '%s'", cmd.Application.Version())
	return nil
}

func main() {
	cligo.Main(cligo.Version("1.0.0"), cligo.Beans(&Plane{}, &Launch{}, &Land{}, &Version{}))
}
