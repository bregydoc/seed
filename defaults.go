package seed

var defaultPrismaProcessor = prismaProcessor{
	Endpoint: "https://us1.prisma.sh/bregy-malpartida-2d1dcf/xxx/dev",
	Datamodel: []string{
		"prisma.graphql",
	},
	Generate: []pGenerator{
		{Generator: "go-client", Output: "../client"},
	},
}

var res = "Resolver"

var defaultGQLGenProcessor = gqlGenProcessor{
	Schema: []string{
		"gqlgen.graphql",
		"literals.graphql",
	},
	Exec: gqlGenField{
		Filename: "generated.go",
	},
	Model: gqlGenField{
		Filename: "models_gen.go",
	},
	Resolver: gqlGenField{
		Filename: "resolver-temp.go",
		Type:     &res,
	},
	Models: map[string]gqlGenModel{},
}
