package seed

import (
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

// Plant returns a segmented gqltypes from your seed file
func (p *Project) Plant() error {
	spinner := spin.Line
	w := wow.New(os.Stdout, spin.Get(spinner), "Creating seed directory")
	w.Start()

	err := createSeedDir(p)
	if err != nil {
		return err
	}

	w.Text("Creating Prisma directory").Spinner(spin.Get(spinner))
	err = createPrismaDir(p)
	if err != nil {
		return err
	}

	w.Text("Creating GQLGen directory").Spinner(spin.Get(spinner))
	err = createGQLDir(p)
	if err != nil {
		return err
	}

	w.Text("Segmenting graphql types").Spinner(spin.Get(spinner))
	types, err := segment(p)
	if err != nil {
		return err
	}

	toPrisma := []GQLType{}
	toGQLGen := []GQLType{}

	// filtering
	for _, t := range types {
		if t.Name == queryType || t.Name == mutationType || t.Name == subscriptionType || t.Name == inputType {
			toGQLGen = append(toGQLGen, t)
		} else {
			toPrisma = append(toPrisma, t)
		}
	}

	// writing
	w.Text("Writing Prisma models").Spinner(spin.Get(spinner))
	err = writeToPrisma(p, toPrisma)
	if err != nil {
		return err
	}

	w.Text("Writing GQLGen query/mutations models").Spinner(spin.Get(spinner))
	err = writeToGQLGen(p, toGQLGen)
	if err != nil {
		return err
	}

	w.Text("Fixing go mod file").Spinner(spin.Get(spinner))
	if !goModExists(p) {
		err = createGoModFile(p)
		if err != nil {
			return err
		}
	}

	w.Text("Writing prisma config file").Spinner(spin.Get(spinner))
	if err = defaultPrismaProcessor.writeConfig(p); err != nil {
		return err
	}

	w.Text("Generating prisma client file").Spinner(spin.Get(spinner))
	if err = defaultPrismaProcessor.generate(p); err != nil {
		return err
	}

	w.Text("Processing and writing GQLGen file").Spinner(spin.Get(spinner))
	if err = defaultGQLGenProcessor.processAndWrite(p, toPrisma); err != nil {
		return err
	}

	w.Text("Generating GQLGen models and resolvers").Spinner(spin.Get(spinner))
	if err = defaultGQLGenProcessor.generateCode(p); err != nil {
		return err
	}

	w.Text("Fixing and writing GQLGen new models and resolvers").Spinner(spin.Get(spinner))
	if err = defaultGQLGenProcessor.fixResolver(p); err != nil {
		return err
	}

	w.Text("Cleaning GQLGen unimplemented models").Spinner(spin.Get(spinner))
	if err = defaultGQLGenProcessor.cleanUnimplemented(p); err != nil {
		return err
	}

	return nil
}
