import { ENV, SUPABASE_ANON_KEY, SUPABASE_SERVICE_KEY, SUPABASE_URL } from "$env/static/private";
import { redirect, type Cookies } from "@sveltejs/kit";
import { createServerClient } from "@supabase/ssr";

import type { Database } from "../types.db";
import Tokens from "./tokens";


async function Supabase(cookies: Cookies, authorizationCode?: string | null) {
  const supabase = createServerClient<Database>(SUPABASE_URL, SUPABASE_ANON_KEY, {
    cookies: {
      get: (key) => cookies.get(key),
      set: (key, value, options) => cookies.set(key, value, { ...options, secure: ENV !== 'DEV' }),
      remove: (key, options) => cookies.delete(key, { ...options, secure: ENV !== 'DEV' })
    },
    db: { schema: 'autofy' }
  })

  if (authorizationCode) {
    await supabase.auth.exchangeCodeForSession(authorizationCode)
    const { data: { session } } = await supabase.auth.getSession()
    await Tokens.save(session?.provider_token, session?.provider_refresh_token, cookies)
    throw redirect(303, '/playlists');
  }

  return supabase
}

export function Service() {
  return createServerClient<Database>(SUPABASE_URL, SUPABASE_SERVICE_KEY, {
    cookies: {
      get: undefined,
      set: undefined,
      remove: undefined
    }
  })
}

export default Supabase
