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
  }
}

return router