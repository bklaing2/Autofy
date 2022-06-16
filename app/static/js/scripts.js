// Adding playlist //////////////////////////////
function submitForm(event) {
  event.preventDefault()

  form = document.getElementById('create-playlist-form')
  addPlaylist(new FormData(form))

  hideOverlay()
}


function addPlaylist(requestData) {
  // Create new playlist element
  playlist = document.createElement('iframe')
  playlist.classList.add('playlist')
  playlist.classList.add('loading')

  playlists = document.getElementById('my-playlists')
  playlists.appendChild(playlist)


  // Set up request to create playlist
  request = new XMLHttpRequest()
  request.responseType = 'json'
  
  request.open("POST", "/create-playlist", true)
  request.send(requestData)

  console.log(requestData)


  // Wait for playlist response
  request.onload = function () {
    playlistIds = request.response.playlistIds

    for (var i = 0; i < playlistIds.length; i++) {
      if (i > 0) {
        playlist = document.createElement('iframe')
        playlist.classList.add('playlist')
        playlist.classList.add('loading')

        playlists = document.getElementById('my-playlists')
        playlists.appendChild(playlist)
      }

      playlist.src = 'https://open.spotify.com/embed/playlist/' + playlistIds[i] + '?utm_source=generator'
      playlist.classList.remove('loading')
    }
  }
}


function searchArtists(input) {

  request = new XMLHttpRequest()
  request.open("GET", "/search-artists?artist=" + input.value, true)
  request.responseType = 'json'

  request.send()

  request.onload = function () {
    datalist = document.getElementById('artists')
    datalist.innerHTML = ''

    request.response.artists.forEach(function(artist) {
      option = document.createElement('option')
      option.setAttribute('value', artist.name)
      option.setAttribute('data-id', artist.id)
      option.setAttribute('data-pic', artist.profile-picture)
      datalist.appendChild(option)
    })
  }
}



// Showing/hiding elements //////////////////////
function showOverlay() {
  overlay = document.getElementById('create-playlist-overlay')
  overlay.removeAttribute('hidden')
}

function hideOverlay() {
  overlay = document.getElementById('create-playlist-overlay')
  overlay.setAttribute('hidden', '')

  hideAdvancedOptions()
}


function toggleAdvancedOptions(event) {
  event.stopPropagation()

  if (document.getElementById('advanced-options').hasAttribute('hidden')) {
    showAdvancedOptions()
  } else {
    hideAdvancedOptions()
  }
}

function showAdvancedOptions() {
  options = document.getElementById('advanced-options')
  button = document.getElementById('advanced-options-button')

  options.removeAttribute('hidden')
  button.innerHTML = '- advanced'
}

function hideAdvancedOptions() {
  options = document.getElementById('advanced-options')
  button = document.getElementById('advanced-options-button')

  options.setAttribute('hidden', '')

  button = document.getElementById('advanced-options-button')
  button.innerHTML = '+ advanced'

  // Delete all artists
  artists = document.getElementsByName('artist')
  artistsParent = artists[0].parentElement
  artists.forEach(artist => {
    console.log(artist)
    artistsParent.remove(artist)
  })

  form = document.getElementById('create-playlist-form')
  form.reset()
}