// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infinitechatgo/service/user/api/internal/logic/group"
	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"
)

func GetUserGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserGroupsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGetUserGroupsLogic(r.Context(), svcCtx)
		err := l.GetUserGroups(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
