<template>
  <fluffy-card class="travel-map">
    <div class="canvas" ref="map"></div>
  </fluffy-card>
</template>

<script>
  import FluffyCard from '@/Fluffy/FluffyCard';

  export default {
    components: { FluffyCard },
    props: {
      end:   { type: Object, default: () => null },
      path:  { type: Array,  required: true      },
      begin: { type: Object, required: true      },
    },
    data () {
      return {
        map: null,
        line: null,
        layer: null,
      };
    },
    watch: {
      path (path) {
        if (!path.length || !this.map || !this.line)
          return;
        const current = path[path.length - 1];
        this.line.setPath(path);
        this.map.setCenter(current);
      }
    },
    mounted () {
      this.map = new google.maps.Map(this.$refs.map, {
        zoom: 14,
        center: this.begin,
      });

      this.line = new google.maps.Polyline({
        path: this.path,
        geodesic: true,
        strokeColor: '#FF0000',
        strokeOpacity: 1.0,
        strokeWeight: 2
      });

      this.layer = new google.maps.BicyclingLayer();

      this.line.setMap(this.map);
      this.layer.setMap(this.map);
    },
  };
</script>

<style>
  .travel-map > .canvas {
    width: 100%;
    height: 320px;
  }
</style>

