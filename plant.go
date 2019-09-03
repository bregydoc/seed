package seed

// Plant returns a segmented gqltypes from your seed file
func Plant(seedFilename string, projectName string) error {
	workDir := "."

	err := createSeedDir(workDir)
	if err != nil {
		return err
	}

	err = createPrismaDir(workDir)
	if err != nil {
		return err
	}

	err = createGQLDir(workDir)
	if err != nil {
		return err
	}

	types, err := segment(seedFilename)
	if err != nil {
		return err
	}

	toPrisma := []GQLType{}
	toGQLGen := []GQLType{}

	// filtering
	for _, t := range types {
		if t.Name == queryType || t.Name == mutationType || t.Name == subscriptionType {
			toGQLGen = append(toGQLGen, t)
		} else {
			toPrisma = append(toPrisma, t)
		}
	}

	// writing

	err = writeToPrisma(workDir, toPrisma)
	if err != nil {
		return err
	}

	err = writeToGQLGen(workDir, toGQLGen)
	if err != nil {
		return err
	}

	err = defaultPrismaProcessor.writeConfig(workDir)
	if err != nil {
		return err
	}

	if !goModExists(workDir) {
		err = createGoModFile(workDir, projectName)
		if err != nil {
			return err
		}
	}

	err = defaultGQLGenProcessor.processAndWrite(workDir, projectName, toPrisma)
	if err != nil {
		return err
	}

	err = defaultGQLGenProcessor.generateCode(workDir)
	if err != nil {
		return err
	}

	return nil
}
