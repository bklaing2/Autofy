// Adding playlist //////////////////////////////
function submitForm(event) {
  event.preventDefault()

  form = document.getElementById('create-playlist-form')
  addPlaylist(new FormData(form))

  hideOverlay()
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