<template>
  <div class="TripMap-container" ref="map"></div>
</template>

<script>
  const toPoint = (position) => ({
    lat: position.coords.latitude,
    lng: position.coords.longitude
  });

  export default {
    props: {
      start: {
        type: Object,
        required: true
      },
      positions: {
        type: Array,
        required: true
      },
    },
    data () {
      return {
        map: null,

      }
    },
    watch: {
      positions () {
        this.route && this.route.setMap(null);
        this.route = null;

        this.route = new google.maps.Polyline({
          path: this.positions.map(toPoint),
          geodesic: true,
          strokeColor: '#FF0000',
          strokeOpacity: 1.0,
          strokeWeight: 2
        });

        this.route.setMap(this.map);
      },
    },
    mounted () {
      this.map = new google.maps.Map(this.$refs.map, {
        zoom: 14,
        center: toPoint(this.start),
      });

      var flightPath = new google.maps.Polyline({
        path: flightPlanCoordinates,
        geodesic: true,
        strokeColor: '#FF0000',
        strokeOpacity: 1.0,
        strokeWeight: 2
      });

      const layer = new google.maps.BicyclingLayer();

      layer.setMap(this.map);
    }
  }
</script>

<style>
  .TripMap-container {
    width: 100vw;
    height: 400px;
  }
</style>

