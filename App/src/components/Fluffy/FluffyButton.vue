<template>
  <button
    v-on="$listeners"
    v-bind="$attrs"
    :class="classes"
    :disabled="isLoading"
  >
    <figure v-if="!!$slots.icon" class="icon">
      <slot name="icon" />
    </figure>
    <span class="text">
      <slot />
    </span>
  </button>
</template>

<script>
  export default {
    props: {
      size: {
        type: String,
        default: 'large',
        validator: (size) => ['small', 'large'].includes(size)
      },
      isLoading: {
        type: Boolean,
        default: false
      }
    },
    computed: {
      classes () {
        return [
          'fluffy-button',
          '-is-' + this.size,
          { '-is-loading': this.isLoading }
        ];
      },
    },
  };
</script>

<style>
  .fluffy-button {
    display: inline-flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    border: 0;
    padding: 0;
    line-height: 1;
    font: inherit;
    background-image: linear-gradient(135deg, #594797 0%, #412777 100%);
    box-shadow: 0 3px 6px rgba(0,0,0,.15);
    animation: pulse 2s infinite;
    transition: opacity .8s ease-in, transform .4s ease-in;
    outline: none;
  }

  .fluffy-button:active {
    opacity: .7;
    transform: scale(.95);
    transition: opacity .05s ease-out, transform .05s ease-out;
  }

  .fluffy-button:disabled {
    opacity: .7;
  }

  .fluffy-button > .text {
    line-height: 1;
    font-weight: 300;
    color: #f1f1f1;
  }

  .fluffy-button.-is-large {
    width: 222px;
    height: 222px;
  }
  .fluffy-button.-is-large > .text {
    font-size: 53pt;
  }

  .fluffy-button.-is-loading {

  }

  @keyframes pulse {
    0% {
      box-shadow: 0 0 0 0 rgba(89,71,151, 0.4);
    } 70% {
      box-shadow: 0 0 0 20px rgba(89,71,151, 0);
    } 100% {
      box-shadow: 0 0 0 0 rgba(89,71,151, 0);
    }
  }
</style>
