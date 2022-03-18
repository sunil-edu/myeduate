package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/msgchannelmessage"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddChannelMessage(ctx context.Context, token string, input ent.CreateMsgChannelMessageInput) (*ent.MsgChannelMessage, error) {
	client := ent.FromContext(ctx)

	msg, err := client.MsgChannelMessage.Create().SetInput(input).Save(ctx)

	if err != nil {
		return nil, err
	}
	r.ChatMessages = append(r.ChatMessages, msg)

	r.mutex.Lock()
	// Notify all active subscriptions that a new message has been posted by posted. In this case we push the now
	// updated ChatMessages array to all clients that care about it.
	//r.ChatObservers["0"] <- r.ChatMessages
	for _, observer := range r.ChatObservers {

		observer <- r.ChatMessages

	}
	r.mutex.Unlock()

	return msg, err
}

func (r *queryResolver) GetChannelMessages(ctx context.Context, token string) ([]*ent.MsgChannelMessage, error) {
	return r.client.MsgChannelMessage.Query().Order(ent.Asc(msgchannelmessage.FieldCreatedAt)).All(ctx)
}

func (r *subscriptionResolver) GetChannelMessagesBySubscription(ctx context.Context, token string) (<-chan []*ent.MsgChannelMessage, error) {
	// Create an ID and channel for each active subscription. We will push changes into this channel.
	id := uuid.New().String()
	msgs := make(chan []*ent.MsgChannelMessage, 1)

	// Start a goroutine to allow for cleaning up subscriptions that are disconnected.
	// This go routine will only get past Done() when a client terminates the subscription. This allows us
	// to only then remove the reference from the list of ChatObservers since it is no longer needed.
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.ChatObservers, id)
		r.mutex.Unlock()
	}()
	r.mutex.Lock()
	// Keep a reference of the channel so that we can push changes into it when new messages are posted.
	r.ChatObservers[id] = msgs
	r.mutex.Unlock()
	// This is optional, and this allows newly subscribed clients to get a list of all the messages that have been
	// posted so far. Upon subscribing the client will be pushed the messages once, further changes are handled
	// in the PostMessage mutation.
	r.ChatObservers[id] <- r.ChatMessages

	return r.ChatObservers[id], nil
}

// Subscription returns myeduate.SubscriptionResolver implementation.
func (r *Resolver) Subscription() myeduate.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
