// Setup ////////////////////////////////////////
const PLAYLISTS = new Set()
const PLAYLISTS_CONTAINER = document.getElementById('my-playlists')

getPlaylists()



function getPlaylists() {

  // Set up request to create playlist
  let request = new XMLHttpRequest()
  request.responseType = 'json'

  request.open("GET", "/get-playlists", true)
  request.send()



  // Wait for playlist response
  request.onload = function () {

    // Iterate through response
    for (let ids of request.response.playlists) {
      
      // Test if playlists already fetched
      const idHash = ids.join('|')
      if (PLAYLISTS.has(idHash)) continue
      PLAYLISTS.add(idHash)


      // Add playlists to playlist container
      for (let id of ids) {
        playlist = document.createElement('iframe')
        playlist.classList.add('playlist')
        playlist.src = 'https://open.spotify.com/embed/playlist/' + id + '?utm_source=generator'

        PLAYLISTS_CONTAINER.appendChild(playlist)
      }
    }
  }
}