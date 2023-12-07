export async function load({ locals }) {
  const { supabase, spotify, signedIn } = locals

  if (!signedIn) return { playlists: [] as string[] }

  const { data } = await supabase
    .from('playlists')
    .select('id')
  
  let playlists = data?.map(p => p.id) ?? []
  if (playlists.length === 0) return { playlists: playlists }

  const exists = await spotify.getUserPlaylists()
  for (let i = playlists.length - 1; i >= 0; i--) {
    const playlist = playlists[i]
    if (exists.body.items.find(p => p.id === playlist)) continue

    supabase
      .from('playlists')
      .delete()
      .eq('id', playlist)
    
    playlists = [ ...playlists.slice(0, i), ...playlists.slice(i+1) ]
  }

  return { playlists: playlists }
}