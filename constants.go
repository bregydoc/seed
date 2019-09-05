package seed

const queryType = "Query"
const mutationType = "Mutation"
const subscriptionType = "Subscription"
const inputType = "input"

const seedDir = "seed"
const prismaDir = "prisma"
const gqlDir = "gqlgen"

const prismaFile = "prisma.graphql"
const gqlGenFile = "gqlgen.graphql"
const gqlLiteralsGenFile = "literals.graphql"

const prismaConfigFile = "prisma.yml"
const gqlConfigFile = "gqlgen.yml"

const goModFilename = "go.mod"

const resolverTempFilename = "resolver-temp.go"

const resolverFilename = "resolver.go"

const notImplementedPanicLine = `panic("not implemented")`
const queryResolver = "query"

const toImplementFilename = "implement.go"
