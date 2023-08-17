let slideIndex = 0;
const slideshowItems = document.querySelectorAll('.slideshow-item');

function showSlide(index) {
  if (index < 0) {
    slideIndex = slideshowItems.length - 1;
  } else if (index >= slideshowItems.length) {
    slideIndex = 0;
  }

  slideshowItems.forEach(item => item.style.display = 'none');
  slideshowItems[slideIndex].style.display = 'block';
}

function changeSlide(n) {
  showSlide(slideIndex += n);
}

function autoScroll() {
  changeSlide(1);
}

showSlide(slideIndex);

const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');

prevBtn.addEventListener('click', () => changeSlide(-1));
nextBtn.addEventListener('click', () => changeSlide(1));

setInterval(autoScroll, 5000); // Change slide every 5 seconds