// Adding playlist //////////////////////////////
function submitForm(event) {
  event.preventDefault()

  form = document.getElementById('create-playlist-form')
  addPlaylist(new FormData(form))

  hideOverlay('create-playlist-overlay')
}



// Showing/hiding elements //////////////////////
function showOverlay(id) {
  overlay = document.getElementById(id)
  overlay.removeAttribute('hidden')
}

function hideOverlay(id) {
  overlay = document.getElementById(id)
  overlay.setAttribute('hidden', '')

  // hideAdvancedOptions()
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