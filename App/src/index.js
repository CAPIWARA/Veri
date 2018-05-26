import Vue from 'vue';
import App from '@/App.vue';

const app = new Vue({
  el: '#app',
  render: (λ) => λ(App)
});

/**
 * Inicializa uma segunda instância do Vue. Essa instância permite que os
 * componentes se comunicarem por eventos globais.
 */
global.EventBus = new Vue();
