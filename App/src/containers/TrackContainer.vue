<template>
  <div>
    <p>Habilitar a API de geolocalização.</p>
    <h5
      v-for="position in positions"
      :key="position.time"
    >{{ JSON.stringify(position.coords) }}</h5>

    <button v-if="isLoading && !isStarted" @click="start()">Iniciar</button>

    <button v-if="isStarted" @click="stop()">Parar</button>
  </div>
</template>

<script>
  export default {
    data () {
      return {
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
        this.watcher = setInterval(() =>
          geolocation.getCurrentPosition(
            this.update,
            (error) => console.dir(error),
            { enableHighAccuracy: true }
          ),
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
