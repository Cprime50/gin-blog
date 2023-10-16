// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aliml92/realworld-gin-sqlc/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/Cprime50/gin-blog/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockStore) AddComment(arg0 context.Context, arg1 db.AddCommentParams) (*db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", arg0, arg1)
	ret0, _ := ret[0].(*db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockStoreMockRecorder) AddComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockStore)(nil).AddComment), arg0, arg1)
}

// CountArticles mocks base method.
func (m *MockStore) CountArticles(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticles", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticles indicates an expected call of CountArticles.
func (mr *MockStoreMockRecorder) CountArticles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticles", reflect.TypeOf((*MockStore)(nil).CountArticles), arg0)
}

// CountArticlesByAuthor mocks base method.
func (m *MockStore) CountArticlesByAuthor(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticlesByAuthor", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticlesByAuthor indicates an expected call of CountArticlesByAuthor.
func (mr *MockStoreMockRecorder) CountArticlesByAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticlesByAuthor", reflect.TypeOf((*MockStore)(nil).CountArticlesByAuthor), arg0, arg1)
}

// CountArticlesByFavorited mocks base method.
func (m *MockStore) CountArticlesByFavorited(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticlesByFavorited", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticlesByFavorited indicates an expected call of CountArticlesByFavorited.
func (mr *MockStoreMockRecorder) CountArticlesByFavorited(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticlesByFavorited", reflect.TypeOf((*MockStore)(nil).CountArticlesByFavorited), arg0, arg1)
}

// CountArticlesByTag mocks base method.
func (m *MockStore) CountArticlesByTag(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticlesByTag", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticlesByTag indicates an expected call of CountArticlesByTag.
func (mr *MockStoreMockRecorder) CountArticlesByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticlesByTag", reflect.TypeOf((*MockStore)(nil).CountArticlesByTag), arg0, arg1)
}

// CountArticlesFeed mocks base method.
func (m *MockStore) CountArticlesFeed(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticlesFeed", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticlesFeed indicates an expected call of CountArticlesFeed.
func (mr *MockStoreMockRecorder) CountArticlesFeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticlesFeed", reflect.TypeOf((*MockStore)(nil).CountArticlesFeed), arg0, arg1)
}

// CreateArticle mocks base method.
func (m *MockStore) CreateArticle(arg0 context.Context, arg1 db.CreateArticleParams) (*db.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", arg0, arg1)
	ret0, _ := ret[0].(*db.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockStoreMockRecorder) CreateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockStore)(nil).CreateArticle), arg0, arg1)
}

// CreateArticleTag mocks base method.
func (m *MockStore) CreateArticleTag(arg0 context.Context, arg1 db.CreateArticleTagParams) (*db.ArticleTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticleTag", arg0, arg1)
	ret0, _ := ret[0].(*db.ArticleTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticleTag indicates an expected call of CreateArticleTag.
func (mr *MockStoreMockRecorder) CreateArticleTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticleTag", reflect.TypeOf((*MockStore)(nil).CreateArticleTag), arg0, arg1)
}

// CreateArticleTx mocks base method.
func (m *MockStore) CreateArticleTx(arg0 context.Context, arg1 db.CreateArticleTxParams) (*db.CreateArticleTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticleTx", arg0, arg1)
	ret0, _ := ret[0].(*db.CreateArticleTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticleTx indicates an expected call of CreateArticleTx.
func (mr *MockStoreMockRecorder) CreateArticleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticleTx", reflect.TypeOf((*MockStore)(nil).CreateArticleTx), arg0, arg1)
}

// CreateTag mocks base method.
func (m *MockStore) CreateTag(arg0 context.Context, arg1 db.CreateTagParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockStoreMockRecorder) CreateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockStore)(nil).CreateTag), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteArticle mocks base method.
func (m *MockStore) DeleteArticle(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockStoreMockRecorder) DeleteArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockStore)(nil).DeleteArticle), arg0, arg1)
}

