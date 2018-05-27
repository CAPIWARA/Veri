<template>
  <fluffy-card :class="classes" @click.native="isFullScreen = true">
    <div class="canvas" ref="map"></div>
    <img
      v-if="isFullScreen"
      class="button"
      src="~@/assets/icons/Zoom-Out.png"
      alt="Redimensionar"
      @click.stop="isFullScreen = false"
    />
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
        // layer: null,
        isFullScreen: false,
      };
    },
    computed: {
      classes () {
        return [
          'travel-map', {
            '-is-fullscreen': this.isFullScreen
          }
        ]
      }
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
        strokeColor: '#594797',
        strokeWeight: 10,
        strokeOpacity: 1.0
      });

      // this.layer = new google.maps.BicyclingLayer();

      this.line.setMap(this.map);
      // this.layer.setMap(this.map);
    },
  };
</script>

<style>
  .travel-map {
    position: absolute;
    top: calc(100vh - 78px);
    left: calc(100vw - 96px);
    overflow: hidden;
    width: 60px;
    height: 60px;
    border: 1px solid #472F7F;
    border-radius: 50%;
    box-shadow: 0 3px 6px rgba(0,0,0,.15);
    transition: all .3s ease-out;
  }

  .travel-map.-is-fullscreen {
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border: 0;
    border-radius: 0;
    transition: all .3s ease-out;
  }

  .travel-map::after {
    position: absolute;
    display: block;
    content: '';
    width: 100vw;
    height: 100vh;
    background-color: transparent;
  }

  .travel-map:not(.-is-fullscreen) > .canvas {
    user-select: none !important;
  }

  .travel-map > .canvas {
    width: 100vw;
    height: 100vh;
  }

  .travel-map > .button {
    position: absolute;
    bottom: 18px;
    right: 36px;
    width: 60px;
    height: 60px;
    overflow: hidden;
    animation: surgir .5s ease;
  }

  @keyframes surgir {
    0% {
      opacity: 0;
    } 70% {
      opacity: .2;
    } 100% {
      opacity: 1;
    }
  }
</style>
