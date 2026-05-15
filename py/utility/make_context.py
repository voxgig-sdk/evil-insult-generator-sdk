# EvilInsultGenerator SDK utility: make_context

from core.context import EvilInsultGeneratorContext


def make_context_util(ctxmap, basectx):
    return EvilInsultGeneratorContext(ctxmap, basectx)
