# @Author: Junebao
# @Time    : 2020/9/12 20:04
# @File    : code_tool.py

import re
import sys

Public = 1
Private = 2


def f_text(text: str, obj) -> str:
    flags = [flag.replace("%", "") for flag in re.findall(r"%[a-z][a-zA-Z]*%", text)]
    flags_tuple = tuple(flags)
    assert len(flags_tuple) == len(flags), "【Error】 Parameters with duplicate names"
    for flag in flags:
        v = getattr(obj, flag)
        text = text.replace(f"%{flag}%", v, 1)
    return text


class ToolException(Exception):
    pass


class CodeTool:
    def __init__(self, path: str):
        self._aip_path = path
        routes = self._aip_path.split("/")
        if len(routes) < 2 or routes[0] != "api":
            raise ToolException("api format does not meet the specification ！")
        self.routes = routes

    @staticmethod
    def get_func_name(routes: [str], name_type: int, suffix: str) -> str:
        prefix = routes[1]
        if name_type == Public:
            prefix = prefix.title()
        for r in routes[2:]:
            prefix += r.title()
        return prefix + suffix

    def _get_route_func_name(self):
        self.route = self.get_func_name(self.routes, Private, "Routes")

    def _get_check_func_name(self):
        self.check = self.get_func_name(self.routes, Public, "Check")

    def _get_server_func_name(self):
        self.server = self.get_func_name(self.routes, Public, "Logic")

    def _get_req_name(self):
        self.req = self.get_func_name(self.routes, Public, "Req")

    def _get_resp_name(self):
        self.resp = self.get_func_name(self.routes, Public, "Resp")

    def _set_route(self):
        func = """
func %route%() []gin.HandlerFunc {
    return []gin.HandlerFunc{
        junebao_top.EasyHandler(check.%check%,
            server.%server%, message.%req%{}),
    }
}"""
        func = f_text(func, self)
        with open(f"./src/routes/{self.routes[1]}.go", "a+") as fp:
            fp.write(func)

    def _set_server(self):
        func = """
func %server%(ctx *gin.Context, req junebaotop.BaseReqInter) junebaotop.BaseRespInter {
    request := req.(*message.%req%)
    resp := message.%resp%{}
    // TODO:...
    log.Println(request)
    resp.Header = junebaotop.SuccessRespHeader
    return resp
}"""
        func = f_text(func, self)
        with open(f"./src/server/{self.routes[1]}.go", "a+") as fp:
            fp.write(func)

    def _set_check(self):
        func = """
func %check%(ctx *gin.Context, req junebao_top.BaseReqInter) (junebao_top.BaseRespInter, error) {
    request := req.(*message.%req%)
    //TODO:...
    return http.StatusOK, nil
}"""
        func = f_text(func, self)
        with open(f"./src/check/{self.routes[1]}.go", "a+") as fp:
            fp.write(func)

    def _set_message(self):
        func = """
type %resp% struct {
    Header junebao_top.BaseRespHeader `json:"header"`
}

type %req% struct {
}

func (r %req%) JSON(ctx *gin.Context) error {
    return ctx.ShouldBindJSON(&r)
}"""
        func = f_text(func, self)
        with open(f"./src/message/{self.routes[1]}.go", "a+") as fp:
            fp.write(func)

    # api/server/list
    def do(self):
        self._get_route_func_name()
        self._get_check_func_name()
        self._get_server_func_name()
        self._get_req_name()
        self._get_resp_name()
        self._set_route()
        self._set_server()
        self._set_check()
        self._set_message()


if __name__ == '__main__':
    path = sys.argv[1]
    ct = CodeTool(path)
    ct.do()

