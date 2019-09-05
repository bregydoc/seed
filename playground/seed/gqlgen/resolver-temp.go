package gqlgen

import (
	"context"

	prisma "github.com/playground/seed/client"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *Resolver) Attraction() AttractionResolver {
	return &attractionResolver{r}
}
func (r *Resolver) Cost() CostResolver {
	return &costResolver{r}
}
func (r *Resolver) Credential() CredentialResolver {
	return &credentialResolver{r}
}
func (r *Resolver) CreditCard() CreditCardResolver {
	return &creditCardResolver{r}
}
func (r *Resolver) Currency() CurrencyResolver {
	return &currencyResolver{r}
}
func (r *Resolver) EmailChannel() EmailChannelResolver {
	return &emailChannelResolver{r}
}
func (r *Resolver) EntryTicket() EntryTicketResolver {
	return &entryTicketResolver{r}
}
func (r *Resolver) Error() ErrorResolver {
	return &errorResolver{r}
}
func (r *Resolver) Expert() ExpertResolver {
	return &expertResolver{r}
}
func (r *Resolver) Faq() FaqResolver {
	return &faqResolver{r}
}
func (r *Resolver) FaqEntry() FaqEntryResolver {
	return &faqEntryResolver{r}
}
func (r *Resolver) Feature() FeatureResolver {
	return &featureResolver{r}
}
func (r *Resolver) IDDocument() IDDocumentResolver {
	return &iDDocumentResolver{r}
}
func (r *Resolver) Lima() LimaResolver {
	return &limaResolver{r}
}
func (r *Resolver) LimaContact() LimaContactResolver {
	return &limaContactResolver{r}
}
func (r *Resolver) Location() LocationResolver {
	return &locationResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) OpeningHours() OpeningHoursResolver {
	return &openingHoursResolver{r}
}
func (r *Resolver) Partner() PartnerResolver {
	return &partnerResolver{r}
}
func (r *Resolver) Pass() PassResolver {
	return &passResolver{r}
}
func (r *Resolver) PassType() PassTypeResolver {
	return &passTypeResolver{r}
}
func (r *Resolver) Person() PersonResolver {
	return &personResolver{r}
}
func (r *Resolver) PersonReviews() PersonReviewsResolver {
	return &personReviewsResolver{r}
}
func (r *Resolver) PhoneChannel() PhoneChannelResolver {
	return &phoneChannelResolver{r}
}
func (r *Resolver) Place() PlaceResolver {
	return &placeResolver{r}
}
func (r *Resolver) Purchase() PurchaseResolver {
	return &purchaseResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Route() RouteResolver {
	return &routeResolver{r}
}
func (r *Resolver) Story() StoryResolver {
	return &storyResolver{r}
}
func (r *Resolver) Ticket() TicketResolver {
	return &ticketResolver{r}
}
func (r *Resolver) Traveler() TravelerResolver {
	return &travelerResolver{r}
}

type attractionResolver struct{ *Resolver }

func (r *attractionResolver) Owner(ctx context.Context, obj *prisma.Attraction) (*prisma.Partner, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).Owner().Exec(ctx)
}
func (r *attractionResolver) Name(ctx context.Context, obj *prisma.Attraction) (*prisma.Text, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).Name().Exec(ctx)
}
func (r *attractionResolver) Description(ctx context.Context, obj *prisma.Attraction) (*prisma.Text, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}
func (r *attractionResolver) ShortDescription(ctx context.Context, obj *prisma.Attraction) (*prisma.Text, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).ShortDescription().Exec(ctx)
}
func (r *attractionResolver) Faqs(ctx context.Context, obj *prisma.Attraction) (*prisma.Faq, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).Faqs().Exec(ctx)
}
func (r *attractionResolver) Place(ctx context.Context, obj *prisma.Attraction) (*prisma.Place, error) {
	return r.Client.Attraction(prisma.AttractionWhereUniqueInput{ID: &obj.ID}).Place().Exec(ctx)
}

type costResolver struct{ *Resolver }

