import type { Actions } from './$types';
import Tokens from '$lib/server/tokens.js'
import { redirect } from '@sveltejs/kit'


export const actions = {
  default: async ({ cookies }) => {
    Tokens.clear(cookies)
    throw redirect(303, '/')
  }
} satisfies Actions
