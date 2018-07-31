local req = require('http')
local cfg = require('config')
req.write(cfg.get('myTestKey'))