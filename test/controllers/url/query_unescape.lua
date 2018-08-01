local req = require('http')
local url = require('url')
req.write(url.queryUnescape('Testing+bison+URL+module'))