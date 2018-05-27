<template>
  <section class="travel-screen">
    <travel-stats
      :begin="begin"
      :isPlaying="isPlaying"
      :coordinates="coordinates"
    />

    <travel-button
      :isLoading="isLoading"
      :isPlaying="isPlaying"
      @stop="stop()"
      @start="start()"
    />

    <travel-map
      v-if="isPlaying"
      :end="end"
      :path="path"
      :begin="begin"
    />
  </section>
</template>

<script>
  import { toPath } from '@/services/travel';
  import TravelMap from '@/components/Travel/TravelMap';
  import TravelStats from '@/components/Travel/TravelStats';
  import TravelButton from '@/components/Travel/TravelButton';

  export default {
    components: { TravelMap, TravelStats, TravelButton },
    props: {
      isLoading:   Boolean,
      isPlaying:   Boolean,
      end:         { type: Object, default: () => null },
      begin:       { type: Object, default: () => null },
      coordinates: { type: Array,  required: true      },
    },
    computed: {
      path () {
        return this.coordinates.map(toPath);
      }
    },
    methods: {
      stop () {
        this.$emit('stop');
      },
      start () {
        this.$emit('start');
      },
    },
  }
</script>

<style>
  .travel-screen {
    height: 100%;
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: flex-start;
  }
</style>
