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
  }
}

return router