package seed

const singleFetch = `return r.Client.{{Parent}}(prisma.{{Parent}}WhereUniqueInput{ID: &obj.ID}).{{Branch}}().Exec(ctx)`

const arrayFetch = `entries, err := r.Client.{{Parent}}(prisma.{{Parent}}WhereUniqueInput{ID: &obj.ID}).{{Branch}}(nil).Exec(ctx)
if err != nil {
	return nil, err
}

final := make([]*prisma.{{BranchType}}, 0)
for i := 0; i < len(entries); i++ {
	final = append(final, &entries[i])
}

return final, nil`
