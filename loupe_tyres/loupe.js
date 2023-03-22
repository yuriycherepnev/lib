$(document).ready(function () {
    let div = document.createElement("div");
    let img = new Image();

    let productCarousel = $("#card-product__carousel");
    img.src = productCarousel.children().children().children().attr("src");

    productCarousel.mouseover(function (e) {
        this.append(div);
        div.append(img);

        div.id = "loupe-element";
        div.classList.add('loupe');
        img.classList.add('img-loupe');

        $(window).mousemover(function (e) {
            div.style.left = e.pageX - productCarousel.offset().left - 150 +'px';
            div.style.top = e.pageY - productCarousel.offset().top - $(window).scrollTop() - 150 +'px';

            img.style.left = - (e.pageX - productCarousel.offset().left) / productCarousel.width() * img.width + 150 + 'px';
            img.style.top = - (e.pageY - productCarousel.offset().top) / productCarousel.height() * img.height + 150 + 'px';

            if (e.pageY + 20 < productCarousel.offset().top) {
                $('#loupe-element').remove();
            }
            if (e.pageY - 20 > (productCarousel.offset().top + productCarousel.height())) {
                $('#loupe-element').remove();
            }
            if (e.pageX + 20 < productCarousel.offset().left) {
                $('#loupe-element').remove();
            }
            if (e.pageX - 20 > (productCarousel.offset().left + productCarousel.width())) {
                $('#loupe-element').remove();
            }
        });
    });

    function loupe() {
        div.style.left = e.pageX - productCarousel.offset().left - 150 +'px';
        div.style.top = e.pageY - productCarousel.offset().top - $(window).scrollTop() - 150 +'px';

        img.style.left = - (e.pageX - productCarousel.offset().left) / productCarousel.width() * img.width + 150 + 'px';
        img.style.top = - (e.pageY - productCarousel.offset().top) / productCarousel.height() * img.height + 150 + 'px';

        if (e.pageY + 20 < productCarousel.offset().top) {
            $('#loupe-element').remove();
        }
        if (e.pageY - 20 > (productCarousel.offset().top + productCarousel.height())) {
            $('#loupe-element').remove();
        }
        if (e.pageX + 20 < productCarousel.offset().left) {
            $('#loupe-element').remove();
        }
        if (e.pageX - 20 > (productCarousel.offset().left + productCarousel.width())) {
            $('#loupe-element').remove();
        }
    }
});


$('.btn').on('click', myFunction);