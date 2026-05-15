-- EvilInsultGenerator SDK error

local EvilInsultGeneratorError = {}
EvilInsultGeneratorError.__index = EvilInsultGeneratorError


function EvilInsultGeneratorError.new(code, msg, ctx)
  local self = setmetatable({}, EvilInsultGeneratorError)
  self.is_sdk_error = true
  self.sdk = "EvilInsultGenerator"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function EvilInsultGeneratorError:error()
  return self.msg
end


function EvilInsultGeneratorError:__tostring()
  return self.msg
end


return EvilInsultGeneratorError
