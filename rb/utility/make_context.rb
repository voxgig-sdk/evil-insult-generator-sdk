# EvilInsultGenerator SDK utility: make_context
require_relative '../core/context'
module EvilInsultGeneratorUtilities
  MakeContext = ->(ctxmap, basectx) {
    EvilInsultGeneratorContext.new(ctxmap, basectx)
  }
end
