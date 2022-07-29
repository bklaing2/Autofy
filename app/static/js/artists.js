const ARTISTS_CONTAINER = document.getElementById('artists')
const ARTISTS_SEARCH_RESULTS_CONTAINER = document.getElementById('search-artists-results')

let handleInput;
function onSearchArtists(input, e) {
  // Only search if user hasn't typed in 0.5s
  clearTimeout(handleInput)
  handleInput = setTimeout(() => { searchArtists(input.value) }, 500)
}



// function onKeyDown(input, e) {
//   if (e.key == 'Tab' || e.key == 'ArrowRight') {
//     acceptPrediction(input)
//   } else {
//     updatePrediction(input)
//   }
// }



// <li class="row" onclick="onArtistSearchResultClick(this)" data-id="TEST-ID-Taylor-Swift" data-name="Taylor Swift" data-img="https://i.scdn.co/image/ab6761610000e5eb9e3acf1eaf3b8846e836f441">
//   <img class="artist-img" src="https://i.scdn.co/image/ab6761610000e5eb9e3acf1eaf3b8846e836f441">
//   <span>Taylor Swift</span>
// </img></li>
function populateSearchResults(artists) {
  console.log(artists)
  // Remove all child nodes
  while (ARTISTS_SEARCH_RESULTS_CONTAINER.firstChild) {
    ARTISTS_SEARCH_RESULTS_CONTAINER.removeChild(ARTISTS_SEARCH_RESULTS_CONTAINER.lastChild);
  }

  // Add new results
  for (let artist of artists) {
    let artistElement = document.createElement('li')
    artistElement.classList.add('row')
    artistElement.setAttribute('onclick', 'onArtistSearchResultClick(this)')
    artistElement.setAttribute('data-id', artist.id)
    artistElement.setAttribute('data-name', artist.name)
    artistElement.setAttribute('data-img', artist.img)

    let artistImg = document.createElement('img')
    artistImg.classList.add('artist-img')
    artistImg.src = artist.img

    let artistName = document.createElement('span')
    artistName.innerHTML = artist.name

    let hiddenId = document.createElement('input')
    hiddenId.type = 'hidden'
    hiddenId.name = 'artists'
    hiddenId.value = artist.id


    artistElement.appendChild(artistImg)
    artistElement.appendChild(artistName)

    
    ARTISTS_SEARCH_RESULTS_CONTAINER.appendChild(artistElement)
  }
}


function searchArtists(artist) {

  request = new XMLHttpRequest()
  request.open("GET", "/search-artists?artist=" + artist, true)
  request.responseType = 'json'

  request.send()

  request.onload = function () { populateSearchResults(request.response) }
}

function onArtistSearchResultClick(artistElement) {
  let artist = {
    id: artistElement.getAttribute('data-id'),
    name: artistElement.getAttribute('data-name'),
    img: artistElement.getAttribute('data-img')
  }

  addArtist(artist)
  hideOverlay('search-artists-overlay')
}


// <div class="artist row" id="">
//   <img src="https://i.scdn.co/image/ab6761610000e5eb9e3acf1eaf3b8846e836f441">
//   <span>Taylor Swift</span>
//   <button class="button">X</button>
// </div>
function addArtist(artist) {
  // TODO: Check that artist isn't already in playlist



  let artistElement = document.createElement('div')
  artistElement.classList.add('artist')
  artistElement.classList.add('row')
  artistElement.id = artist.id

  let artistImg = document.createElement('img')
  artistImg.src = artist.img

  let artistName = document.createElement('span')
  artistName.innerHTML = artist.name

  let deleteArtist = document.createElement('button')
  deleteArtist.innerHTML = 'X'
  deleteArtist.classList.add('button')
  deleteArtist.type = 'button'
  deleteArtist.setAttribute('onclick', 'removeArtist(this)')

  let hiddenId = document.createElement('input')
  hiddenId.type = 'hidden'
  hiddenId.name = 'artists'
  hiddenId.value = artist.id


  artistElement.appendChild(artistImg)
  artistElement.appendChild(artistName)
  artistElement.appendChild(deleteArtist)
  artistElement.appendChild(hiddenId)

  
  ARTISTS_CONTAINER.appendChild(artistElement)
}

function removeArtist(deleteArtist) {
  let artistElement = deleteArtist.parentElement
  ARTISTS_CONTAINER.removeChild(artistElement)
}