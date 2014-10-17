var map = {};

function initialize() {
    var mapProp = {
        center: new google.maps.LatLng(51.508742, -0.120850),
        zoom: 2,
        mapTypeId: google.maps.MapTypeId.TERRAIN,
        disableDoubleClickZoom: true
    };
    map = new google.maps.Map(document.getElementById("googleMap"), mapProp);
    loadMarkers(map)
    google.maps.event.addListener(map, 'dblclick', function (event) {
        placeMarker(event.latLng);
    });
}


function getInfoWindow(name, desc, id) {
    var result = '<div id="content" style="color: black; width: 200px">'+
                       '<div id="siteNotice">'+
                       '</div>'+
                       '<h1 id="firstHeading" class="firstHeading" style="font-size: 18px;">' +name + '</h1>'+
                       '<div id="bodyContent">'+
                       '<p>' + (desc || '') + '</p>'+
                       '<p><a href="/album/' +id + '">'+
                       'Open album...</a></p>'+
                       '</div>'+
                       '</div>';

    return result;
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
                    infoWindow: new google.maps.InfoWindow({
                       content: getInfoWindow(el.Name, el.Description, el.Id)
                    }),
                    clickListener: function() {
                        this.infoWindow.open(map, this);
                    }
                });

                google.maps.event.addListener(marker, 'click', marker.clickListener);

            }
        }});
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
                    title: "New Album",
                    id: result.id,
                    infoWindow: new google.maps.InfoWindow({
                       content: getInfoWindow("New Album", "", result.id)
                    }),
                    clickListener: function() {
                        this.infoWindow.open(map, this);
                    }
                });
                google.maps.event.addListener(marker, 'click', marker.clickListener);
            }
        }
    });
}

google.maps.event.addDomListener(window, 'load', initialize);