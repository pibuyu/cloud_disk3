// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadReply struct {
	Ext      string `json:"ext"`
	Name     string `json:"name"`
	Message  string `json:"message"`
	Identity string `json:"identity"`
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"` //文件后缀名
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"` //文件在cos服务器上的路径
}

type LoginReply struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type SendCodeReply struct {
	Message string `json:"message"`
}

type SendCodeRequest struct {
	Email string `json:"email"`
}

type UserDeatilReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDeatilRequest struct {
	Identity string `json:"identity"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	Name               string `json:"name"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int    `json:"size"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list",optional`
	Count int         `json:"count",optional`
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"` //文件夹id
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserRegisterReply struct {
	Message string `json:"message"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRepositorySaveReply struct {
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}
