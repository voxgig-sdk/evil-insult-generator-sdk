# EvilInsultGenerator SDK utility: feature_add
module EvilInsultGeneratorUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