// DeleteArticleTx mocks base method.
func (m *MockStore) DeleteArticleTx(arg0 context.Context, arg1 db.DeleteArticleTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticleTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticleTx indicates an expected call of DeleteArticleTx.
func (mr *MockStoreMockRecorder) DeleteArticleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticleTx", reflect.TypeOf((*MockStore)(nil).DeleteArticleTx), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteCommentTx mocks base method.
func (m *MockStore) DeleteCommentTx(arg0 context.Context, arg1 db.DeleteCommentTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommentTx indicates an expected call of DeleteCommentTx.
func (mr *MockStoreMockRecorder) DeleteCommentTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentTx", reflect.TypeOf((*MockStore)(nil).DeleteCommentTx), arg0, arg1)
}

// DoesFavoriteExist mocks base method.
func (m *MockStore) DoesFavoriteExist(arg0 context.Context, arg1 db.DoesFavoriteExistParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesFavoriteExist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesFavoriteExist indicates an expected call of DoesFavoriteExist.
func (mr *MockStoreMockRecorder) DoesFavoriteExist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesFavoriteExist", reflect.TypeOf((*MockStore)(nil).DoesFavoriteExist), arg0, arg1)
}

// DoesUserExist mocks base method.
func (m *MockStore) DoesUserExist(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesUserExist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesUserExist indicates an expected call of DoesUserExist.
func (mr *MockStoreMockRecorder) DoesUserExist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesUserExist", reflect.TypeOf((*MockStore)(nil).DoesUserExist), arg0, arg1)
}

// FavoriteArticle mocks base method.
func (m *MockStore) FavoriteArticle(arg0 context.Context, arg1 db.FavoriteArticleParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteArticle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteArticle indicates an expected call of FavoriteArticle.
func (mr *MockStoreMockRecorder) FavoriteArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteArticle", reflect.TypeOf((*MockStore)(nil).FavoriteArticle), arg0, arg1)
}

// FavoriteArticleTx mocks base method.
func (m *MockStore) FavoriteArticleTx(arg0 context.Context, arg1 db.FavoriteArticleTxParams) (*db.FavoriteArticleTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteArticleTx", arg0, arg1)
	ret0, _ := ret[0].(*db.FavoriteArticleTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoriteArticleTx indicates an expected call of FavoriteArticleTx.
func (mr *MockStoreMockRecorder) FavoriteArticleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteArticleTx", reflect.TypeOf((*MockStore)(nil).FavoriteArticleTx), arg0, arg1)
}

// FollowUser mocks base method.
func (m *MockStore) FollowUser(arg0 context.Context, arg1 db.FollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FollowUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FollowUser indicates an expected call of FollowUser.
func (mr *MockStoreMockRecorder) FollowUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FollowUser", reflect.TypeOf((*MockStore)(nil).FollowUser), arg0, arg1)
}

// GetArticleAuthorID mocks base method.
func (m *MockStore) GetArticleAuthorID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleAuthorID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleAuthorID indicates an expected call of GetArticleAuthorID.
func (mr *MockStoreMockRecorder) GetArticleAuthorID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleAuthorID", reflect.TypeOf((*MockStore)(nil).GetArticleAuthorID), arg0, arg1)
}

// GetArticleAuthorIDBySlug mocks base method.
func (m *MockStore) GetArticleAuthorIDBySlug(arg0 context.Context, arg1 string) (*db.GetArticleAuthorIDBySlugRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleAuthorIDBySlug", arg0, arg1)
	ret0, _ := ret[0].(*db.GetArticleAuthorIDBySlugRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleAuthorIDBySlug indicates an expected call of GetArticleAuthorIDBySlug.
func (mr *MockStoreMockRecorder) GetArticleAuthorIDBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleAuthorIDBySlug", reflect.TypeOf((*MockStore)(nil).GetArticleAuthorIDBySlug), arg0, arg1)
}

// GetArticleBySlug mocks base method.
func (m *MockStore) GetArticleBySlug(arg0 context.Context, arg1 string) (*db.GetArticleBySlugRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleBySlug", arg0, arg1)
	ret0, _ := ret[0].(*db.GetArticleBySlugRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleBySlug indicates an expected call of GetArticleBySlug.
func (mr *MockStoreMockRecorder) GetArticleBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleBySlug", reflect.TypeOf((*MockStore)(nil).GetArticleBySlug), arg0, arg1)
}

// GetArticleIDBySlug mocks base method.
func (m *MockStore) GetArticleIDBySlug(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleIDBySlug", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleIDBySlug indicates an expected call of GetArticleIDBySlug.
func (mr *MockStoreMockRecorder) GetArticleIDBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleIDBySlug", reflect.TypeOf((*MockStore)(nil).GetArticleIDBySlug), arg0, arg1)
}

// GetArticles mocks base method.
func (m *MockStore) GetArticles(arg0 context.Context, arg1 db.GetArticlesParams) ([]*db.GetArticlesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetArticlesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockStoreMockRecorder) GetArticles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockStore)(nil).GetArticles), arg0, arg1)
}

