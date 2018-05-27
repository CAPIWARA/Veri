<template>
  <fluffy-card class="travel-map">
    <div class="canvas" ref="map"></div>
  </fluffy-card>
</template>

<script>
  import { toPath } from '@/services/travel';
  import FluffyCard from '@/components/Fluffy/FluffyCard';

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
        this.map.panTo(current);
      }
    },
    mounted () {
      this.map = new google.maps.Map(this.$refs.map, {
        zoom: 18,
        center: toPath(this.begin),
        disableDefaultUI: true
      });

      this.line = new google.maps.Polyline({
        path: this.path,
        geodesic: true,
        strokeColor: '#FF0000',
        strokeWeight: 6,
        strokeOpacity: 1.0
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