func (r *costResolver) Currency(ctx context.Context, obj *prisma.Cost) (*prisma.Currency, error) {
	return r.Client.Cost(prisma.CostWhereUniqueInput{ID: &obj.ID}).Currency().Exec(ctx)
}

type credentialResolver struct{ *Resolver }

func (r *credentialResolver) Owner(ctx context.Context, obj *prisma.Credential) (*prisma.Person, error) {
	return r.Client.Credential(prisma.CredentialWhereUniqueInput{ID: &obj.ID}).Owner().Exec(ctx)
}

type creditCardResolver struct{ *Resolver }

func (r *creditCardResolver) Owner(ctx context.Context, obj *prisma.CreditCard) (*prisma.Traveler, error) {
	return r.Client.CreditCard(prisma.CreditCardWhereUniqueInput{ID: &obj.ID}).Owner().Exec(ctx)
}

type currencyResolver struct{ *Resolver }

func (r *currencyResolver) Name(ctx context.Context, obj *prisma.Currency) (*prisma.Text, error) {
	return r.Client.Currency(prisma.CurrencyWhereUniqueInput{ID: &obj.ID}).Name().Exec(ctx)
}

type emailChannelResolver struct{ *Resolver }

func (r *emailChannelResolver) Person(ctx context.Context, obj *prisma.EmailChannel) (*prisma.Person, error) {
	return r.Client.EmailChannel(prisma.EmailChannelWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}

type entryTicketResolver struct{ *Resolver }

func (r *entryTicketResolver) Ticket(ctx context.Context, obj *prisma.EntryTicket) (*prisma.Ticket, error) {
	return r.Client.EntryTicket(prisma.EntryTicketWhereUniqueInput{ID: &obj.ID}).Ticket().Exec(ctx)
}
func (r *entryTicketResolver) Partner(ctx context.Context, obj *prisma.EntryTicket) (*prisma.Partner, error) {
	return r.Client.EntryTicket(prisma.EntryTicketWhereUniqueInput{ID: &obj.ID}).Partner().Exec(ctx)
}
func (r *entryTicketResolver) Traveler(ctx context.Context, obj *prisma.EntryTicket) (*prisma.Traveler, error) {
	return r.Client.EntryTicket(prisma.EntryTicketWhereUniqueInput{ID: &obj.ID}).Traveler().Exec(ctx)
}

type errorResolver struct{ *Resolver }

func (r *errorResolver) Message(ctx context.Context, obj *prisma.Error) (*prisma.Text, error) {
	return r.Client.Error(prisma.ErrorWhereUniqueInput{ID: &obj.ID}).Message().Exec(ctx)
}

type expertResolver struct{ *Resolver }

func (r *expertResolver) Person(ctx context.Context, obj *prisma.Expert) (*prisma.Person, error) {
	return r.Client.Expert(prisma.ExpertWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}
func (r *expertResolver) Stories(ctx context.Context, obj *prisma.Expert) ([]*prisma.Story, error) {
	entries, err := r.Client.Expert(prisma.ExpertWhereUniqueInput{ID: &obj.ID}).Stories(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Story, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type faqResolver struct{ *Resolver }

func (r *faqResolver) Attraction(ctx context.Context, obj *prisma.Faq) (*prisma.Attraction, error) {
	return r.Client.Faq(prisma.FaqWhereUniqueInput{ID: &obj.ID}).Attraction().Exec(ctx)
}
func (r *faqResolver) Entries(ctx context.Context, obj *prisma.Faq) ([]*prisma.FaqEntry, error) {
	entries, err := r.Client.Faq(prisma.FaqWhereUniqueInput{ID: &obj.ID}).Entries(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.FaqEntry, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type faqEntryResolver struct{ *Resolver }

func (r *faqEntryResolver) Faq(ctx context.Context, obj *prisma.FaqEntry) (*prisma.Faq, error) {
	return r.Client.FaqEntry(prisma.FaqEntryWhereUniqueInput{ID: &obj.ID}).Faq().Exec(ctx)
}
func (r *faqEntryResolver) Question(ctx context.Context, obj *prisma.FaqEntry) (*prisma.Text, error) {
	return r.Client.FaqEntry(prisma.FaqEntryWhereUniqueInput{ID: &obj.ID}).Question().Exec(ctx)
}
func (r *faqEntryResolver) Answer(ctx context.Context, obj *prisma.FaqEntry) (*prisma.Text, error) {
	return r.Client.FaqEntry(prisma.FaqEntryWhereUniqueInput{ID: &obj.ID}).Answer().Exec(ctx)
}

type featureResolver struct{ *Resolver }

func (r *featureResolver) Name(ctx context.Context, obj *prisma.Feature) (*prisma.Text, error) {
	return r.Client.Feature(prisma.FeatureWhereUniqueInput{ID: &obj.ID}).Name().Exec(ctx)
}

type iDDocumentResolver struct{ *Resolver }

func (r *iDDocumentResolver) Person(ctx context.Context, obj *prisma.IDDocument) (*prisma.Person, error) {
	return r.Client.IDDocument(prisma.IDDocumentWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}

type limaResolver struct{ *Resolver }

func (r *limaResolver) MainAttractions(ctx context.Context, obj *prisma.Lima) ([]*prisma.Attraction, error) {
	entries, err := r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).MainAttractions(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Attraction, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *limaResolver) MainPlaces(ctx context.Context, obj *prisma.Lima) ([]*prisma.Place, error) {
	entries, err := r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).MainPlaces(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Place, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *limaResolver) MainPartners(ctx context.Context, obj *prisma.Lima) ([]*prisma.Partner, error) {
	entries, err := r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).MainPartners(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Partner, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *limaResolver) HotStories(ctx context.Context, obj *prisma.Lima) ([]*prisma.Story, error) {
	entries, err := r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).HotStories(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Story, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *limaResolver) Contact(ctx context.Context, obj *prisma.Lima) (*prisma.LimaContact, error) {
	return r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).Contact().Exec(ctx)
}
func (r *limaResolver) Faq(ctx context.Context, obj *prisma.Lima) (*prisma.Faq, error) {
	return r.Client.Lima(prisma.LimaWhereUniqueInput{ID: &obj.ID}).Faq().Exec(ctx)
}

type limaContactResolver struct{ *Resolver }

func (r *limaContactResolver) Message(ctx context.Context, obj *prisma.LimaContact) (*prisma.Text, error) {
	return r.Client.LimaContact(prisma.LimaContactWhereUniqueInput{ID: &obj.ID}).Message().Exec(ctx)
}

type locationResolver struct{ *Resolver }

func (r *locationResolver) Description(ctx context.Context, obj *prisma.Location) (*prisma.Text, error) {
	return r.Client.Location(prisma.LocationWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}

type mutationResolver struct{ *Resolver }

type openingHoursResolver struct{ *Resolver }

func (r *openingHoursResolver) Description(ctx context.Context, obj *prisma.OpeningHours) (*prisma.Text, error) {
	return r.Client.OpeningHours(prisma.OpeningHoursWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}

type partnerResolver struct{ *Resolver }

func (r *partnerResolver) Attender(ctx context.Context, obj *prisma.Partner) (*prisma.Person, error) {
	return r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Attender().Exec(ctx)
}
func (r *partnerResolver) Owner(ctx context.Context, obj *prisma.Partner) (*prisma.Person, error) {
	return r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Owner().Exec(ctx)
}
func (r *partnerResolver) Attractions(ctx context.Context, obj *prisma.Partner) ([]*prisma.Attraction, error) {
	entries, err := r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Attractions(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Attraction, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *partnerResolver) Location(ctx context.Context, obj *prisma.Partner) (*prisma.Location, error) {
	return r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Location().Exec(ctx)
}
func (r *partnerResolver) Features(ctx context.Context, obj *prisma.Partner) ([]*prisma.Feature, error) {
	entries, err := r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Features(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Feature, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *partnerResolver) ServiceHours(ctx context.Context, obj *prisma.Partner) (*prisma.OpeningHours, error) {
	return r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).ServiceHours().Exec(ctx)
}
func (r *partnerResolver) Registry(ctx context.Context, obj *prisma.Partner) ([]*prisma.EntryTicket, error) {
	entries, err := r.Client.Partner(prisma.PartnerWhereUniqueInput{ID: &obj.ID}).Registry(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.EntryTicket, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type passResolver struct{ *Resolver }

func (r *passResolver) Type(ctx context.Context, obj *prisma.Pass) (*prisma.PassType, error) {
	return r.Client.Pass(prisma.PassWhereUniqueInput{ID: &obj.ID}).Type().Exec(ctx)
}
func (r *passResolver) Owner(ctx context.Context, obj *prisma.Pass) (*prisma.Person, error) {
	return r.Client.Pass(prisma.PassWhereUniqueInput{ID: &obj.ID}).Owner().Exec(ctx)
}
func (r *passResolver) PurchaseReceive(ctx context.Context, obj *prisma.Pass) (*prisma.Purchase, error) {
	return r.Client.Pass(prisma.PassWhereUniqueInput{ID: &obj.ID}).PurchaseReceive().Exec(ctx)
}
func (r *passResolver) Tickets(ctx context.Context, obj *prisma.Pass) ([]*prisma.Ticket, error) {
	entries, err := r.Client.Pass(prisma.PassWhereUniqueInput{ID: &obj.ID}).Tickets(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Ticket, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type passTypeResolver struct{ *Resolver }

func (r *passTypeResolver) Name(ctx context.Context, obj *prisma.PassType) (*prisma.Text, error) {
	return r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Name().Exec(ctx)
}
func (r *passTypeResolver) ShortDescription(ctx context.Context, obj *prisma.PassType) (*prisma.Text, error) {
	return r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).ShortDescription().Exec(ctx)
}
func (r *passTypeResolver) Description(ctx context.Context, obj *prisma.PassType) (*prisma.Text, error) {
	return r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}
func (r *passTypeResolver) Disclaimer(ctx context.Context, obj *prisma.PassType) (*prisma.Text, error) {
	return r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Disclaimer().Exec(ctx)
}
func (r *passTypeResolver) Attractions(ctx context.Context, obj *prisma.PassType) ([]*prisma.Attraction, error) {
	entries, err := r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Attractions(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Attraction, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *passTypeResolver) Cost(ctx context.Context, obj *prisma.PassType) (*prisma.Cost, error) {
	return r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Cost().Exec(ctx)
}
func (r *passTypeResolver) Children(ctx context.Context, obj *prisma.PassType) ([]*prisma.Pass, error) {
	entries, err := r.Client.PassType(prisma.PassTypeWhereUniqueInput{ID: &obj.ID}).Children(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Pass, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type personResolver struct{ *Resolver }

func (r *personResolver) Country(ctx context.Context, obj *prisma.Person) (*prisma.Text, error) {
	return r.Client.Person(prisma.PersonWhereUniqueInput{ID: &obj.ID}).Country().Exec(ctx)
}
func (r *personResolver) Phones(ctx context.Context, obj *prisma.Person) ([]*prisma.PhoneChannel, error) {
	entries, err := r.Client.Person(prisma.PersonWhereUniqueInput{ID: &obj.ID}).Phones(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.PhoneChannel, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *personResolver) Emails(ctx context.Context, obj *prisma.Person) ([]*prisma.EmailChannel, error) {
	entries, err := r.Client.Person(prisma.PersonWhereUniqueInput{ID: &obj.ID}).Emails(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.EmailChannel, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *personResolver) Identification(ctx context.Context, obj *prisma.Person) (*prisma.IDDocument, error) {
	return r.Client.Person(prisma.PersonWhereUniqueInput{ID: &obj.ID}).Identification().Exec(ctx)
}
func (r *personResolver) Credentials(ctx context.Context, obj *prisma.Person) ([]*prisma.Credential, error) {
	entries, err := r.Client.Person(prisma.PersonWhereUniqueInput{ID: &obj.ID}).Credentials(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Credential, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type personReviewsResolver struct{ *Resolver }

func (r *personReviewsResolver) Person(ctx context.Context, obj *prisma.PersonReviews) (*prisma.Person, error) {
	return r.Client.PersonReviews(prisma.PersonReviewsWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}
func (r *personReviewsResolver) Title(ctx context.Context, obj *prisma.PersonReviews) (*prisma.Text, error) {
	return r.Client.PersonReviews(prisma.PersonReviewsWhereUniqueInput{ID: &obj.ID}).Title().Exec(ctx)
}
func (r *personReviewsResolver) Description(ctx context.Context, obj *prisma.PersonReviews) (*prisma.Text, error) {
	return r.Client.PersonReviews(prisma.PersonReviewsWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}
func (r *personReviewsResolver) Comments(ctx context.Context, obj *prisma.PersonReviews) ([]*prisma.PersonReviews, error) {
	entries, err := r.Client.PersonReviews(prisma.PersonReviewsWhereUniqueInput{ID: &obj.ID}).Comments(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.PersonReviews, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type phoneChannelResolver struct{ *Resolver }

func (r *phoneChannelResolver) Person(ctx context.Context, obj *prisma.PhoneChannel) (*prisma.Person, error) {
	return r.Client.PhoneChannel(prisma.PhoneChannelWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}

type placeResolver struct{ *Resolver }

func (r *placeResolver) Description(ctx context.Context, obj *prisma.Place) (*prisma.Text, error) {
	return r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}
func (r *placeResolver) ShortDescription(ctx context.Context, obj *prisma.Place) (*prisma.Text, error) {
	return r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).ShortDescription().Exec(ctx)
}
func (r *placeResolver) Location(ctx context.Context, obj *prisma.Place) (*prisma.Location, error) {
	return r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).Location().Exec(ctx)
}
func (r *placeResolver) Routes(ctx context.Context, obj *prisma.Place) ([]*prisma.Route, error) {
	entries, err := r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).Routes(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Route, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *placeResolver) Reviews(ctx context.Context, obj *prisma.Place) ([]*prisma.PersonReviews, error) {
	entries, err := r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).Reviews(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.PersonReviews, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *placeResolver) LikedBy(ctx context.Context, obj *prisma.Place) ([]*prisma.Person, error) {
	entries, err := r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).LikedBy(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Person, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *placeResolver) Ambassadors(ctx context.Context, obj *prisma.Place) ([]*prisma.Expert, error) {
	entries, err := r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).Ambassadors(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Expert, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *placeResolver) OpeningHours(ctx context.Context, obj *prisma.Place) (*prisma.OpeningHours, error) {
	return r.Client.Place(prisma.PlaceWhereUniqueInput{ID: &obj.ID}).OpeningHours().Exec(ctx)
}

type purchaseResolver struct{ *Resolver }

func (r *purchaseResolver) Products(ctx context.Context, obj *prisma.Purchase) ([]*prisma.ProductEntry, error) {
	entries, err := r.Client.Purchase(prisma.PurchaseWhereUniqueInput{ID: &obj.ID}).Products(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.ProductEntry, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *purchaseResolver) PaymentMethod(ctx context.Context, obj *prisma.Purchase) (*prisma.PaymentMethod, error) {
	return r.Client.Purchase(prisma.PurchaseWhereUniqueInput{ID: &obj.ID}).PaymentMethod().Exec(ctx)
}

type queryResolver struct{ *Resolver }

type routeResolver struct{ *Resolver }

func (r *routeResolver) Creator(ctx context.Context, obj *prisma.Route) (*prisma.Expert, error) {
	return r.Client.Route(prisma.RouteWhereUniqueInput{ID: &obj.ID}).Creator().Exec(ctx)
}
func (r *routeResolver) Places(ctx context.Context, obj *prisma.Route) ([]*prisma.Place, error) {
	entries, err := r.Client.Route(prisma.RouteWhereUniqueInput{ID: &obj.ID}).Places(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Place, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *routeResolver) LinkedStory(ctx context.Context, obj *prisma.Route) (*prisma.Story, error) {
	return r.Client.Route(prisma.RouteWhereUniqueInput{ID: &obj.ID}).LinkedStory().Exec(ctx)
}
func (r *routeResolver) LikedBy(ctx context.Context, obj *prisma.Route) ([]*prisma.Person, error) {
	entries, err := r.Client.Route(prisma.RouteWhereUniqueInput{ID: &obj.ID}).LikedBy(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Person, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *routeResolver) Reviews(ctx context.Context, obj *prisma.Route) ([]*prisma.PersonReviews, error) {
	entries, err := r.Client.Route(prisma.RouteWhereUniqueInput{ID: &obj.ID}).Reviews(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.PersonReviews, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}

type storyResolver struct{ *Resolver }

func (r *storyResolver) Author(ctx context.Context, obj *prisma.Story) (*prisma.Expert, error) {
	return r.Client.Story(prisma.StoryWhereUniqueInput{ID: &obj.ID}).Author().Exec(ctx)
}
func (r *storyResolver) Title(ctx context.Context, obj *prisma.Story) (*prisma.Text, error) {
	return r.Client.Story(prisma.StoryWhereUniqueInput{ID: &obj.ID}).Title().Exec(ctx)
}
func (r *storyResolver) Body(ctx context.Context, obj *prisma.Story) (*prisma.Text, error) {
	return r.Client.Story(prisma.StoryWhereUniqueInput{ID: &obj.ID}).Body().Exec(ctx)
}
func (r *storyResolver) Reviews(ctx context.Context, obj *prisma.Story) ([]*prisma.PersonReviews, error) {
	entries, err := r.Client.Story(prisma.StoryWhereUniqueInput{ID: &obj.ID}).Reviews(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.PersonReviews, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *storyResolver) LinkedRoute(ctx context.Context, obj *prisma.Story) (*prisma.Route, error) {
	return r.Client.Story(prisma.StoryWhereUniqueInput{ID: &obj.ID}).LinkedRoute().Exec(ctx)
}

type ticketResolver struct{ *Resolver }

func (r *ticketResolver) Parent(ctx context.Context, obj *prisma.Ticket) (*prisma.Pass, error) {
	return r.Client.Ticket(prisma.TicketWhereUniqueInput{ID: &obj.ID}).Parent().Exec(ctx)
}
func (r *ticketResolver) Attraction(ctx context.Context, obj *prisma.Ticket) (*prisma.Attraction, error) {
	return r.Client.Ticket(prisma.TicketWhereUniqueInput{ID: &obj.ID}).Attraction().Exec(ctx)
}
func (r *ticketResolver) Description(ctx context.Context, obj *prisma.Ticket) (*prisma.Text, error) {
	return r.Client.Ticket(prisma.TicketWhereUniqueInput{ID: &obj.ID}).Description().Exec(ctx)
}
func (r *ticketResolver) AvailableTime(ctx context.Context, obj *prisma.Ticket) (*prisma.Text, error) {
	return r.Client.Ticket(prisma.TicketWhereUniqueInput{ID: &obj.ID}).AvailableTime().Exec(ctx)
}

type travelerResolver struct{ *Resolver }

func (r *travelerResolver) RegistrationTicket(ctx context.Context, obj *prisma.Traveler) (*prisma.ActionTicket, error) {
	return r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).RegistrationTicket().Exec(ctx)
}
func (r *travelerResolver) Person(ctx context.Context, obj *prisma.Traveler) (*prisma.Person, error) {
	return r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).Person().Exec(ctx)
}
func (r *travelerResolver) Cards(ctx context.Context, obj *prisma.Traveler) ([]*prisma.CreditCard, error) {
	entries, err := r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).Cards(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.CreditCard, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *travelerResolver) Purchases(ctx context.Context, obj *prisma.Traveler) ([]*prisma.Purchase, error) {
	entries, err := r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).Purchases(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Purchase, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *travelerResolver) Passes(ctx context.Context, obj *prisma.Traveler) ([]*prisma.Pass, error) {
	entries, err := r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).Passes(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.Pass, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
func (r *travelerResolver) Registry(ctx context.Context, obj *prisma.Traveler) ([]*prisma.EntryTicket, error) {
	entries, err := r.Client.Traveler(prisma.TravelerWhereUniqueInput{ID: &obj.ID}).Registry(nil).Exec(ctx)
	if err != nil {
		return nil, err
	}

	final := make([]*prisma.EntryTicket, 0)
	for i := 0; i < len(entries); i++ {
		final = append(final, &entries[i])
	}

	return final, nil
}
