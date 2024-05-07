keyboard$.subscribe(function(key) {
  if (key.mode === "global" && key.type === "x") {
    /* Add custom keyboard handler here */
    key.claim() 
  }
})

function closeImg() {
    document.getElementById('imgBaseDiv').remove();
}

function clickAction(img) {
    let medusa = document.createElement('div');
    medusa.setAttribute('id', 'imgBaseDiv');
    medusa.setAttribute('onclick', 'closeImg()');
    let image = document.createElement('img');
    image.setAttribute('src', img.src);
    medusa.appendChild(image);
    document.body.appendChild(medusa);
}

window.onload = function () {
    for (let item of document.getElementsByTagName('img')) {
        // if (item.classList.contains('pass') === false) {
            // item.setAttribute('onclick', 'clickAction(this)');
        // }
    }
}

$(document).ready(function () {
  let productImageGroups = []
  $('.img-fluid').each(function () {
    let productImageSource = $(this).attr('src')
    let productImageTag = $(this).attr('tag')
    let productImageTitle = $(this).attr('title')
    if (productImageTitle) {
      productImageTitle = 'title="' + productImageTitle + '" '
    }
    else {
      productImageTitle = ''
    }
    $(this).
        wrap('<a class="boxedThumb ' + productImageTag + '" ' +
            productImageTitle + 'href="' + productImageSource + '"></a>')
    productImageGroups.push('.' + productImageTag)
  })
  jQuery.unique(productImageGroups)
  productImageGroups.forEach(productImageGroupsSet)

  function productImageGroupsSet (value) {
    $(value).simpleLightbox()
  }
})