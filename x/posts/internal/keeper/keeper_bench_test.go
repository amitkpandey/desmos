package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// RandomPostIDOrSubspace returns a random PostID
func RandomPostIDOrSubspace() types.PostID {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return types.PostID(hex.EncodeToString(hash[:]))
}

// RandomMessage returns a random String with len <= 500
func RandomMessage(r *rand.Rand) string {
	bytes := make([]byte, r.Intn(100))
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

// RandomPost returns a post with a 50% chance to have random medias and random poll
func RandomPost() types.Post {
	r := rand.New(rand.NewSource(100))
	accounts := sim.RandomAccounts(r, r.Intn(20))

	post := types.NewPost(
		RandomPostIDOrSubspace(),
		RandomPostIDOrSubspace(),
		RandomMessage(r),
		r.Intn(101) <= 50,
		RandomPostIDOrSubspace().String(),
		map[string]string{},
		time.Now(),
		accounts[r.Intn(len(accounts))].Address,
	)

	if r.Intn(101) <= 50 {
		post = post.WithMedias(simulation.RandomMedias(r, accounts))
	}

	if r.Intn(101) <= 50 {
		if pollData := simulation.RandomPollData(r); pollData != nil {
			post = post.WithPollData(*pollData)
		}
	}

	return post
}

//RandomQueryParams returns randomized QueryPostsParams
func RandomQueryParams(r *rand.Rand) types.QueryPostsParams {
	sortBy := types.PostSortByCreationDate
	sortOrder := types.PostSortOrderAscending
	allowsComments := r.Intn(101) <= 50

	if r.Intn(101) <= 50 {
		sortBy = types.PostSortByID
	}

	if r.Intn(101) <= 50 {
		sortOrder = types.PostSortOrderDescending
	}

	return types.QueryPostsParams{
		Page:           r.Intn(10),
		Limit:          r.Intn(100),
		SortBy:         sortBy,
		SortOrder:      sortOrder,
		ParentID:       nil,
		CreationTime:   nil,
		AllowsComments: &allowsComments,
		Subspace:       "",
		Creator:        nil,
		Hashtags:       nil,
	}
}

func BenchmarkKeeper_SavePost(b *testing.B) {
	fmt.Println("Benchmark: Save a post")
	ctx, k := SetupTestInput()
	post := RandomPost()

	b.SetParallelism(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, post)
	}
}

func BenchmarkKeeper_GetPost(b *testing.B) {
	fmt.Println("Benchmark: Get a post")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	randomPost := posts[r.Intn(len(posts))]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetPost(ctx, randomPost.PostID)
	}

}

func BenchmarkKeeper_GetPostsFiltered(b *testing.B) {
	fmt.Println("Benchmark: Get posts filtered")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	randomQueryParams := RandomQueryParams(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = k.GetPostsFiltered(ctx, randomQueryParams)
	}
}

func BenchmarkKeeper_SavePostReaction(b *testing.B) {
	fmt.Println("Benchmark Save a post reaction")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		k.SavePostReaction(ctx, post.PostID, reaction)
	}
}

func BenchmarkKeeper_GetPostReactions(b *testing.B) {
	fmt.Println("Benchmark Get a post reaction")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		k.SavePostReaction(ctx, post.PostID, reaction)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetPostReactions(ctx, post.PostID)
	}
}
