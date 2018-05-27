<template>
  <div class="travel-stats">
    <h4 class="time">{{ time | toTime }}</h4>
    <!-- <h3 class="distance">{{ distance }}</h3>
    <h5 class="velocity">{{ velocity }}</h5>
    <h5 class="coins">{{ coins }}</h5> -->
  </div>
</template>

<script>
  import { toTime } from '@/helpers/formatters';

  export default {
    props: {
      begin:       { type: Object, default: () => ({}) },
      coordinates: { type: Array,  dafault: () => []   },
    },
    data () {
      return {
        time: 0,
        counter: null,
      };
    },
    filters: { toTime },
    destroyed () {
      clearInterval(this.counter);
    },
    mounted () {
      const count = () => {
        if (!this.begin) {
          this.time = 0;
          return
        }
        this.time = Date.now() - this.begin.time;
      };

      count();

      this.counter = setInterval(count, 1000);
    }
  };
</script>
