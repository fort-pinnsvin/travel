var map = {};

function initialize() {
    var mapProp = {
        center: new google.maps.LatLng(51.508742, -0.120850),
        zoom: 5,
        mapTypeId: google.maps.MapTypeId.TERRAIN
    };
    map = new google.maps.Map(document.getElementById("googleMap"), mapProp);
    loadMarkers(map)
    google.maps.event.addListener(map, 'dblclick', function (event) {
        placeMarker(event.latLng);
    });
}

function loadMarkers(map) {
    var array = [];
    $.ajax({
        type: "GET",
        url: "markers",
        success: function (msg) {
            array = JSON.parse(msg)
            console.table(array)

            for (var i = 0; i < array.length; i++) {
                var el = array[i]
                var marker = new google.maps.Marker({
                    position: new google.maps.LatLng(parseFloat(el.Latitude), parseFloat(el.Longitude)),
                    map: map,
                    title: el.Name,
                    id: el.Id,
                    clickListener: function () {
                        window.location.href = '/album/' + this.id;
                    }

                google.maps.event.addListener(marker, 'click', marker.clickListener);
            }
        }
    });
}

function placeMarker(location) {
    $.ajax({
        type: "GET",
        url: "markers/create",
        data: "name=New+Album&lat=" + location.lat() + "&long=" + location.lng(),
        success: function (msg) {
            result = JSON.parse(msg)
            console.log(result)
            if (result.error == 0) {
                var marker = new google.maps.Marker({
                    position: location,
                    map: map,
                    title: "New Album"
                });
                google.maps.event.addListener(marker, 'click', function () {
                    window.location.href = '/album/' + result.id;
                });
            }
        }
    });
}

google.maps.event.addDomListener(window, 'load', initialize);