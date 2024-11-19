import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ locals }) => {
  const { spotify, signedIn } = locals
  const user = signedIn ? (await spotify.getMe()).body : undefined

  return { user }
}
