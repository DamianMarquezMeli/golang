var map;
function initMap() {

  map = new google.maps.Map(document.getElementById("map"), {
    zoom: 13,
    center:  new google.maps.LatLng(-26.8338135,-65.2052511),
    rotateControl: true,
    mapTypeControl: true,
    mapTypeControlOptions: {
      style: google.maps.MapTypeControlStyle.HORIZONTAL_BAR,
      mapTypeIds: [
        google.maps.MapTypeId.ROADMAP,
        google.maps.MapTypeId.TERRAIN,
        google.maps.MapTypeId.SATELLITE,
        google.maps.MapTypeId.HYBRID
      ]
    }
  });



  alert("hola");
}

