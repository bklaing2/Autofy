// Setup ////////////////////////////////////////
const PLAYLISTS = new Set()
const PLAYLISTS_CONTAINER = document.getElementById('my-playlists')

getPlaylists()




// Functions ////////////////////////////////////
function getPlaylists() {
  console.log('Getting playlists...')

  // Set up request to create playlist
  let request = new XMLHttpRequest()
  request.responseType = 'json'

  request.open("GET", "/get-playlists", true)
  request.send()



  // Wait for playlist response
  request.onload = function () {
    let hasGenerating = false

    // Iterate through response
    for (let playlist of request.response.playlists) {

      // Get parent element, or create one defaulting to loading icon
      let parent_element = document.getElementById(playlist.id)
      if (parent_element === null) {
        parent_element = document.createElement('div')
        parent_element.id = playlist.id
        parent_element.classList.add('playlist')
        parent_element.classList.add('loading')

        PLAYLISTS_CONTAINER.appendChild(parent_element)
      }


      // Skip this loop iteration if playlist is still generating
      if (playlist.spotifyPlaylists === 'generating') {
        hasGenerating = true
        continue
      }


      // Skip this loop iteration if playlist already fetched
      if (PLAYLISTS.has(playlist.id)) continue
      PLAYLISTS.add(playlist.id)

      parent_element.classList.remove('loading')


      
      // Add playlists to playlist container
      for (let id of playlist.spotifyPlaylists) {
        playlist_element = document.createElement('iframe')
        playlist_element.id = id
        playlist_element.classList.add('playlist')
        playlist_element.src = 'https://open.spotify.com/embed/playlist/' + id + '?utm_source=generator'

        parent_element.appendChild(playlist_element)
      }
    }



    // If any playlists are still generating, call this function again in 10 seconds
    if (hasGenerating) setTimeout(getPlaylists, 10000)
  }
}




function addPlaylist(requestData) {
  // Set up request to create playlist
  request = new XMLHttpRequest()
  request.responseType = 'json'
  
  request.open("POST", "/create-playlist", true)
  request.send(requestData)

  console.log(requestData)


  // Wait for playlist response
  request.onload = function () {
    getPlaylists()
  }
}