# Web 框架

Web 框架 接受 HTTP Request，返回 HTTP Response 给 Web Server。

它主要处理 HTTP 中一些细碎而繁琐的部分，例如：

- 设置动态路由，将请求转发给对应的 Controller.
- 维护请求的上下文信息，做到日志、鉴权等功能。
- 分组控制，插入中间件，维护中间件的信息。
- 渲染模板（在前后端分离的现在越来越少见)
