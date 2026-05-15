
import { Context } from './Context'


class EvilInsultGeneratorError extends Error {

  isEvilInsultGeneratorError = true

  sdk = 'EvilInsultGenerator'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  EvilInsultGeneratorError
}

