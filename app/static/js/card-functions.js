function expandCard(event, cardId) {
  event.stopPropagation()

  card = document.getElementById(cardId)
  card.classList.add('expanded')
  card.classList.remove('button')

  console.log('expand')
}

function closeCard(event, cardId) {
  event.stopPropagation()
  card = document.getElementById(cardId)
  card.classList.add('button')
  card.classList.remove('expanded')

  console.log('close')
}



function submitForm(event, cardId) {
  event.preventDefault()
  closeCard(event, cardId)

  form = document.getElementById(cardId).children[1]
  addPlaylist(new FormData(form))
}


function addPlaylist(requestData) {
  // Create new playlist element
  playlist = document.createElement('iframe')
  playlist.classList.add('playlist')
  playlist.classList.add('loading')

  playlists = document.getElementById('my-playlists').children[1]
  playlists.appendChild(playlist)


  // Set up request to create playlist
  var request = new XMLHttpRequest();
  request.responseType = 'json'
  
  request.open("POST", "/create-playlist", true);
  request.send(requestData);


  // Wait for playlist response
  request.onload = function () {
    playlistId = request.response.playlist_id
    playlist.src = 'https://open.spotify.com/embed/playlist/' + playlistId + '?utm_source=generator'
    playlist.classList.remove('loading')
  };
}