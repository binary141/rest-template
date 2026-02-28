import { createRouter, createWebHistory } from 'vue-router';
import { useUsersStore } from '@/stores/users';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomePage.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginPage.vue'),
      meta: { allowAnonymous: true },
    },
    {
      path: '/account',
      name: 'account',
      component: () => import('@/views/AccountPage.vue'),
    },
  ],
});

router.beforeEach((to, from, next) => {
  const store = useUsersStore();
  if (to.meta.allowAnonymous || store.loggedIn) {
    next();
  } else {
    next({ name: 'login', query: { next: to.fullPath } });
  }
});

export default router;
