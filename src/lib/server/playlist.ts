import type { Database } from "$lib/types";
import type SpotifyWebApi from "spotify-web-api-node";


type Playlist = Database['public']['Tables']['playlists']['Row']
type IdValue = { id: string; value: string; }


async function populate(playlist: Playlist['id'], artists: string[], spotify: SpotifyWebApi) {
  const tracks = await getArtistsTracks(artists, spotify)
  await addTracksToPlaylist(playlist, tracks, spotify)
  console.log('Finished')
}


async function update(playlist: Playlist, spotify: SpotifyWebApi) {
  let artistsToAdd: string[] = []
  let artistsToRemove: string[] = []

  const oldFollowed = playlist.followed_artists ?? []
  const newFollowed = playlist.followed_artists ? await getFollowedArtists(spotify) : []

  const allOldArtists = oldFollowed.concat(playlist.artists)
  const allNewArtists = newFollowed.concat(playlist.artists)

  // Handle followed artists
  if (playlist.followed_artists && (playlist.update_when_user_follows_artist || playlist.update_when_user_unfollows_artist)) {
    if (playlist.update_when_user_follows_artist)
      artistsToAdd = [ ...new Set(artistsToAdd.concat(newFollowed.filter(a => !allOldArtists.includes(a)))) ]

    if (playlist.update_when_user_unfollows_artist)
      artistsToRemove = [ ...new Set(artistsToRemove.concat(oldFollowed.filter(a => !allNewArtists.includes(a)))) ]
  }


  // Handle tracks released since playlist was updated
  let newTracks: string[] = []
  if (playlist.update_when_artist_posts && Day(playlist.updated_at) !== Day()) {
    let albums: IdValue[] = []
    
    const artistsIntersection = oldFollowed.filter(a => newFollowed.includes(a)).concat(playlist.artists)
    for (const artist of new Set(artistsIntersection))
      albums = albums.concat(await getArtistAlbumsWithReleaseDate(artist, spotify))
    
    albums = albums.filter(a => Day(a.value) + 1 >= Day(playlist.updated_at))
    for (const album of albums)
      newTracks = newTracks.concat(await getAlbumTracks(album.id, spotify))
  }


  // Add/remove tracks to/from playlist
  const tracksToAdd: string[] = [ ...await getArtistsTracks(artistsToAdd, spotify), ...newTracks ]
  const tracksToRemove: string[] = await getArtistsTracks(artistsToRemove, spotify)

  await removeTracksFromPlaylist(playlist.id, tracksToRemove, spotify)
  await addTracksToPlaylist(playlist.id, tracksToAdd, spotify)


  // Return artists
  const artists = new Set(playlist.artists.concat(artistsToAdd))
  for (const elem of artistsToRemove) artists.delete(elem)
}


async function updateSettings(oldPlaylist: Playlist | null, newPlaylist: Playlist | null, spotify: SpotifyWebApi) {
  if (!oldPlaylist || !newPlaylist) return newPlaylist

  let artistsToAdd: string[] = []
  let artistsToRemove: string[] = []
  
  const allOldArtists = oldPlaylist.artists.concat(oldPlaylist.followed_artists ?? [])
  const allNewArtists = newPlaylist.artists.concat(newPlaylist.followed_artists ?? [])
  
  // Handle artists added/removed
  const oldFollowed = oldPlaylist.followed_artists
  const newFollowed = newPlaylist.followed_artists

  if (!!oldFollowed !== !!newFollowed) {
    if (newFollowed) artistsToAdd = newFollowed.filter(a => !allOldArtists.includes(a)) ?? []
    if (oldFollowed) artistsToRemove = oldFollowed.filter(a => allOldArtists.includes(a)) ?? []
  }

  artistsToAdd = artistsToAdd.concat(newPlaylist.artists.filter(a => !allOldArtists.includes(a)))
  artistsToRemove = artistsToRemove.concat(oldPlaylist.artists.filter(a => !allNewArtists.includes(a)))
  artistsToAdd = [ ...new Set(artistsToAdd) ]
  artistsToRemove = [ ...new Set(artistsToRemove) ]

  const tracksToAdd = await getArtistsTracks(artistsToAdd, spotify)
  const tracksToRemove = await getArtistsTracks(artistsToRemove, spotify)

  await removeTracksFromPlaylist(oldPlaylist.id, tracksToRemove, spotify)
  await addTracksToPlaylist(oldPlaylist.id, tracksToAdd, spotify)
}



