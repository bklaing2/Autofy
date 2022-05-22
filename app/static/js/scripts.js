timeSinceLastSearch = 0

// Adding artists ///////////////////////////////
function addArtist(button) {
  //  <div name="artist" class="row">
  //    <div class="bg-img">
  //      <button type="button" class="delete-button" onclick="deleteArtist(this)">X</button>
  //    </div>
  //
  //    <div class="search-artist">
  //      <input
  //          type="text"
  //          placeholder="Search..."
  //          onkeydown="onKeyDown(this, event)"
  //          onfocus="spawnPrediction(this)"
  //          onblur="destroyPrediction(this)"
  //      >
  //      <input type="hidden" name="artists" value="123">
  //    </div>
  //  </div>

  row = document.createElement('div')
  row.setAttribute('name', 'artist')
  row.classList.add('row')

  bgImg = document.createElement('div')
  bgImg.classList.add('bg-img')
  row.appendChild(bgImg)

  deleteButton = document.createElement('button')
  deleteButton.setAttribute('type', 'button')
  deleteButton.setAttribute('onclick', 'deleteArtist(this)')
  deleteButton.classList.add('delete-button')
  deleteButton.innerHTML = 'X'
  bgImg.appendChild(deleteButton)

  searchArtist = document.createElement('div')
  searchArtist.classList.add('search-artist')
  row.appendChild(searchArtist)

  input = document.createElement('input')
  input.setAttribute('type', 'text')
  input.setAttribute('placeholder', 'Search...')
  input.setAttribute('onkeydown', 'onKeyDown(this, event)')
  input.setAttribute('onfocus', 'spawnPrediction(this)')
  input.setAttribute('onblur', 'destroyPrediction(this)')
  searchArtist.appendChild(input)

  id = document.createElement('input')
  id.setAttribute('type', 'hidden')
  id.setAttribute('name', 'artists')
  searchArtist.appendChild(id)


  parent = document.getElementById('choose-artists')
  parent.insertBefore(row, button)

  input.focus()
}

function deleteArtist(button) {
  artist = button.parentElement.parentElement
  container = artist.parentElement

  container.removeChild(artist)
}


function onKeyDown(input, e) {
  if (e.key == 'Tab' || e.key == 'ArrowRight') {
    acceptPrediction(input)
  } else {
    updatePrediction(input)
  }
}

function spawnPrediction(input) {
  parent = input.parentElement
  span = document.createElement('span')

  span.classList.add('prediction')
  span.classList.add('secondary')
  span.innerHTML = input.value

  parent.insertBefore(span, input)
}

function destroyPrediction(input) {
  acceptPrediction(input)

  parent = input.parentElement
  span = input.previousElementSibling

  parent.removeChild(span)
}

function updatePrediction(input) {


  prediction = input.previousElementSibling
  img = input.parentElement.previousElementSibling
  id = input.nextElementSibling

  query = input.value
  if (query === '') {
    prediction.innerHTML = ''
    img.style.backgroundImage = 'none'
    id.value = ''
    return
  }

  prediction.innerHTML = query + prediction.innerHTML.substring(query.length)


  if (Date.now() - timeSinceLastSearch < 500) return
  timeSinceLastSearch = Date.now()

  // Search for artist
  request = new XMLHttpRequest()
  request.open("GET", "/search-for-artist?artist=" + query, true)
  request.responseType = 'json'

  request.send()

  request.onload = function () {
    artist = request.response

    // If there is a matching artist
    if (Object.keys(artist).length > 0) {

      // Update prediction text, artist image, and hidden id
      prediction.innerHTML = query + artist.name.substring(query.length)
      img.style.backgroundImage = 'url(' + artist.profilePicture + ')'
      id.value = artist.id
     }

     // If there isn't a matching artist
     else {
      prediction.innerHTML = ''
      img.style.backgroundImage = 'none'
      id.value = ''
     }
  }
}

function acceptPrediction(input) {
  prediction = input.previousElementSibling

  if (prediction.innerHTML != '') {
    input.value = prediction.innerHTML
  }
}


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
    playlistId = request.response.playlist_id
    playlist.src = 'https://open.spotify.com/embed/playlist/' + playlistId + '?utm_source=generator'
    playlist.classList.remove('loading')
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