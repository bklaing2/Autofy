export async function load({ locals }) {
  const { spotify, signedIn } = locals
  return { user: signedIn ? (await spotify.getMe()).body : undefined }
}