# EvilInsultGenerator SDK exists test

require "minitest/autorun"
require_relative "../EvilInsultGenerator_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = EvilInsultGeneratorSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
