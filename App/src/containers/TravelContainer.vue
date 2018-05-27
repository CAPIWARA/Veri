<template>
  <travel-screen
    :begin="begin"
    :isLoading="isLoading"
    :isPlaying="!!watcher"
    :coordinates="coordinates"
    @start="start()"
    @stop="stop()"
  />
</template>

<script>
  import * as services from '@/services/geolocation';
  import TravelScreen from '@/components/Travel/TravelScreen';

  export default {
    components: { TravelScreen },
    data () {
      return {
        coordinates: [],
        begin: null,
        watcher: null,
        isLoading: true,
      }
    },
    methods: {
      async start () {
        this.isLoading = true;
        this.begin = await services.get();
        this.watcher = services.watch((coordinate) => {
          this.coordinates.push(coordinate);
        });
        this.isLoading = false;
      },

      stop () {
        services.unwatch(this.watcher);
        this.watcher = null;
      },
    },
    async mounted () {
      await services.initialize();
      this.isLoading = false;
    }
  };
</script>

