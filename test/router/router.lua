local router = {
  ['/http/redirect'] = {
    get = 'http/redirect.lua'
  },
  ['/http/write'] = {
    get = 'http/write.lua'
  },
  ['/http/request_method'] = {
    get = 'http/request_method.lua'
  },
  ['/http/uri'] = {
    get = 'http/uri.lua'
  },
  ['/http/param/:name'] = {
    get = 'http/param.lua'
  },
  ['/http/serve_file'] = {
    get = 'http/serve_file.lua'
  },
  ['/http/set_cookie'] = {
    get = 'http/set_cookie.lua'
  },
  ['/http/get_cookie'] = {
    get = 'http/get_cookie.lua'
  },
  ['/http/remote_address'] = {
    get = 'http/remote_address.lua'
  },
  ['/config/get'] = {
    get = 'config/get.lua'
  },
  ['/template/render'] = {
    get = 'template/render.lua'
  },
  ['/url/query_escape'] = {
    get = 'url/query_escape.lua'
  },
  ['/url/query_unescape'] = {
    get = 'url/query_unescape.lua'
  },
  ['/url/path_escape'] = {
    get = 'url/path_escape.lua'
  },
  ['/url/path_unescape'] = {
    get = 'url/path_unescape.lua'
  },
  ['/cache/set'] = {
    get = 'cache/set.lua'
  },
  ['/cache/get'] = {
    get = 'cache/get.lua'
  }
}

return router