<template>
  <div class="travel-stats">
    <h4 class="time">{{ time | toTime }}</h4>
    <h3 class="distance">
      {{ distance | toDistance }}<small class="unit">km</small>
    </h3>
    <!--
    <h5 class="coins">
      <img src="~@/assets/icons/Coin.png" alt="Moeda da plataforma" />
      {{ coins | toCoins }}
    </h5>
    <h5 class="velocity">
      {{ velocity }}
    </h5>
    -->
  </div>
</template>

<script>
  import { toTime, toDistance } from '@/helpers/formatters';
  import { distances } from '@/services/geolocation';

  export default {
    props: {
      begin:       { type: Object,  default: () => ({}) },
      isPlaying:   { type: Boolean, required: true      },
      coordinates: { type: Array,   dafault: () => []   },
    },
    data () {
      return {
        time: 0,
        counter: null,
      };
    },
    filters: { toTime, toDistance },
    computed: {
      distance () {
        if (!this.begin || !this.coordinates || this.coordinates.length === 0)
          return 0;
        return distances([ this.begin, ...this.coordinates ]);
      }
    },
    mounted () {
      const count = () => {
        if (!this.begin || !this.isPlaying)
          return
        this.time = Date.now() - this.begin.time;
      };

      count();

      this.counter = setInterval(count, 1000);
    },
    destroyed () {
      clearInterval(this.counter);
    },
  };
</script>

<style>
  .travel-stats {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 22px;
    padding-bottom: 44px;
  }

  .travel-stats > .time {
    line-height: 1;
    font-size: 30pt;
    font-weight: 400;
  }

  .travel-stats > .distance {
    line-height: 1;
    font-size: 81pt;
    font-weight: 600;
  }

  .travel-stats > .distance > .unit {
    line-height: 1;
    font-size: 22pt;
    font-weight: 600;
  }
</style>
