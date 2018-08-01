local cache = require('cache')
local req = require('http')
req.write(cache.get("bison-test"))