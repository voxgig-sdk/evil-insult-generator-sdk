package = "voxgig-sdk-evil-insult-generator"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/evil-insult-generator-sdk.git"
}
description = {
  summary = "EvilInsultGenerator SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["evil-insult-generator_sdk"] = "evil-insult-generator_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
