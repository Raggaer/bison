local req = require('http')
local session = require('session')
req.write(session.get("bison-test"))