async function getFollowedArtists (spotify: SpotifyWebApi) {
  let artists: string[] = []
    let after: string | undefined

    do {
      const data = await spotify.getFollowedArtists({ limit: 50, after: after })
      artists = artists.concat(data.body.artists.items.map(a => a.id))
      after = data.body.artists.cursors.after
    } while (after)

    return artists
}


// TODO:: speed it up by running in parallel
async function getArtistAlbums (artist: string, spotify: SpotifyWebApi) {
  let albums: string[] = []

  let hasNext = true
  for (let i = 0; hasNext; i++) {
    const data = await spotify.getArtistAlbums(artist, {
      offset: i * 50,
      limit: 50,
      include_groups: 'album',
      country: 'US'
    })

    albums = albums.concat(data.body.items.map(a => a.id))
    hasNext = !!data.body.next
  }

  return albums
}


// TODO:: speed it up by running in parallel
async function getArtistAlbumsWithReleaseDate (artist: string, spotify: SpotifyWebApi) {

  let albums: IdValue[] = []

  let hasNext = true
  for (let i = 0; hasNext; i++) {
    const data = await spotify.getArtistAlbums(artist, {
      offset: i * 50,
      limit: 50,
      include_groups: 'album',
      country: 'US'
    })

    albums = albums.concat(data.body.items.map(a => ({ id: a.id, value: a.release_date })))
    hasNext = !!data.body.next
  }

  return albums
}


// TODO:: speed it up by running in parallel
async function getAlbumTracks (album: string, spotify: SpotifyWebApi) {

  let tracks: string[] = []

  let hasNext = true
  for (let i = 0; hasNext; i++) {
    const data = await spotify.getAlbumTracks(album, { offset: i * 50, limit: 50 })
    tracks = tracks.concat(data.body.items.map(a => a.uri))
    hasNext = !!data.body.next
  }

  return tracks
}


async function getArtistsTracks (artists: string[], spotify: SpotifyWebApi) {
  let albums: string[] = []
  for (const artist of artists) albums = albums.concat(await getArtistAlbums(artist, spotify))
  console.log(`Albums: ${albums.length}`)

  let tracks: string[] = []
  for (const album of albums) tracks = tracks.concat(await getAlbumTracks(album, spotify))
  console.log(`Tracks: ${tracks.length}`)

  return tracks
}


async function addTracksToPlaylist (playlist: Playlist['id'], tracks: string[], spotify: SpotifyWebApi) {
  for (let i = 0; i < tracks.length; i += 100)
    await spotify.addTracksToPlaylist(playlist, tracks.slice(i, i+100))
}

async function removeTracksFromPlaylist (playlist: Playlist['id'], tracks: string[], spotify: SpotifyWebApi) {
  for (let i = 0; i < tracks.length; i += 100)
    await spotify.removeTracksFromPlaylist(playlist, tracks.slice(i, i+100).map(t => ({ uri: t })))
}




async function getArtists(artistIds: string[], spotify: SpotifyWebApi) {
  let artists: { id: string, name: string, img?: string }[] = []

  for (let i = 0; i < artistIds.length; i += 50) {
    const { body } = await spotify.getArtists(artistIds.slice(i, i+50))
    artists = artists.concat(body.artists.map(a => ({ id: a.id, name: a.name, img: a.images[0]?.url })))
  }

  return artists
}



function Day (date?: string) {
  return date ? new Date(date).setHours(0, 0, 0, 0) : new Date().setHours(0, 0, 0, 0)
}

export default {
  populate,
  update,
  updateSettings,

  getArtists,
  getFollowedArtists
}