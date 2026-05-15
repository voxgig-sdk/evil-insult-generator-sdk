# EvilInsultGenerator SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module EvilInsultGeneratorFeatures
  def self.make_feature(name)
    case name
    when "base"
      EvilInsultGeneratorBaseFeature.new
    when "test"
      EvilInsultGeneratorTestFeature.new
    else
      EvilInsultGeneratorBaseFeature.new
    end
  end
end
