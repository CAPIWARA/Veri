import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Minha corrida',
      component: () => import('@/containers/TravelContainer'),
    }
  ]
});

export default router;
