package service

import (
	v1 "Kratos-realworld/api/realworld/v1"
	"context"
)

func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentRequest) (reply *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	return &v1.TagListReply{}, nil
}

func (s *RealWorldService) FavouriteArticle(ctx context.Context, req *v1.FavouriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UnfavouriteArticle(ctx context.Context, req *v1.FavouriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
