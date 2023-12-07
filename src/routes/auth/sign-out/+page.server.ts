import Tokens from '$lib/server/tokens.js'
import { redirect } from '@sveltejs/kit'


export const actions = {
	default: async ({ cookies, locals }) => {
    const { supabase } = locals
    await supabase.auth.signOut()
    Tokens.clear(cookies)
    throw redirect(303, '/')
  }
}