// GetArticlesByAuthor mocks base method.
func (m *MockStore) GetArticlesByAuthor(arg0 context.Context, arg1 db.GetArticlesByAuthorParams) ([]*db.GetArticlesByAuthorRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByAuthor", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetArticlesByAuthorRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticlesByAuthor indicates an expected call of GetArticlesByAuthor.
func (mr *MockStoreMockRecorder) GetArticlesByAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByAuthor", reflect.TypeOf((*MockStore)(nil).GetArticlesByAuthor), arg0, arg1)
}

// GetArticlesByFavorited mocks base method.
func (m *MockStore) GetArticlesByFavorited(arg0 context.Context, arg1 db.GetArticlesByFavoritedParams) ([]*db.GetArticlesByFavoritedRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByFavorited", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetArticlesByFavoritedRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticlesByFavorited indicates an expected call of GetArticlesByFavorited.
func (mr *MockStoreMockRecorder) GetArticlesByFavorited(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByFavorited", reflect.TypeOf((*MockStore)(nil).GetArticlesByFavorited), arg0, arg1)
}

// GetArticlesByTag mocks base method.
func (m *MockStore) GetArticlesByTag(arg0 context.Context, arg1 db.GetArticlesByTagParams) ([]*db.GetArticlesByTagRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByTag", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetArticlesByTagRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticlesByTag indicates an expected call of GetArticlesByTag.
func (mr *MockStoreMockRecorder) GetArticlesByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByTag", reflect.TypeOf((*MockStore)(nil).GetArticlesByTag), arg0, arg1)
}

// GetArticlesFeed mocks base method.
func (m *MockStore) GetArticlesFeed(arg0 context.Context, arg1 db.GetArticlesFeedParams) ([]*db.GetArticlesFeedRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesFeed", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetArticlesFeedRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticlesFeed indicates an expected call of GetArticlesFeed.
func (mr *MockStoreMockRecorder) GetArticlesFeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesFeed", reflect.TypeOf((*MockStore)(nil).GetArticlesFeed), arg0, arg1)
}

// GetCommentAuthorID mocks base method.
func (m *MockStore) GetCommentAuthorID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentAuthorID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentAuthorID indicates an expected call of GetCommentAuthorID.
func (mr *MockStoreMockRecorder) GetCommentAuthorID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentAuthorID", reflect.TypeOf((*MockStore)(nil).GetCommentAuthorID), arg0, arg1)
}

// GetCommentsBySlug mocks base method.
func (m *MockStore) GetCommentsBySlug(arg0 context.Context, arg1 string) ([]*db.GetCommentsBySlugRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsBySlug", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetCommentsBySlugRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsBySlug indicates an expected call of GetCommentsBySlug.
func (mr *MockStoreMockRecorder) GetCommentsBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsBySlug", reflect.TypeOf((*MockStore)(nil).GetCommentsBySlug), arg0, arg1)
}

