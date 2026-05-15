-- ProjectName SDK exists test

local sdk = require("evil-insult-generator_sdk")

describe("EvilInsultGeneratorSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
