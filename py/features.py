# EvilInsultGenerator SDK feature factory

from feature.base_feature import EvilInsultGeneratorBaseFeature
from feature.test_feature import EvilInsultGeneratorTestFeature


def _make_feature(name):
    features = {
        "base": lambda: EvilInsultGeneratorBaseFeature(),
        "test": lambda: EvilInsultGeneratorTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