// GetFollowees mocks base method.
func (m *MockStore) GetFollowees(arg0 context.Context, arg1 string) ([]*db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowees", arg0, arg1)
	ret0, _ := ret[0].([]*db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowees indicates an expected call of GetFollowees.
func (mr *MockStoreMockRecorder) GetFollowees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowees", reflect.TypeOf((*MockStore)(nil).GetFollowees), arg0, arg1)
}

// GetTags mocks base method.
func (m *MockStore) GetTags(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockStoreMockRecorder) GetTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockStore)(nil).GetTags), arg0)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// IsFollowing mocks base method.
func (m *MockStore) IsFollowing(arg0 context.Context, arg1 db.IsFollowingParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFollowing", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFollowing indicates an expected call of IsFollowing.
func (mr *MockStoreMockRecorder) IsFollowing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFollowing", reflect.TypeOf((*MockStore)(nil).IsFollowing), arg0, arg1)
}

// IsFollowingList mocks base method.
func (m *MockStore) IsFollowingList(arg0 context.Context, arg1 db.IsFollowingListParams) ([]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFollowingList", arg0, arg1)
	ret0, _ := ret[0].([]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFollowingList indicates an expected call of IsFollowingList.
func (mr *MockStoreMockRecorder) IsFollowingList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFollowingList", reflect.TypeOf((*MockStore)(nil).IsFollowingList), arg0, arg1)
}

// UnfavoriteArticle mocks base method.
func (m *MockStore) UnfavoriteArticle(arg0 context.Context, arg1 db.UnfavoriteArticleParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfavoriteArticle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnfavoriteArticle indicates an expected call of UnfavoriteArticle.
func (mr *MockStoreMockRecorder) UnfavoriteArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfavoriteArticle", reflect.TypeOf((*MockStore)(nil).UnfavoriteArticle), arg0, arg1)
}

// UnfavoriteArticleTx mocks base method.
func (m *MockStore) UnfavoriteArticleTx(arg0 context.Context, arg1 db.UnfavoriteArticleTxParams) (*db.UnfavoriteArticleTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfavoriteArticleTx", arg0, arg1)
	ret0, _ := ret[0].(*db.UnfavoriteArticleTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnfavoriteArticleTx indicates an expected call of UnfavoriteArticleTx.
func (mr *MockStoreMockRecorder) UnfavoriteArticleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfavoriteArticleTx", reflect.TypeOf((*MockStore)(nil).UnfavoriteArticleTx), arg0, arg1)
}

// UnfollowUser mocks base method.
func (m *MockStore) UnfollowUser(arg0 context.Context, arg1 db.UnfollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfollowUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnfollowUser indicates an expected call of UnfollowUser.
func (mr *MockStoreMockRecorder) UnfollowUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfollowUser", reflect.TypeOf((*MockStore)(nil).UnfollowUser), arg0, arg1)
}

// UpdateArticle mocks base method.
func (m *MockStore) UpdateArticle(arg0 context.Context, arg1 db.UpdateArticleParams) (*db.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", arg0, arg1)
	ret0, _ := ret[0].(*db.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockStoreMockRecorder) UpdateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockStore)(nil).UpdateArticle), arg0, arg1)
}

// UpdateArticleTx mocks base method.
func (m *MockStore) UpdateArticleTx(arg0 context.Context, arg1 db.UpdateArticleTxParams) (*db.UpdateArticleTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticleTx", arg0, arg1)
	ret0, _ := ret[0].(*db.UpdateArticleTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticleTx indicates an expected call of UpdateArticleTx.
func (mr *MockStoreMockRecorder) UpdateArticleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticleTx", reflect.TypeOf((*MockStore)(nil).UpdateArticleTx), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}