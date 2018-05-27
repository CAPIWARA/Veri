<template>
  <div>
    <p>Habilitar a API de geolocalização.</p>

    <trip-map :positions="positions" />

    <button v-if="isLoading && !isStarted" @click="start()">Iniciar</button>

    <button v-if="isStarted" @click="stop()">Parar</button>
  </div>
</template>

<script>
  import TripMap from '@/components/TripMap';
  import { getGeolocation } from '@/services/geolocation';

  export default {
    components: { TripMap },
    data () {
      return {
        start: null,
        watcher: null,
        positions: [],
        isStarted: false,
        isLoading: false,
      }
    },

    methods: {
      update (position) {
        this.positions.push({
          time: position.timestamp,
          coords: {
            altitude: position.coords.altitude,
            latitude: position.coords.latitude,
            longitude: position.coords.longitude
          }
        });
      },

      async initialize () {
        const permissions = window.navigator.permissions;
        await new Promise((resolve) => setTimeout(resolve, 5 * 1000));
        const answer = await permissions.query({ name: 'geolocation' });
        this.isLoading = true;
      },

      start () {
        const geolocation = window.navigator.geolocation;
        this.isStarted = true;
        this.watcher = setInterval(
          () => getGeolocation().then(this.update),
          1000 * 10
        );
      },

      stop () {
        clearInterval(this.watcher);
        this.isStarted = false;
      }
    },

    mounted () {
      this.initialize();
    }
  };
</script>

<!--
navigator.permissions.query({name:'geolocation'}).then(function(result) {
  if (result.state == 'granted') {
    report(result.state);
    geoBtn.style.display = 'none';
  } else if (result.state == 'prompt') {
    report(result.state);
    geoBtn.style.display = 'none';
    navigator.geolocation.getCurrentPosition(revealPosition,positionDenied,geoSettings);
  } else if (result.state == 'denied') {
    report(result.state);
    geoBtn.style.display = 'inline';
  }
  result.onchange = function() {
    report(result.state);
  }
});
-->
