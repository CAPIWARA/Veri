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
  import { save } from '@/services/server';
  import { v1 as generate } from 'uuid';
  import TravelScreen from '@/components/Travel/TravelScreen';


  export default {
    components: { TravelScreen },
    data () {
      return {
        id: null,
        begin: null,
        watcher: null,
        isLoading: true,
        coordinates: [],
      }
    },
    methods: {
      async start () {
        this.isLoading = true;
        this.id = generate();
        this.begin = await services.get();
        save(this.id, this.begin);
        this.watcher = services.watch((coordinate) => {
          this.coordinates.push(coordinate);
          save(this.id, coordinate);
        });
        this.isLoading = false;
      },

      async stop () {
        services.unwatch(this.watcher);
        this.watcher = null;
        await fetch('https://veri-205319.firebaseio.com/coorinates.json', {
          method: 'POST',

        });
      },
    },
    async mounted () {
      await services.initialize();
      this.isLoading = false;
    }
  };
</script>

