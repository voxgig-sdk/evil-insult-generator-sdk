# ProjectName SDK exists test

import pytest
from evilinsultgenerator_sdk import EvilInsultGeneratorSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = EvilInsultGeneratorSDK.test(None, None)
        assert testsdk is not None
