
//js snippets from browser tests

var jq = document.createElement('script');
jq.src = "https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js";
document.getElementsByTagName('head')[0].appendChild(jq);
// ... give time for script to load, then type (or see below for non wait option)
jQuery.noConflict();


json = jQuery.getJSON('http://192.168.99.100/search?query=rick_and_morty', function(data) {

    for (index = 0; index < data.length; ++index) {
        fullUrl = data[index];
        code = fullUrl.slice(32);
        document.write("<iframe title='YouTube video player' type=\"text/html\" width='693' height='390' src='http://www.youtube.com/embed/"+code+"' frameborder='0' allow='accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture' allowfullscreen></iframe>");
    }


});


for (index = 0; index < data.length; ++index) {
    fullUrl = data[index];
    code = fullUrl.slice(32);
    document.write("<iframe title='YouTube video player' type=\"text/html\" width='693' height='390' src='http://www.youtube.com/embed/"+code+"' frameborder='0' allow='accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture' allowfullscreen></iframe>");
}


document.write("<iframe title='YouTube video player' type=\"text/html\" width='693' height='390' src='http://www.youtube.com/embed/"+code+"?autoplay=1&enablejsapi=1' frameborder='0' allow='autoplay'></iframe>");
    jQuery(document).ready(function( $ ) {
    $('.video-selector iframe')[0].contentWindow.postMessage('{"event":"command","func":"playVideo","args":""}', '*');
    });



    document.write("<iframe title='YouTube video player' type=\"text/html\" width='693' height='390' src='http://www.youtube.com/embed/8WDNgfnZ3HM?autoplay=1&mute=1' frameborder='0' allow='accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture' allowfullscreen></iframe>");
            