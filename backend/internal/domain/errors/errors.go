package domain

import "errors"

// -- repositoryでのエラー --

var ErrNotFound = errors.New("entity not found")

// -- usecaseでのエラー -- 

// 入力データが無効（形式違い、必須項目不足など)
var ErrInvalidInput = errors.New("invalid input data")

// 登録しようとしたメールアドレスがすでにある
var ErrEmailAlreadyExists = errors.New("email already exists")

// ログイン時の認証情報が正しくない
var ErrAuthenticationFailed = errors.New("authentication failed")

// ユーザーがその操作を行う権限がない
var ErrPermissionDenied = errors.New("permission denied")